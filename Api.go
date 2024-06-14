package bwssdk

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// API represents the API client with necessary credentials and configuration
type API struct {
	MerchantNo     string // Merchant number
	APIKey         string // API key
	GatewayAddress string // Gateway address
	CallURL        string // Callback URL
}

// NewAPI creates a new API instance
// merchantNo: Merchant number
// apiKey: API key
// gatewayAddress: Gateway address
// callURL: Callback URL
func NewAPI(merchantNo, apiKey, gatewayAddress, callURL string) *API {
	return &API{
		MerchantNo:     merchantNo,
		APIKey:         apiKey,
		GatewayAddress: gatewayAddress,
		CallURL:        callURL,
	}
}

// Request makes an API request
// method: API endpoint method
// body: Request body parameters
// Returns: A map containing the response data or an error if the request fails
func (api *API) Request(method string, body interface{}) (map[string]interface{}, error) {
	time := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strconv.Itoa(rand.Intn(899999) + 100000)

	// Marshal the request body to JSON
	var bodyData []byte
	var err error
	if method == "/mch/support-coins" {
		bodyData, err = json.Marshal(body)
	} else {
		bodyData, err = json.Marshal([]interface{}{body})
	}
	if err != nil {
		return nil, err
	}

	// Generate signature
	sign := api.signature(string(bodyData), time, nonce)

	// Create request parameters
	params := map[string]string{
		"timestamp": time,
		"nonce":     nonce,
		"sign":      sign,
		"body":      string(bodyData),
	}

	// Perform the HTTP request
	response, err := api.doRequest(method, params)
	if err != nil {
		return nil, err
	}

	// Parse the response body
	result := make(map[string]interface{})
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	// Check for errors in the response
	// if err := api.checkErrorAndThrow(result); err != nil {
	// 	return nil, err
	// }

	return result, nil
}

// signature generates a signature for the API request
// body: Request body as JSON string
// time: Current timestamp as string
// nonce: Random nonce value
// Returns: MD5 hash signature as a string
func (api *API) signature(body, time, nonce string) string {
	data := []byte(body + api.APIKey + nonce + time)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

// doRequest sends the HTTP request to the API
// method: API endpoint method
// params: Request parameters as a map
// Returns: The response body as bytes or an error if the request fails
// doRequest sends the HTTP request to the API and checks the HTTP status code
func (api *API) doRequest(method string, params map[string]string) ([]byte, error) {
	// 将参数编码为 JSON
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %v", err)
	}

	// 构建请求 URL
	url := api.GatewayAddress + method

	fmt.Println("url: " + url)
	fmt.Println(string(jsonData))
	fmt.Println()

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %v", err)
	}
	defer resp.Body.Close()

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return responseBody, nil
}

// checkErrorAndThrow checks for errors in the API response
// result: Response data as a map
// Returns: An error if the response contains an error code
func (api *API) checkErrorAndThrow(result map[string]interface{}) error {
	if code, ok := result["code"].(int64); !ok || int(code) != 200 {
		message := result["message"].(string)
		return fmt.Errorf("ByteDispatchException: code=%d, message=%s", int(code), message)
	}
	return nil
}

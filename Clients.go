package bwssdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Clients represents the API client with necessary credentials and configuration
type Clients struct {
	API
}

// NewClients creates a new Clients instance
// merchantNo: Merchant number
// apiKey: API key
// gatewayAddress: Gateway address
// callURL: Callback URL
func NewClients(merchantNo, apiKey, gatewayAddress, callURL string) *Clients {
	return &Clients{
		API: API{
			MerchantNo:     merchantNo,
			APIKey:         apiKey,
			GatewayAddress: gatewayAddress,
			CallURL:        callURL,
		},
	}
}

// SupportCoins fetches supported coins information for the merchant
// showBalance: Whether to query the balance (true: fetch, false: don't fetch)
// Returns: A map containing the response data or an error if the request fails
func (c *Clients) SupportCoins(showBalance bool) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"merchantId":  c.MerchantNo,
		"showBalance": showBalance,
	}
	return c.Request("/mch/support-coins", body)
}

// CreateAddress creates a new address
// mainCoinType: Main coin type number
// walletId: Wallet ID, defaults to generating address based on the main wallet
// alias: Address alias
// Returns: A map containing the response data or an error if the request fails
func (c *Clients) CreateAddress(mainCoinType string, walletId, alias *string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"merchantId":   c.MerchantNo,
		"mainCoinType": mainCoinType,
		"callUrl":      c.CallURL,
	}
	if walletId != nil {
		body["walletId"] = *walletId
	}
	if alias != nil {
		body["alias"] = *alias
	}
	return c.Request("/mch/address/create", body)
}

// CheckAddress validates the legality of an address
// mainCoinType: Main coin type number
// address: Address to be validated
// Returns: A map containing the response data or an error if the request fails
func (c *Clients) CheckAddress(mainCoinType, address string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"merchantId":   c.MerchantNo,
		"mainCoinType": mainCoinType,
		"address":      address,
	}
	return c.Request("/mch/check/address", body)
}

// ExistAddress checks if an address exists
// mainCoinType: Main coin type number
// address: Address to be checked
// Returns: A map containing the response data or an error if the request fails
func (c *Clients) ExistAddress(mainCoinType, address string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"merchantId":   c.MerchantNo,
		"mainCoinType": mainCoinType,
		"address":      address,
	}
	return c.Request("/mch/exist/address", body)
}

// Withdraw requests a withdrawal
// businessId: Business ID, must be unique
// mainCoinType: Main coin type number
// coinType: Sub coin type number
// address: Address to withdraw to
// amount: Amount to withdraw
// memo: Memo or remark
// Returns: A map containing the response data or an error if the request fails
func (c *Clients) Withdraw(businessId, mainCoinType, coinType, address string, amount float64, memo *string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"merchantId":   c.MerchantNo,
		"mainCoinType": mainCoinType,
		"coinType":     coinType,
		"address":      address,
		"businessId":   businessId,
		"amount":       amount,
		"callUrl":      c.CallURL,
	}
	if memo != nil {
		body["memo"] = *memo
	}
	return c.Request("/mch/withdraw", body)
}

// PrintLog logs custom messages
// msg: Message to log
func (c *Clients) PrintLog(msg string) {
	logDir := "log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, 0777)
	}
	path := fmt.Sprintf("%s/%s.log", logDir, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()
	logMsg := fmt.Sprintf("【%s】%s\r\n\r\n", time.Now().Format("2006-01-02 15:04:05"), msg)
	file.WriteString(logMsg)
}

// Callback handles callback requests
// Returns: A string indicating success or an error if the request fails
func (c *Clients) Callback(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	nonce := r.FormValue("nonce")
	timestamp := r.FormValue("timestamp")
	sign := r.FormValue("sign")

	// Verify signature
	signCheck := c.signature(body, timestamp, nonce)
	if sign != signCheck {
		http.Error(w, "Signature error", http.StatusBadRequest)
		return
	}

	var callbackBody struct {
		TradeType int `json:"tradeType"`
		Status    int `json:"status"`
	}
	if err := json.Unmarshal([]byte(body), &callbackBody); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	// Business logic based on tradeType and status
	switch callbackBody.TradeType {
	case 1: // Deposit callback
		if callbackBody.Status == 3 {
			// Business logic for successful deposit
		}
	case 2: // Withdrawal callback
		switch callbackBody.Status {
		case 0:
			// Business logic for pending review
		case 1:
			// Business logic for review success
		case 2:
			// Business logic for review rejection
		case 3:
			// Business logic for successful transaction
		case 4:
			// Business logic for failed transaction
		}
	}

	w.Write([]byte("success"))
}

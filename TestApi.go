package bwssdk

import (
	"fmt"
	"log"
)

// Main 函数，
func main() {
	merchantNo := "80006"
	apiKey := "071FE257d241C0b3f4faeEe13019871E1eb399f5"
	gatewayAddress := "https://api.ByteWallets.io"
	callURL := "https://localhost/callUrl"

	client := NewClients(merchantNo, apiKey, gatewayAddress, callURL)

	// 查询支持的币种信息
	result, err := client.SupportCoins(true)
	if err != nil {
		log.Fatalf("SupportCoins Error: %v", err)
	}
	fmt.Printf("SupportCoins Result: %+v\n", result)

	// 创建新地址
	result, err = client.CreateAddress("62", nil, nil)
	if err != nil {
		log.Fatalf("CreateAddress Error: %v", err)
	}
	fmt.Printf("CreateAddress Result: %+v\n", result)

	// 验证地址合法性
	result, err = client.CheckAddress("62", "TEpK1aWkeuskfo8reeeMqG7hdJ5tRytyAF")
	if err != nil {
		log.Fatalf("CheckAddress Error: %v", err)
	}
	fmt.Printf("CheckAddress Result: %+v\n", result)

	// 检查地址是否存在
	result, err = client.ExistAddress("62", "TEpK1aWkeuskfo8reeeMqG7hdJ5tRytyAF")
	if err != nil {
		log.Fatalf("ExistAddress Error: %v", err)
	}
	fmt.Printf("ExistAddress Result: %+v\n", result)

	// 申请提币
	memo := "Withdrawal for invoice #12345"
	result, err = client.Withdraw("sn00001", "62", "62", "TEpK1aWkeuskfo8reeeMqG7hdJ5tRytyAF", 10, &memo)
	if err != nil {
		log.Fatalf("Withdraw Error: %v", err)
	}
	fmt.Printf("Withdraw Result: %+v\n", result)
}

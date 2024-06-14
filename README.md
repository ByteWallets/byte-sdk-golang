
# API Client Example

## Overview
This example demonstrates how to use the `Clients` class to call various API methods, including querying supported coins, creating new addresses, validating addresses, checking if an address exists, and applying for withdrawals.

## Prerequisites
Make sure you have Go installed. If not, you can download and install it from the [official Go website](https://golang.org/dl/).

## Getting Started

1. Clone or download this repository to your local machine.
2. Ensure you have configured your API settings, including the merchant number, API key, gateway address, and callback URL.
3. Create a `main.go` file in your project directory and copy the following content into the file.

```go
package main

import (
	"fmt"
	"log"
)

// Define the API and Clients structs
type API struct {
	MerchantNo     string
	APIKey         string
	GatewayAddress string
	CallURL        string
}

type Clients struct {
	API
}

// NewClients creates a new Clients instance
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

// Example methods for the Clients struct
func (c *Clients) SupportCoins(showBalance bool) (map[string]interface{}, error) {
	// Dummy implementation for example purposes
	return map[string]interface{}{"dummy": "data"}, nil
}

func (c *Clients) CreateAddress(mainCoinType string, walletId, alias *string) (map[string]interface{}, error) {
	// Dummy implementation for example purposes
	return map[string]interface{}{"dummy": "data"}, nil
}

func (c *Clients) CheckAddress(mainCoinType, address string) (map[string]interface{}, error) {
	// Dummy implementation for example purposes
	return map[string]interface{}{"dummy": "data"}, nil
}

func (c *Clients) ExistAddress(mainCoinType, address string) (map[string]interface{}, error) {
	// Dummy implementation for example purposes
	return map[string]interface{}{"dummy": "data"}, nil
}

func (c *Clients) Withdraw(businessId, mainCoinType, coinType, address string, amount float64, memo *string) (map[string]interface{}, error) {
	// Dummy implementation for example purposes
	return map[string]interface{}{"dummy": "data"}, nil
}

func main() {
	merchantNo := "30000"
	apiKey := "c789324a4XXXXXXX88ec8a3872"
	gatewayAddress := "https://sig11.Byte.io"
	callURL := "https://localhost/callUrl"

	client := NewClients(merchantNo, apiKey, gatewayAddress, callURL)

	// Query supported coins
	result, err := client.SupportCoins(true)
	if err != nil {
		log.Fatalf("SupportCoins Error: %v", err)
	}
	fmt.Printf("SupportCoins Result: %+v
", result)

	// Create new address
	result, err = client.CreateAddress("195", nil, nil)
	if err != nil {
		log.Fatalf("CreateAddress Error: %v", err)
	}
	fmt.Printf("CreateAddress Result: %+v
", result)

	// Validate address
	result, err = client.CheckAddress("195", "TEpK1aWkjDue6j8reeeMqG7hdJ5tRytyAF")
	if err != nil {
		log.Fatalf("CheckAddress Error: %v", err)
	}
	fmt.Printf("CheckAddress Result: %+v
", result)

	// Check if address exists
	result, err = client.ExistAddress("195", "TEpK1aWkjDue6j8reeeMqG7hdJ5tRytyAF")
	if err != nil {
		log.Fatalf("ExistAddress Error: %v", err)
	}
	fmt.Printf("ExistAddress Result: %+v
", result)

	// Apply for withdrawal
	memo := "Withdrawal for invoice #12345"
	result, err = client.Withdraw("sn00001", "195", "195", "TEpK1aWkjDue6j8reeeMqG7hdJ5tRytyAF", 10, &memo)
	if err != nil {
		log.Fatalf("Withdraw Error: %v", err)
	}
	fmt.Printf("Withdraw Result: %+v
", result)
}
```

4. Run the following command in your project directory to build and run the program:

```sh
go run main.go
```

## Method Descriptions

### `SupportCoins(showBalance bool)`

Queries supported coins information.

- **Parameters**:
  - `showBalance`: Whether to query the balance (true to fetch, false not to fetch).
- **Returns**: A map containing the supported coins information.

### `CreateAddress(mainCoinType string, walletId, alias *string)`

Creates a new address.

- **Parameters**:
  - `mainCoinType`: Main coin type number.
  - `walletId`: Wallet ID, default is nil.
  - `alias`: Address alias, default is nil.
- **Returns**: A map containing the new address information.

### `CheckAddress(mainCoinType, address string)`

Validates the legality of an address.

- **Parameters**:
  - `mainCoinType`: Main coin type number.
  - `address`: Address to be validated.
- **Returns**: Validation result.

### `ExistAddress(mainCoinType, address string)`

Checks if an address exists.

- **Parameters**:
  - `mainCoinType`: Main coin type number.
  - `address`: Address to be checked.
- **Returns**: Existence check result.

### `Withdraw(businessId, mainCoinType, coinType, address string, amount float64, memo *string)`

Applies for a withdrawal.

- **Parameters**:
  - `businessId`: Business ID, must be unique.
  - `mainCoinType`: Main coin type number.
  - `coinType`: Sub coin type number.
  - `address`: Withdrawal address.
  - `amount`: Amount to withdraw.
  - `memo`: Memo or remark.
- **Returns**: Withdrawal application result.

## Contact

If you have any questions or suggestions, please contact the author.

Thank you for using this example program!

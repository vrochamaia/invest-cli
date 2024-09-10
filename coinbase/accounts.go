package coinbase

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	requestHost = "api.coinbase.com"
	requestPath = "/api/v3/brokerage/accounts"
)

type AvailableBalance struct {
	Value string `json:"value"`
}

type CoinbaseAccount struct {
	Name             string           `json:"name"`
	Currency         string           `json:"currency"`
	AvailableBalance AvailableBalance `json:"available_balance"`
}

type CoinbaseAccounts struct {
	Accounts []CoinbaseAccount `json:"accounts"`
}

func formatResponse(input string) []CoinbaseAccount {
	var result CoinbaseAccounts

	error := json.Unmarshal([]byte(input), &result)

	if error != nil {
		panic(error)
	}

	return result.Accounts
}

func Accounts(isDevelopment bool) {

	var accounts string

	if isDevelopment {
		fmt.Println("Using mock data...")

		jsonFile, _ := os.ReadFile("./payloadexample.json")
		accounts = string(jsonFile)
	} else {
		accounts = Get(GetRequestInput{RequestHost: requestHost, RequestPath: requestPath})

	}

	formattedResponse := formatResponse(accounts)

	fmt.Println(formattedResponse)
}

package coinbase

import (
	"encoding/json"
	"fmt"
	"investcli/coinconvert"
	"os"
	"strconv"
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

	accountsMap := make(map[string]float64)
	CADTotalAmount := 0.0

	for _, account := range result.Accounts {
		accountBalance, _ := strconv.ParseFloat(account.AvailableBalance.Value, 64)

		if accountBalance > 0 {
			CADAmount := coinconvert.CoinConvert(coinconvert.CoinConvertInput{FromCurrency: account.Currency, ToCurrency: "CAD", Amount: accountBalance})

			accountsMap[account.Currency] = CADAmount

			CADTotalAmount += CADAmount
		}
	}

	for key, value := range accountsMap {
		fmt.Println(key, fmt.Sprintf("CA$ %.2f", value), "| %", fmt.Sprintf("%.2f", value/CADTotalAmount*100))
		fmt.Println("")
	}

	fmt.Println("Total amount", fmt.Sprintf("CA$ %.2f", CADTotalAmount))

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

	formatResponse(accounts)
}

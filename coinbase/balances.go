package coinbase

import (
	"encoding/json"
	"fmt"
	"investcli/coin"
	"investcli/utils"
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

func parseResponse(input string) []coin.Balance {
	var result CoinbaseAccounts

	error := json.Unmarshal([]byte(input), &result)

	if error != nil {
		panic(error)
	}

	var coinBalances []coin.Balance

	for _, account := range result.Accounts {
		parsedBalance, _ := strconv.ParseFloat(account.AvailableBalance.Value, 64)

		coinBalances = append(coinBalances, coin.Balance{Currency: account.Currency, AvailableBalance: parsedBalance})
	}

	return coinBalances
}

func Balances() []coin.Balance {

	var accounts string

	if utils.IsTestEnv() {
		fmt.Println("Using Coinbase mock data...")

		jsonFile, _ := os.ReadFile("./coinbase-example.json")
		accounts = string(jsonFile)
	} else {
		accounts = Get(GetRequestInput{RequestHost: requestHost, RequestPath: requestPath})
	}

	return parseResponse(accounts)
}

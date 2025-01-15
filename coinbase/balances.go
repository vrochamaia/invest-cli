package coinbase

import (
	"encoding/json"
	"fmt"
	"investcli/http"
	"investcli/utils"
	"investcli/wallet"
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

type authenticationTokenInput struct {
	requestHost   string
	requestPath   string
	requestMethod string
}

func authenticationToken(input authenticationTokenInput) (string, error) {

	uri := fmt.Sprintf("%s %s%s", input.requestMethod, input.requestHost, input.requestPath)

	jwtToken, error := buildJWT(uri)

	return jwtToken, error
}

func parseResponse(input string) []wallet.Balance {
	var result CoinbaseAccounts

	error := json.Unmarshal([]byte(input), &result)

	if error != nil {
		panic(error)
	}

	var coinBalances []wallet.Balance

	for _, account := range result.Accounts {
		parsedBalance, _ := strconv.ParseFloat(account.AvailableBalance.Value, 64)

		coinBalances = append(coinBalances, wallet.Balance{Currency: account.Currency, AvailableBalance: parsedBalance})
	}

	return coinBalances
}

func Balances() []wallet.Balance {
	if utils.IsTestEnv() {
		fmt.Println("Using Coinbase mock data...")

		jsonFile, _ := os.ReadFile("./coinbase-mock-data.json")

		return parseResponse(string(jsonFile))
	}

	requestMethod := "GET"

	token, error := authenticationToken(authenticationTokenInput{requestHost: requestHost,
		requestPath: requestPath, requestMethod: requestMethod})

	if error != nil {
		fmt.Println("Error while getting Coinbase authentication token: ", error)

		return []wallet.Balance{}
	}

	accounts := http.Request(http.RequestInput{RequestMethod: requestMethod, RequestHost: requestHost, RequestPath: requestPath, Headers: map[string]string{"Authorization": "Bearer " + token}})

	return parseResponse(accounts)
}

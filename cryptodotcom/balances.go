package cryptodotcom

import (
	"encoding/json"
	"fmt"
	"investcli/coin"
	"investcli/http"
	"investcli/utils"
	"os"
	"strconv"
	"time"
)

type AccountBalance struct {
	Balance  string `json:"quantity"`
	Currency string `json:"instrument_name"`
}

type AccountData struct {
	Balances []AccountBalance `json:"position_balances"`
}

type UserBalanceResult struct {
	Data []AccountData `json:"data"`
}

type CryptoDotComUserBalance struct {
	Result UserBalanceResult `json:"result"`
}

func parseResponse(input string) []coin.Balance {
	var userBalance CryptoDotComUserBalance

	error := json.Unmarshal([]byte(input), &userBalance)

	if error != nil {
		panic(error)
	}

	var coinBalances []coin.Balance

	// assuming user has only one master account
	accountBalances := userBalance.Result.Data[0].Balances

	for _, accountBalance := range accountBalances {
		parsedBalance, _ := strconv.ParseFloat(accountBalance.Balance, 64)

		coinBalances = append(coinBalances, coin.Balance{Currency: accountBalance.Currency, AvailableBalance: parsedBalance})
	}

	return coinBalances
}

func Balances() []coin.Balance {
	if utils.IsTestEnv() {
		fmt.Println("Using Crypto.com mock data...")

		jsonFile, _ := os.ReadFile("./cryptodotcom-mock-data.json")

		return parseResponse(string(jsonFile))
	}

	requestBody, error := signRequest(map[string]interface{}{
		"id":     10,
		"method": "private/user-balance",
		"params": map[string]interface{}{},
		"nonce":  time.Now().UnixMilli(),
	})

	if error != nil {
		fmt.Println("Error signing Crypto.com request: ", error)

		return []coin.Balance{}
	}

	response := http.Request(http.RequestInput{
		RequestHost:   "api.crypto.com",
		RequestMethod: "POST",
		RequestPath:   "/exchange/v1/private/user-balance",
		Body:          requestBody,
		Headers:       map[string]string{"Content-Type": "application/json"},
	})

	return parseResponse(response)
}

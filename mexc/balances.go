package mexc

import (
	"encoding/json"
	"fmt"
	"investcli/http"
	"investcli/utils"
	"investcli/wallet"
	"os"
	"strconv"
	"time"
)

type mexcApiKey struct {
	ApiKeys utils.ApiKeys `json:"mexc"`
}

type MexcBalance struct {
	Currency string `json:"asset"`
	Balance  string `json:"free"`
}

type MexcBalances struct {
	Balances []MexcBalance `json:"balances"`
}

func Balances() []wallet.Balance {
	if utils.IsTestEnv() {
		fmt.Println("Using Mexc mock data...")

		jsonFile, _ := os.ReadFile("./mexc-mock-data.json")

		return formatResponse(string(jsonFile))
	}

	apiKey := utils.GetDataFromJson[mexcApiKey]("./secrets.json").ApiKeys

	if apiKey.Key == "" || apiKey.PrivateKey == "" {
		fmt.Println("Could not get MEXC API keys")

		return []wallet.Balance{}
	}

	timestamp := fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))

	signature := requestSignature(map[string]string{"timestamp": timestamp}, apiKey.PrivateKey)

	response := http.Request(http.RequestInput{
		RequestHost:   "api.mexc.com",
		RequestMethod: "GET",
		RequestPath:   fmt.Sprintf("/api/v3/account?timestamp=%s&signature=%s", timestamp, signature),
		Headers:       map[string]string{"Content-Type": "application/json", "x-mexc-apikey": apiKey.Key},
	})

	return formatResponse(response)
}

func formatResponse(input string) []wallet.Balance {
	var mexcBalances MexcBalances

	error := json.Unmarshal([]byte(input), &mexcBalances)

	if error != nil {
		panic(error)
	}

	var walletBalances []wallet.Balance

	for _, accountBalance := range mexcBalances.Balances {
		parsedBalance, _ := strconv.ParseFloat(accountBalance.Balance, 64)

		walletBalances = append(walletBalances, wallet.Balance{Currency: accountBalance.Currency, AvailableBalance: parsedBalance})
	}

	return walletBalances
}

package coin

import (
	"encoding/json"
	"fmt"
	"investcli/coinconvert"
	"os"
)

type Balance struct {
	Currency         string
	AvailableBalance float64
}

func fetchDesiredWeights() map[string]float32 {
	var desiredWeights map[string]float32

	jsonFile, _ := os.ReadFile("./desired-wallet.json")

	error := json.Unmarshal([]byte(string(jsonFile)), &desiredWeights)

	if error != nil {
		panic(error)
	}

	return desiredWeights
}

func CalculateProportionAmongBalances(balances []Balance) {
	accountsMap := make(map[string]float64)
	CADTotalAmount := 0.0

	desiredWeights := fetchDesiredWeights()

	for _, balance := range balances {
		availableBalance := balance.AvailableBalance
		currency := balance.Currency

		if availableBalance > 0 {
			CADAmount := coinconvert.CoinConvert(coinconvert.CoinConvertInput{FromCurrency: currency, ToCurrency: "CAD", Amount: availableBalance})

			// The same crypto token can appear more than once in the array
			accountsMap[currency] = accountsMap[currency] + CADAmount

			CADTotalAmount += CADAmount
		}
	}

	for key, value := range accountsMap {
		currentWeight := value / CADTotalAmount * 100
		desiredWeight := desiredWeights[key]

		fmt.Println(key, fmt.Sprintf("CA$ %.2f", value), "|", fmt.Sprintf("Current Weigth: %.2f", currentWeight), "%", "|", fmt.Sprintf("Desired weigth: %.2f", desiredWeight), "%")
		fmt.Println("")
	}

	fmt.Println("Total amount:", fmt.Sprintf("CA$ %.2f", CADTotalAmount))
}

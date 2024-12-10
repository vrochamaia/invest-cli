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

	fileName := "desired-wallet.json"

	jsonFile, _ := os.ReadFile(fileName)

	error := json.Unmarshal([]byte(string(jsonFile)), &desiredWeights)

	if error != nil {
		fmt.Printf("Could not fetch desired wallet. This is expected if you didn't set up the %s file.\n", fileName)
	}

	weightsPercentageSum := 0.0
	for _, value := range desiredWeights {
		weightsPercentageSum += float64(value)
	}

	fmt.Printf("Weights Percentage Sum: %.2f\n\n", weightsPercentageSum)

	return desiredWeights
}

func CalculateProportionAmongBalances(balances []Balance) {
	accountsMap := make(map[string]float64)
	totalBalance := 0.0

	desiredWeights := fetchDesiredWeights()

	for _, balance := range balances {
		availableBalance := balance.AvailableBalance
		currency := balance.Currency

		if availableBalance > 0 {
			usdAmount := coinconvert.CoinConvert(coinconvert.CoinConvertInput{FromCurrency: currency, ToCurrency: "USD", Amount: availableBalance})

			// The same crypto token can appear more than once in the array
			accountsMap[currency] = accountsMap[currency] + usdAmount

			totalBalance += usdAmount
		}
	}

	for key, value := range accountsMap {
		currentWeight := value / totalBalance * 100
		desiredWeight := desiredWeights[key]

		desiredMinusCurrentWeight := float64(desiredWeight) - currentWeight
		valueRequiredToBalance := (desiredMinusCurrentWeight * totalBalance) / 100

		fmt.Println(key, fmt.Sprintf("$ %.2f", value), "|", fmt.Sprintf("Current Weigth: %.2f", currentWeight), "%", "|", fmt.Sprintf("Desired weigth: %.2f", desiredWeight), "%", "|", fmt.Sprintf("Required value to ideal balance: $ %.2f", valueRequiredToBalance))
		fmt.Println("")
	}

	fmt.Println("Total balance:", fmt.Sprintf("$ %.2f", totalBalance))
}

package coin

import (
	"fmt"
	"investcli/coinconvert"
)

type Balance struct {
	Currency         string
	AvailableBalance float64
}

func CalculateProportionBetweenBalances(balances []Balance) {
	accountsMap := make(map[string]float64)
	CADTotalAmount := 0.0

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
		fmt.Println(key, fmt.Sprintf("CA$ %.2f", value), "|", fmt.Sprintf("%.2f", value/CADTotalAmount*100), "%")
		fmt.Println("")
	}

	fmt.Println("Total amount:", fmt.Sprintf("CA$ %.2f", CADTotalAmount))
}

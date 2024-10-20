package coinconvert

import (
	"encoding/json"
	"fmt"
	"investcli/http"
	"investcli/utils"
)

type CoinConvertInput struct {
	FromCurrency string
	ToCurrency   string
	Amount       float64
}

func CoinConvert(input CoinConvertInput) float64 {
	if utils.IsTestEnv() {
		return 50
	}

	response := http.Request(http.RequestInput{RequestMethod: "GET", RequestHost: "api.coinconvert.net", RequestPath: fmt.Sprintf("/convert/%s/%s?amount=%f", input.FromCurrency, input.ToCurrency, input.Amount)})

	var parsedResponse map[string]interface{}

	error := json.Unmarshal([]byte(response), &parsedResponse)

	if error != nil {
		panic(error)
	}

	if parsedResponse["status"] != "success" {
		panic(fmt.Sprintf("Failed to convert the currency: %s", parsedResponse["message"]))
	}

	return parsedResponse[input.ToCurrency].(float64)
}

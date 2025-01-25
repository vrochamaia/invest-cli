package coinconvert

import (
	"encoding/json"
	"fmt"
	"investcli/http"
	"investcli/utils"
	"strconv"
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

	response := http.Request(http.RequestInput{RequestMethod: "GET", RequestHost: "api.coinconvert.net", RequestPath: fmt.Sprintf("/convert/%s/%s?amount=%s", input.FromCurrency, input.ToCurrency, strconv.FormatFloat(input.Amount, 'f', -1, 64))})

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

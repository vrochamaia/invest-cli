package coinbase

import (
	"encoding/json"
	"os"
)

type ApiKey struct {
	KeyName   string `json:"name"`
	KeySecret string `json:"privateKey"`
}

func fetchApiKeyFromSecrets() ApiKey {
	jsonFile, error := os.ReadFile("./secrets.json")

	if error != nil {
		panic(error)
	}

	var apiKey ApiKey

	jsonParseError := json.Unmarshal(jsonFile, &apiKey)

	if jsonParseError != nil {
		panic(jsonParseError)
	}

	return apiKey
}

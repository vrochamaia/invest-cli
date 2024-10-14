package utils

import (
	"encoding/json"
	"os"
)

type ApiKeys struct {
	Key        string `json:"key"`
	PrivateKey string `json:"privateKey"`
}

func GetDataFromJson[T any](filePath string) *T {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var data T

	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		panic(err)
	}

	return &data
}

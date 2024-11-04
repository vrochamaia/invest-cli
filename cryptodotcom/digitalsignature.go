package cryptodotcom

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"investcli/utils"
	"sort"
	"strings"
)

type cryptoDotComApiKey struct {
	ApiKeys utils.ApiKeys `json:"cryptoDotCom"`
}

func signRequest(requestBody map[string]interface{}) (string, error) {
	apiKey := utils.GetDataFromJson[cryptoDotComApiKey]("./secrets.json").ApiKeys

	if apiKey.Key == "" || apiKey.PrivateKey == "" {
		return "", errors.New("could not get Crypto.com API keys")
	}

	id := requestBody["id"]
	method := requestBody["method"].(string)
	params := requestBody["params"].(map[string]interface{})
	nonce := requestBody["nonce"].(int64)

	paramsString := objectToString(params)

	sigPayload := fmt.Sprintf("%s%v%s%s%d", method, id, apiKey.Key, paramsString, nonce)

	hmac := hmac.New(sha256.New, []byte(apiKey.PrivateKey))
	hmac.Write([]byte(sigPayload))
	signature := hex.EncodeToString(hmac.Sum(nil))

	requestBody["sig"] = signature
	requestBody["api_key"] = apiKey.Key

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	return string(jsonBody), nil
}

func arrayToString(arr []interface{}) string {
	var result strings.Builder
	for _, v := range arr {
		switch val := v.(type) {
		case map[string]interface{}:
			result.WriteString(objectToString(val))
		case []interface{}:
			result.WriteString(arrayToString(val))
		default:
			result.WriteString(fmt.Sprint(val))
		}
	}
	return result.String()
}

func objectToString(obj map[string]interface{}) string {
	if obj == nil {
		return ""
	}
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var result strings.Builder
	for _, k := range keys {
		result.WriteString(k)
		switch val := obj[k].(type) {
		case []interface{}:
			result.WriteString(arrayToString(val))
		case map[string]interface{}:
			result.WriteString(objectToString(val))
		default:
			result.WriteString(fmt.Sprint(val))
		}
	}
	return result.String()
}

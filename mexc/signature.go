package mexc

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strings"
)

func requestSignature(params map[string]string, privateKey string) string {
	params = removeEmptyValue(params)
	queryString := buildQueryString(params)

	h := hmac.New(sha256.New, []byte(privateKey))
	h.Write([]byte(queryString))
	signature := hex.EncodeToString(h.Sum(nil))

	return signature
}

func removeEmptyValue(params map[string]string) map[string]string {
	for k, v := range params {
		if v == "" {
			delete(params, k)
		}
	}
	return params
}

func buildQueryString(params map[string]string) string {
	var queryString []string
	for k, v := range params {
		queryString = append(queryString, url.QueryEscape(k)+"="+url.QueryEscape(v))
	}
	return strings.Join(queryString, "&")
}

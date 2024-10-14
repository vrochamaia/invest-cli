package cryptodotcom

import (
	"fmt"
	"investcli/http"
	"time"
)

func Accounts() {
	requestBody := signRequest(map[string]interface{}{
		"id":     10,
		"method": "private/user-balance",
		"params": map[string]interface{}{},
		"nonce":  time.Now().UnixMilli(),
	})

	response := http.Request(http.RequestInput{
		RequestHost:   "api.crypto.com",
		RequestMethod: "POST",
		RequestPath:   "/exchange/v1/private/user-balance",
		Body:          requestBody,
		Headers:       map[string]string{"Content-Type": "application/json"},
	})

	fmt.Println(response)
}

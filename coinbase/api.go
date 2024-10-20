package coinbase

import (
	"fmt"
	"investcli/http"
)

type GetRequestInput struct {
	RequestHost string
	RequestPath string
}

type authenticationTokenInput struct {
	requestHost   string
	requestPath   string
	requestMethod string
}

func authenticationToken(input authenticationTokenInput) string {

	uri := fmt.Sprintf("%s %s%s", input.requestMethod, input.requestHost, input.requestPath)

	jwtToken, _ := buildJWT(uri)

	return jwtToken
}

func Get(input GetRequestInput) string {
	requestMethod := "GET"

	token := authenticationToken(authenticationTokenInput{requestHost: input.RequestHost,
		requestPath: input.RequestPath, requestMethod: requestMethod})

	return http.Request(http.RequestInput{RequestMethod: requestMethod, RequestHost: input.RequestHost, RequestPath: input.RequestPath, Headers: map[string]string{"Authorization": "Bearer " + token}})
}

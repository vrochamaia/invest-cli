package coinbase

import (
	"fmt"
	"io"
	"net/http"
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
	const requestMethod = "GET"

	request, error := http.NewRequest(requestMethod, fmt.Sprintf("https://%s%s", input.RequestHost, input.RequestPath), nil)

	if error != nil {
		panic(error)
	}

	request.Header.Add("Authorization", "Bearer "+authenticationToken(authenticationTokenInput{requestHost: input.RequestHost,
		requestPath: input.RequestPath, requestMethod: requestMethod}))

	client := &http.Client{}

	response, error := client.Do(request)

	if error != nil {
		panic(error)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Sprintf("Failed to get response: %s", response.Status)
	}

	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)

	if error != nil {
		return fmt.Sprint("Failed to read the response: ", error)
	}

	return string(body)
}

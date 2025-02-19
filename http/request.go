package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type RequestInput struct {
	RequestHost   string
	RequestMethod string
	RequestPath   string
	Headers       map[string]string
	Body          string
}

func Request(input RequestInput) string {
	request, error := http.NewRequest(input.RequestMethod, fmt.Sprintf("https://%s%s", input.RequestHost, input.RequestPath), nil)

	if error != nil {
		panic(error)
	}

	if len(input.Headers) > 0 {
		for key, value := range input.Headers {
			request.Header.Add(key, value)
		}
	}

	if input.Body != "" {
		request.Body = io.NopCloser(strings.NewReader(input.Body))
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, error := client.Do(request)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		panic(fmt.Sprint("Failed to read the response: ", error))
	}

	if response.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Failed to get response: %s. Response Body: %s. Request Host: %s", response.Status, responseBody, input.RequestHost))
	}

	return string(responseBody)
}

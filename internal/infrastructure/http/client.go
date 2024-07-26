package http

import (
	"net/http"
)

type (
	HttpClientInterface interface {
		DoRequest(url string) int
	}

	HttpClient struct {
	}
)

func NewHttpClient() HttpClientInterface {
	return &HttpClient{}
}

func (c *HttpClient) DoRequest(url string) int {
	resp, err := http.Get(url)

	if err != nil {
		return 0
	}

	defer resp.Body.Close()

	return resp.StatusCode
}

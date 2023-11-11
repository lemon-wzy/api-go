package http_request

import (
	"net/http"
	"time"
)

type HttpClient struct {
	client *http.Client
}

func DefaultClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: 9 * time.Second,
		},
	}
}

func CustomTimeOutClient(timeout int64) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

func (client *HttpClient) Get() *HttpClient {
	return client
}

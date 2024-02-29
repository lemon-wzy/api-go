package http_request

import (
	"net/http"
	"time"
)

// HttpClient http客户端类型
type HttpClient struct {
	client *http.Client
}

// DefaultClient 默认客户端
func DefaultClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: 9 * time.Second,
		},
	}
}

// CustomTimeOutClient 自定义客户端的超时时间
func CustomTimeOutClient(timeout int64) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

// Get 获取客户端
// return *HttpClient
func (client *HttpClient) Get() *HttpClient {
	return client
}

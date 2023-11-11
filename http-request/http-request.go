package http_request

import (
	"api-go/common"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

func doPostClient(requestUrl string, requestData map[string]any, requestHeader map[string]string) any {
	defer common.RequestRecover()
	httpClient := http.DefaultClient
	jsonData, jsonErr := json.Marshal(requestData)
	if jsonErr != nil {
		panic(jsonErr)
	}
	request, reqErr := http.NewRequest(http.MethodPost, requestUrl, bytes.NewReader(jsonData))
	if reqErr != nil {
		panic(reqErr)
	}
	if requestHeader != nil {
		for k, v := range requestHeader {
			request.Header.Add(k, v)
		}
	}
	resp, repErr := httpClient.Do(request)
	if repErr != nil {
		panic(reqErr)
	}
	body, parseErr := io.ReadAll(resp.Body)
	if parseErr != nil {
		panic(parseErr)
	}
	return string(body)
}

func doGetClient(requireUrl string, requireParams map[string]string, requireHeader map[string]string) any {
	defer common.RequestRecover()

	httpClient := http.Client{
		Timeout: 9 * time.Second,
	}
	params := url.Values{}
	Url, err := url.Parse(requireUrl)
	if err != nil {
		panic(err)
	}
	if requireParams != nil {
		for k, v := range requireParams {
			params.Set(k, v)
		}
	}
	Url.RawQuery += params.Encode()
	urlPath := Url.String()
	req, e := http.NewRequest(http.MethodGet, urlPath, nil)
	if e != nil {
		panic(e)
	}
	if requireHeader != nil {
		for k, v := range requireParams {
			req.Header.Add(k, v)
		}
	}
	resp, repErr := httpClient.Do(req)
	if repErr != nil {
		panic(repErr)
	}
	body, parseErr := io.ReadAll(resp.Body)
	if parseErr != nil {
		panic(parseErr)
	}
	return string(body)
}

// DoRequest http请求
// url 完整的请求路径
// method 请求方式
// header 请求头
// params 请求参数
func DoRequest(url string, method string, header map[string]string, params map[string]any) any {
	switch method {
	case http.MethodGet:
		paramsStr := convert(params)
		return doGetClient(url, paramsStr, header)
	case http.MethodPost:
		return doPostClient(url, params, header)
	default:
		return "请指定请求方式!"
	}
}

// 参数转换
// params 传入参数
func convert(params map[string]any) map[string]string {
	result := make(map[string]string)
	for k, v := range params {
		result[k] = v.(string)
	}
	return result
}

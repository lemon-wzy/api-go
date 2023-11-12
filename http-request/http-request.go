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

// doPostClient
// url 完整的请求路径
// params 请求参数
// header 请求头
// httpClient *HttpClient
// update to custom timeout
func doPostClient(requestUrl string, requestData map[string]any, requestHeader map[string]string, httpClient *HttpClient) any {
	defer common.RequestRecover()
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
	//update to custom timeout
	resp, repErr := httpClient.client.Do(request)
	if repErr != nil {
		panic(reqErr)
	}
	body, parseErr := io.ReadAll(resp.Body)
	if parseErr != nil {
		panic(parseErr)
	}
	return string(body)
}

// doGetClient
// url 完整的请求路径
// params 请求参数
// header 请求头
// httpClient *HttpClient
// update to custom timeout
func doGetClient(requireUrl string, requireParams map[string]string, requireHeader map[string]string, httpClient *HttpClient) any {
	defer common.RequestRecover()
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
	//update to custom timeout
	resp, repErr := httpClient.client.Do(req)
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
// httpClient *HttpClient
func DoRequest(url string, method string, header map[string]string, params map[string]any, httpClient *HttpClient) any {
	switch method {
	case http.MethodGet:
		paramsStr := convert(params)
		//update to custom timeout
		return doGetClient(url, paramsStr, header, httpClient)
	case http.MethodPost:
		//update to custom timeout
		return doPostClient(url, params, header, httpClient)
	default:
		return "请指定请求方式!"
	}
}

// DoRequestByConsoleArgs
// consoleArgs 命令行参数
// httpClient *HttpClient 请求客户端
func DoRequestByConsoleArgs(consoleArgs *common.ConsoleArgs, client *HttpClient) any {
	if client == nil {
		client = &HttpClient{
			client: &http.Client{
				Timeout: 9 * time.Second,
			},
		}
	}
	switch consoleArgs.M {
	case http.MethodGet:
		paramStr := convert(consoleArgs.P)
		return doGetClient(consoleArgs.U, paramStr, consoleArgs.H, client)
	case http.MethodPost:
		return doPostClient(consoleArgs.U, consoleArgs.P, consoleArgs.H, client)
	default:
		return "请指定请求方式!"
	}
}

// convert 参数转换
// params 传入参数
// return 转换后的参数
func convert(params map[string]any) map[string]string {
	result := make(map[string]string)
	for k, v := range params {
		result[k] = v.(string)
	}
	return result
}

package parse

import (
	"api-go/common"
	"encoding/json"
)

//自定义命令行解析类型

type paramValue map[string]any

// NewParamValue 新建命令行参数
// value 命令行参数
// p 命令行参数指针
// return 命令行参数
func NewParamValue(value map[string]any, p *map[string]any) *paramValue {
	*p = value
	return (*paramValue)(p)
}

// String 命令行参数转json字符串
// return 命令行参数json字符串
func (p *paramValue) String() string {
	if p == nil {
		return ""
	}
	defer common.ParseRecover()
	s, _ := json.Marshal(*p)
	dataStr := string(s)
	return dataStr
}

// Set 设置命令行参数
// value 命令行参数
// return error
func (p *paramValue) Set(value string) error {
	if value == "" {
		return nil
	}
	defer common.ParseRecover()
	err := json.Unmarshal([]byte(value), &p)
	if err != nil {
		panic(err)
	}
	return nil
}

// headerValue 请求头类型
type headerValue map[string]string

// NewHeaderValue 新建请求头
// value 请求头
// p 请求头指针
// return 请求头
func NewHeaderValue(value map[string]string, p *map[string]string) *headerValue {
	*p = value
	return (*headerValue)(p)
}

// String 命令行参数转json字符串
// return 命令行参数json字符串
func (h *headerValue) String() string {
	if h == nil {
		return ""
	}
	defer common.ParseRecover()
	s, _ := json.Marshal(*h)
	dataStr := string(s)
	return dataStr
}

// Set 设置命令行参数
// value 命令行参数
// return error
func (h *headerValue) Set(s string) error {
	if s == "" {
		return nil
	}
	defer common.ParseRecover()
	err := json.Unmarshal([]byte(s), &h)
	if err != nil {
		panic(err)
	}
	return nil
}

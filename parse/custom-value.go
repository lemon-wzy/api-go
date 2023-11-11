package parse

import (
	"api-go/common"
	"encoding/json"
)

//自定义命令行解析类型

type paramValue map[string]any

func NewParamValue(value map[string]any, p *map[string]any) *paramValue {
	*p = value
	return (*paramValue)(p)
}

func (p *paramValue) String() string {
	if p == nil {
		return ""
	}
	defer common.ParseRecover()
	s, _ := json.Marshal(*p)
	dataStr := string(s)
	return dataStr
}
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

type headerValue map[string]string

func NewHeaderValue(value map[string]string, p *map[string]string) *headerValue {
	*p = value
	return (*headerValue)(p)
}

func (h *headerValue) String() string {
	if h == nil {
		return ""
	}
	defer common.ParseRecover()
	s, _ := json.Marshal(*h)
	dataStr := string(s)
	return dataStr
}

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

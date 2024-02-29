package common

import (
	"fmt"
	"os"
)

// RequestRecover 自定义请求错误
func RequestRecover() {
	r := recover()
	if r != nil {
		fmt.Println("请求失败:", r)
		os.Exit(1)
	}
}

// ParseRecover 自定义解析错误
func ParseRecover() {
	r := recover()
	if r != nil {
		fmt.Println("参数解析失败:", r)
		os.Exit(1)
	}
}

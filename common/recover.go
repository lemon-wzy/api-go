package common

import (
	"fmt"
	"os"
)

func RequestRecover() {
	r := recover()
	if r != nil {
		fmt.Println("请求失败:", r)
		os.Exit(1)
	}
}

func ParseRecover() {
	r := recover()
	if r != nil {
		fmt.Println("参数解析失败:", r)
		os.Exit(1)
	}
}

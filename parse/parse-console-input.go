package parse

import (
	http_request "api-go/http-request"
	"flag"
	"fmt"
	"os"
)

// 用于解析命令行输入参数
var (
	//请求方式
	m string
	//请求路径
	u string
	//请求参数
	p map[string]any

	//请求头
	H map[string]string

	//帮助
	help, h bool
)
var flagSet = flag.NewFlagSet("custom flag", flag.ExitOnError)

func init() {
	flagSet.StringVar(&u, "u", "", "request url")
	flagSet.StringVar(&m, "m", "", "request method")
	flagSet.Var(NewHeaderValue(map[string]string{}, &H), "H", "request head")
	flagSet.Var(NewParamValue(map[string]any{}, &p), "p", "request params")

	flagSet.BoolVar(&h, "h", false, "help")
	flagSet.BoolVar(&help, "help", false, "help")
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
func Init() {
	if h || help {
		flagSet.Usage()
		os.Exit(0)
	}
	//call request and print response
	json := http_request.DoRequest(u, m, H, p)
	fmt.Println(json, "\n\r")
}

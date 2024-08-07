package parse

import (
	"api-go/common"
	http_request "api-go/http-request"
	"flag"
	"fmt"
	"os"
)

var consoleArgs = &common.ConsoleArgs{}
var flagSet = flag.NewFlagSet("custom flag", flag.ExitOnError)

// 初始化解析命令行参数
func init() {
	flagSet.StringVar(&consoleArgs.U, "u", "", "request url ")
	flagSet.StringVar(&consoleArgs.M, "m", "GET", "request method GET or POST")
	flagSet.Int64Var(&consoleArgs.T, "t", 9, "request timeout unit second")
	flagSet.Var(NewHeaderValue(map[string]string{}, &consoleArgs.H), "H", "request header")
	flagSet.Var(NewParamValue(map[string]any{}, &consoleArgs.P), "p", "request params")
	flagSet.BoolVar(&consoleArgs.HELP, "h", false, "command help")
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}

// Init 初始化
func Init() {
	if consoleArgs.HELP || consoleArgs.U == "" {
		flagSet.Usage()
		os.Exit(0)
	}
	var result any
	if consoleArgs.T != 9 {
		result = http_request.DoRequestByConsoleArgs(consoleArgs, http_request.CustomTimeOutClient(consoleArgs.T))
	} else {
		//call request and print response
		result = http_request.DoRequestByConsoleArgs(consoleArgs, http_request.DefaultClient())
	}
	fmt.Println(result, "\n\r")
}

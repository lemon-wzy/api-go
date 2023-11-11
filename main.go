package main

import "api-go/parse"

/*
*
用于解析命令行输入参数
*/
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

//
//func init() {
//	flag.StringVar(&u, "u", "", "request url")
//	flag.StringVar(&m, "m", "", "request method")
//	flag.Var(parse.NewHeaderValue(map[string]string{}, &H), "H", "request head")
//	flag.Var(parse.NewParamValue(map[string]any{}, &p), "p", "request params")
//	flag.BoolVar(&h, "h", false, "help")
//	flag.BoolVar(&help, "help", false, "help")
//}

func main() {
	//flag.Parse()
	//fmt.Printf("method：%v\nparams：%v\nurl：%v\nhelp：%v\nheader:%v\n",
	//	m, p, u, h, H)
	//if h || help {
	//	flag.Usage()
	//}
	parse.Init()
}

package common

// ConsoleArgs 命令行参数
// M 请求方式
// U 请求地址
// P 请求参数
// H 请求头
// T 请求超时
// HELP 帮助
type ConsoleArgs struct {
	M    string
	U    string
	P    map[string]any
	H    map[string]string
	T    int64
	HELP bool
}

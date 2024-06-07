package main

import "C"

import (
	"fmt"
	"log/slog"
)

//export SayHello
func SayHello(name *C.char) *C.char {
	// 将 C语言中 的字符串转成 go字符串
	goName := C.GoString(name)
	// 执行 go方法
	goStr := sayHelloProxy(goName)
	// 将 go字符串 转成 C语言字符串
	return C.CString(goStr)
}

func sayHelloProxy(name string) string {
	slog.Info("message from ChineseGreeter.SayHello", slog.String("name", name))
	str := fmt.Sprintf("Chinese : 你好 %v ... !", name)
	return str
}

func main() {
}

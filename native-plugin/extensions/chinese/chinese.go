package main

import "C"

import (
	"fmt"
	"log/slog"
)

//export SayHello
func SayHello(name *C.char) {
	goName := C.GoString(name)
	slog.Info("message from ChineseGreeter.SayHello", slog.String("name", goName))
	str := fmt.Sprintf("Chinese : Hello %v ... !", goName)
	fmt.Println(str)
	// return C.CString(C.GoString(str))
}

func main() {
}

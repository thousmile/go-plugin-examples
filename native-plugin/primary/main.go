package main

//#cgo CFLAGS: -I./
//#cgo LDFLAGS: -L. -lplugins
////#cgo LDFLAGS: -L${SRCDIR}/plugins -lplugins
//
//#include "libplugins.h"
import "C"

func main() {
	// .dll 或者 .so 动态库必须在 同级目录下。否则调用失败。
	// ${SRCDIR} 没用
	C.SayHello(C.CString("Hello, World"))
	/*	cStr := C.SayHello(C.CString("Hello, World"))
		goStr := C.GoString(cStr)
		fmt.Println(goStr)*/
}

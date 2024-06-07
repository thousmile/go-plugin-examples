package main

//#cgo CFLAGS: -I./
//#cgo LDFLAGS: -L. -lplugins
////#cgo LDFLAGS: -L${SRCDIR}/plugins -lplugins
//
//#include "libplugins.h"
import "C"
import (
	"log/slog"
)

func main() {
	// .dll 或者 .so 动态库必须在 同级目录下。否则调用失败。
	// ${SRCDIR} 没用
	cStr := C.SayHello(C.CString("tom"))
	goStr := C.GoString(cStr)
	slog.Info("SayHello", slog.String("result", goStr))
}

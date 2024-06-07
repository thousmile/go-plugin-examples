package main

import (
	"fmt"
	"log/slog"
)

var Impl ChineseGreeter

type ChineseGreeter struct {
}

func (m ChineseGreeter) SayHello(name string) string {
	slog.Info("message from ChineseGreeter.SayHello")
	return fmt.Sprintf("Chinese : 你好 %s ... !", name)
}

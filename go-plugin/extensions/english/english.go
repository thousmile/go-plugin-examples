package main

import (
	"fmt"
	"log/slog"
)

var Impl EnglishGreeter

type EnglishGreeter struct {
}

func (m EnglishGreeter) SayHello(name string) string {
	slog.Info("message from EnglishGreeter.SayHello")
	return fmt.Sprintf("English : Hello %s ... !", name)
}

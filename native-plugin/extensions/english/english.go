package main

import (
	"fmt"
	"log/slog"
)

//export SayHello
func SayHello(name string) {
	slog.Info("message from EnglishGreeter.SayHello")
	str := fmt.Sprintf("English : Hello %s ... !", name)
	return str
}

func main() {
}

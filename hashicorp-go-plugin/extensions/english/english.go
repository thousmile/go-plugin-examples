package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"go-plugin-demo1/shared"
	"os"
)

type EnglishGreeter struct {
	hclog.Logger
}

func (m EnglishGreeter) SayHello(name string) string {
	m.Info("message from EnglishGreeter.SayHello")
	return fmt.Sprintf("English : Hello %s ... !", name)
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "English Plugin",
		Level:      hclog.Info,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	greeter := &EnglishGreeter{Logger: logger}
	var pluginMap = map[string]plugin.Plugin{
		shared.GreeterApi: &shared.GreeterPlugin{Impl: greeter},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"go-plugin-demo1/shared"
	"os"
)

type ChineseGreeter struct {
	hclog.Logger
}

func (m ChineseGreeter) SayHello(name string) string {
	m.Info("message from ChineseGreeter.SayHello")
	return fmt.Sprintf("Chinese : 你好 %s ... !", name)
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "Chinese Plugin",
		Level:      hclog.Info,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	greeter := &ChineseGreeter{Logger: logger}
	var pluginMap = map[string]plugin.Plugin{
		shared.GreeterApi: &shared.GreeterPlugin{Impl: greeter},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

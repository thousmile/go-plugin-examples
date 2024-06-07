package main

import (
	"flag"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"go-plugin-demo1/shared"
	"log"
	"os"
	"os/exec"
)

func main() {
	var pluginName string
	flag.StringVar(&pluginName, "plugin", "../plugins/english.exe", "plugin name")
	flag.Parse()
	logger := hclog.New(&hclog.LoggerOptions{
		Output: os.Stdout,
		Level:  hclog.Info,
	})
	var pluginMap = map[string]plugin.Plugin{
		shared.GreeterApi: &shared.GreeterPlugin{},
	}
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(pluginName),
		Logger:          logger,
	})
	defer client.Kill()
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}
	raw, err := rpcClient.Dispense(shared.GreeterApi)
	if err != nil {
		log.Fatal(err)
	}
	greeter := raw.(shared.Greeter)
	resp := greeter.SayHello("tom")
	log.Println(resp)
}

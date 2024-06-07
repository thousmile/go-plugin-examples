package shared

import (
	"github.com/hashicorp/go-plugin"
	_ "github.com/hashicorp/go-plugin"
	"net/rpc"
)

type Greeter interface {
	SayHello(name string) string
}

type GreeterRPC struct{ client *rpc.Client }

func (g *GreeterRPC) SayHello(name string) string {
	var resp string
	err := g.client.Call("Plugin.SayHello", name, &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

type GreeterRPCServer struct {
	Impl Greeter
}

func (s *GreeterRPCServer) SayHello(args string, resp *string) error {
	*resp = s.Impl.SayHello(args)
	return nil
}

type GreeterPlugin struct {
	// Impl Injection
	Impl Greeter
}

func (p GreeterPlugin) Server(broker *plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: p.Impl}, nil
}

func (GreeterPlugin) Client(broker *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}

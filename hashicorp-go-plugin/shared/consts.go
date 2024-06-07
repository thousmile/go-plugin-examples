package shared

import "github.com/hashicorp/go-plugin"

var (
	HandshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "BASIC_PLUGIN",
		MagicCookieValue: "hashicorp-go-plugin",
	}
)

const (
	GreeterApi = "greeter"
)

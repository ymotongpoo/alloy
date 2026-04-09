package receiver

import (
	fnet "github.com/grafana/alloy/internal/component/common/net"
	"github.com/grafana/alloy/internal/component/sigil"
)

type Arguments struct {
	Server    *fnet.ServerConfig          `alloy:",squash"`
	ForwardTo []sigil.GenerationsReceiver `alloy:"forward_to,attr"`
}

// SetToDefault implements syntax.Defaulter.
func (a *Arguments) SetToDefault() {
	*a = Arguments{
		Server: fnet.DefaultServerConfig(),
	}
}

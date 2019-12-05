package centrifuge

import (
	"ChainBridgeV2/core"
	"ChainBridgeV2/types"
)

var _ core.Listener = &Listener{}

type Listener struct {
	conn core.Connection
}

func NewListener(conn core.Connection) *Listener {
	return &Listener{conn: conn}
}

func (l *Listener) RegisterEventHandler(name types.EventName, handler func()) {
	panic("not implemented")
}

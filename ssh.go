// Package ssh contains an Endpoint implementation for the Git SSH protocol.
package ssh

import (
	"context"
	"fmt"
	"github.com/mughub/mughub/bare"
	"github.com/spf13/viper"
	"net"
)

type endpoint struct{}

func (e *endpoint) ListenAndServe(ctx context.Context) error {
	return nil
}

// NewEndpoint returns an Endpoint implementation which serves
// the Git SSH protocol.
//
func NewEndpoint() bare.Endpoint {
	return &endpoint{}
}

func getTCPListener(cfg *viper.Viper) net.Listener {
	addr := fmt.Sprintf("%s:%d", cfg.GetString("addr"), cfg.GetInt("port"))
	l, err := net.Listen("tcp", addr)
	if err != nil {
		l.Close()
		panic(err)
	}
	return l
}

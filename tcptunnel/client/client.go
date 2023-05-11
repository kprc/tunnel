package client

import (
	"net"
)

type ITunnelClient interface {
	Connect() (net.Conn, error)
	Close() error
}

type TunnelClient struct {
	pool   int
	limits int
	remote string
	//             pool id
	realConns map[int]net.Conn
	//             pool     tunnel-conn
	tunnelConns  map[int]map[int]net.Conn
	remoteServer string
}

func NewTunnelClient(pool int, limits int, remote string) *TunnelClient {
	return &TunnelClient{
		pool:   pool,
		limits: limits,
		remote: remote,
	}
}

func (tc *TunnelClient) Connect() (net.Conn, error) {
	return nil, nil
}

func (tc *TunnelClient) Close() error {
	return nil
}

package conn

import (
	"time"
)

type TunnelConn struct {
	real   int
	tunnel int
}

func (tc *TunnelConn) Read(b []byte) (n int, err error) {
	return 0, nil
}

func (tc *TunnelConn) Write(b []byte) (n int, err error) {
	return 0, nil
}

func (tc *TunnelConn) Close() error {
	return nil
}

func (tc *TunnelConn) LocalAddr() Addr {
	return nil
}

func (tc *TunnelConn) RemoteAddr() Addr {
	return nil
}

func (tc *TunnelConn) SetDeadline(t time.Time) error {
	return nil
}

func (tc *TunnelConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (tc *TunnelConn) SetWriteDeadline(t time.Time) error {
	return nil
}

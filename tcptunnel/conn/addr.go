package conn

type Addr interface {
	Network() string // name of the network (for example, "tcp", "udp")
	String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}

type TunnelAddr struct {
}

func (ta *TunnelAddr) Network() string {
	return ""
}

func (ta *TunnelAddr) String() string {
	return ""
}

package grpcserver

import (
	"net"
)

// TODO: can be generated
// Option -.
type Option func(*Server)

// Port -.
func Port(port string) Option {
	return func(s *Server) {
		s.address = net.JoinHostPort("", port)
	}
}

package server

import (
	"github.com/bborbe/ip/handler"
	"github.com/bborbe/server"
)

func NewServer(port int) server.Server {
	return server.NewServerPort(port, handler.New())
}

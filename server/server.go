package server

import (
	"github.com/bborbe/server"
	"github.com/bborbe/password/generator"
"github.com/bborbe/password/handler"
)

func NewServer(port int, passwordGenerator generator.PasswordGenerator) server.Server {
	return server.NewServerPort(port, handler.New(passwordGenerator))
}

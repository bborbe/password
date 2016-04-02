package server

import (
	"github.com/bborbe/password/generator"
	"github.com/bborbe/password/handler"
	"github.com/bborbe/server"
)

func NewServer(port int, passwordGenerator generator.PasswordGenerator) server.Server {
	return server.NewServerPort(port, handler.New(passwordGenerator))
}

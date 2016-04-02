package main

import (
	"flag"

	"github.com/bborbe/log"
	password_generator "github.com/bborbe/password/generator"
	password_server "github.com/bborbe/password/server"
)

var logger = log.DefaultLogger

const (
	DEFAULT_PORT       int = 8080
	PARAMETER_LOGLEVEL     = "loglevel"
	PARAMETER_PORT         = "port"
)

func main() {
	logLevelPtr := flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")
	portPtr := flag.Int(PARAMETER_PORT, DEFAULT_PORT, "port")
	flag.Parse()
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Tracef("set log level to %s", *logLevelPtr)
	passwordGenerator := password_generator.New()
	srv := password_server.NewServer(*portPtr, passwordGenerator)
	srv.Run()
}

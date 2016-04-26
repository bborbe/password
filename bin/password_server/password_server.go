package main

import (
	"flag"

	"fmt"
	"net/http"
	"os"

	"github.com/bborbe/log"
	password_generator "github.com/bborbe/password/generator"
	"github.com/bborbe/password/handler"
	"github.com/facebookgo/grace/gracehttp"
)

var logger = log.DefaultLogger

const (
	DEFAULT_PORT       int = 8080
	PARAMETER_LOGLEVEL     = "loglevel"
	PARAMETER_PORT         = "port"
)

var (
	logLevelPtr = flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")
	portPtr     = flag.Int(PARAMETER_PORT, DEFAULT_PORT, "port")
)

func main() {
	flag.Parse()

	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	server, err := createServer(*portPtr)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	gracehttp.Serve(server)
}

func createServer(port int) (*http.Server, error) {
	passwordGenerator := password_generator.New()
	h := handler.New(passwordGenerator)
	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: h}, nil
}

package main

import (
	"fmt"
	"net/http"
	"os"

	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/log"
	password_generator "github.com/bborbe/password/generator"
	"github.com/bborbe/password/handler"
	"github.com/facebookgo/grace/gracehttp"
	"runtime"
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
	defer logger.Close()
	flag.Parse()

	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	runtime.GOMAXPROCS(runtime.NumCPU())

	server, err := createServer(*portPtr)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
	logger.Debugf("start server")
	gracehttp.Serve(server)
}

func createServer(port int) (*http.Server, error) {
	if port <= 0 {
		return nil, fmt.Errorf("parameter %s invalid", PARAMETER_PORT)
	}
	passwordGenerator := password_generator.New()
	h := handler.New(passwordGenerator)
	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: h}, nil
}

package main

import (
	"fmt"
	"net/http"
	"runtime"

	debug_handler "github.com/bborbe/http_handler/debug"

	flag "github.com/bborbe/flagenv"
	password_generator "github.com/bborbe/password/generator"
	"github.com/bborbe/password/handler"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
)

const (
	DEFAULT_PORT   int = 8080
	PARAMETER_PORT     = "port"
)

var (
	portPtr = flag.Int(PARAMETER_PORT, DEFAULT_PORT, "port")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := do(
		*portPtr,
	)
	if err != nil {
		glog.Exit(err)
	}
}

func do(
	port int,
) error {
	server, err := createServer(
		port,
	)
	if err != nil {
		return err
	}
	glog.V(2).Infof("start server")
	return gracehttp.Serve(server)
}

func createServer(
	port int,
) (*http.Server, error) {
	if port <= 0 {
		return nil, fmt.Errorf("parameter %s invalid", PARAMETER_PORT)
	}
	passwordGenerator := password_generator.New()
	handler := handler.New(passwordGenerator)

	if glog.V(4) {
		handler = debug_handler.New(handler)
	}

	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler}, nil
}

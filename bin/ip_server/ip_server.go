package main

import (
	"fmt"
	"net/http"

	debug_handler "github.com/bborbe/http_handler/debug"

	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/ip/handler"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
)

const (
	PARAMETER_PORT      = "port"
	DEFAULT_PORT    int = 8080
	PARAMETER_DEBUG     = "debug"
)

var (
	portPtr  = flag.Int(PARAMETER_PORT, DEFAULT_PORT, "Port")
	debugPtr = flag.Bool(PARAMETER_DEBUG, false, "debug")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := do(
		*portPtr,
		*debugPtr,
	)
	if err != nil {
		glog.Exit(err)
	}

}

func do(
	port int,
	debug bool,
) error {
	server, err := createServer(
		port,
		debug,
	)
	if err != nil {
		return err
	}
	glog.V(2).Infof("start server")
	return gracehttp.Serve(server)
}

func createServer(
	port int,
	debug bool,
) (*http.Server, error) {
	handler := handler.New()

	if debug {
		handler = debug_handler.New(handler)
	}

	return &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler}, nil
}

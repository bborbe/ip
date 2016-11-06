package main

import (
	"net/http"

	debug_handler "github.com/bborbe/http_handler/debug"

	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/ip/handler"
	"github.com/bborbe/ip/model"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
)

const (
	PARAMETER_PORT     = "port"
	DEFAULT_PORT   int = 8080
)

var (
	portPtr = flag.Int(PARAMETER_PORT, DEFAULT_PORT, "Port")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := do(); err != nil {
		glog.Exit(err)
	}

}

func do() error {
	server, err := createServer()
	if err != nil {
		return err
	}
	glog.V(2).Infof("start server")
	return gracehttp.Serve(server)
}

func createServer() (*http.Server, error) {
	port := model.Port(*portPtr)
	handler := handler.New()
	if glog.V(4) {
		handler = debug_handler.New(handler)
	}
	glog.V(2).Infof("create http server on %s", port.Address())
	return &http.Server{Addr: port.Address(), Handler: handler}, nil
}

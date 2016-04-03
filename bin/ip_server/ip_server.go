package main

import (
	ip_server "github.com/bborbe/ip/server"
	"github.com/bborbe/log"
	flag  "github.com/bborbe/flagenv"
)

var logger = log.DefaultLogger

const (
	DEFAULT_PORT int = 8080
	PARAMETER_LOGLEVEL = "loglevel"
	PARAMETER_PORT = "port"
)

func main() {
	logLevelPtr := flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")
	portPtr := flag.Int(PARAMETER_PORT, DEFAULT_PORT, "port")
	flag.Parse()
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Tracef("set log level to %s", *logLevelPtr)
	logger.Infof("listen on port %d", *portPtr)
	srv := ip_server.NewServer(*portPtr)
	srv.Run()
}

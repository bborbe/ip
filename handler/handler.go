package handler

import (
	"fmt"
	"net/http"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type statusHandler struct {
}

func New() http.Handler {
	s := new(statusHandler)
	return s
}

func (s *statusHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	logger.Debugf("get ip")
	ip := getIp(request)
	responseWriter.Header().Add("Content-Type", "text/plain")
	responseWriter.WriteHeader(200)
	fmt.Fprint(responseWriter, ip)
}

func getIp(request *http.Request) string {
	logger.Tracef("header %v", request.Header)
	forwardedAddr := request.Header.Get("HTTP_X_FORWARDED_FOR")
	if len(forwardedAddr) != 0 {
		return forwardedAddr
	}
	remoteAddr := request.Header.Get("REMOTE_ADDR")
	if len(remoteAddr) != 0 {
		return remoteAddr
	}
	return request.RemoteAddr
}

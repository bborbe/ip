package handler

import (
	"fmt"
	"net/http"

	"github.com/bborbe/log"
	"strings"
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
	ip, err := getIp(request)
	if err != nil {
		responseWriter.WriteHeader(500)
		fmt.Fprint(responseWriter, "Internal Server Error: %v", err)
		return
	}
	responseWriter.Header().Add("Content-Type", "text/plain")
	responseWriter.WriteHeader(200)
	fmt.Fprint(responseWriter, ip)
}

func getIp(request *http.Request) (string, error) {
	logger.Tracef("header %v", request.Header)
	forwardedAddr := request.Header.Get("HTTP_X_FORWARDED_FOR")
	if len(forwardedAddr) != 0 {
		return forwardedAddr, nil
	}
	remoteAddr := request.Header.Get("REMOTE_ADDR")
	if len(remoteAddr) != 0 {
		return remoteAddr, nil
	}
	parts := strings.Split(request.RemoteAddr, ":")
	if len(parts) > 0 && len(parts[0]) > 0 {
		return parts[0], nil
	}
	return "", fmt.Errorf("remote ip not found")
}

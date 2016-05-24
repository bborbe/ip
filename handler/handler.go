package handler

import (
	"fmt"
	"net/http"

	"strings"

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
	ip, err := getIp(request)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(responseWriter, "Internal Server Error: %v", err)
		return
	}
	responseWriter.Header().Add("Content-Type", "text/plain")
	responseWriter.WriteHeader(http.StatusOK)
	fmt.Fprint(responseWriter, ip)
}

func getIp(request *http.Request) (string, error) {
	logger.Tracef("header %v", request.Header)
	forwardedAddr := getHeader(request, "X-Forwarded-For")
	logger.Tracef("header X-Forwarded-For %s", forwardedAddr)
	if len(forwardedAddr) != 0 {
		return forwardedAddr, nil
	}
	remoteAddr := getHeader(request, "X-Remote-Addr")
	logger.Tracef("header X-Remote-Addr %s", remoteAddr)
	if len(remoteAddr) != 0 {
		return remoteAddr, nil
	}
	parts := strings.Split(request.RemoteAddr, ":")
	logger.Tracef("remoteAddr %s", request.RemoteAddr)
	if len(parts) > 0 && len(parts[0]) > 0 {
		return parts[0], nil
	}
	return "", fmt.Errorf("remote ip not found")
}

func getHeader(request *http.Request, name string) string {
	result := request.Header.Get(name)
	if len(result) > 0 {
		return result
	}
	return request.Header.Get(parameterNameToEnvName(name))
}

func parameterNameToEnvName(name string) string {
	return strings.Replace(strings.ToUpper(name), "-", "_", -1)
}

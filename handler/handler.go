package handler

import (
	"fmt"
	"net/http"

	"strings"

	"github.com/golang/glog"
)

type statusHandler struct {
}

func New() http.Handler {
	s := new(statusHandler)
	return s
}

func (s *statusHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	var ip string
	var err error
	glog.V(4).Infof("get ip")
	if ip, err = getIp(request); err != nil {
		glog.Warningf("get ip failed: %v", err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(responseWriter, "Internal Server Error: %v", err)
		return
	}
	responseWriter.Header().Add("Content-Type", "text/plain")
	responseWriter.WriteHeader(http.StatusOK)
	glog.V(2).Infof("return ip %s to client", ip)
	fmt.Fprint(responseWriter, ip)
}

func getIp(request *http.Request) (string, error) {
	glog.V(4).Infof("header %v", request.Header)
	address := getAddress(request, "X-Forwarded-For", "X-Remote-Addr")
	return parseIpFromAddress(address)
}

func parseIpFromAddress(address string) (string, error) {
	glog.V(4).Infof("parse ip from address %s", address)
	if len(address) == 0 {
		return "", fmt.Errorf("remote ip not found")
	}
	parts := strings.Split(address, ", ")
	if len(parts) == 0 || len(parts[0]) == 0 {
		return "", fmt.Errorf("remote ip not found")
	}
	parts = strings.Split(parts[0], ":")
	if len(parts) == 0 || len(parts[0]) == 0 {
		return "", fmt.Errorf("remote ip not found")
	}
	return parts[0], nil
}

func getAddress(request *http.Request, names ...string) string {
	address := getHeaders(request, names...)
	if len(address) > 0 {
		return address
	}
	return request.RemoteAddr
}

func getHeaders(request *http.Request, names ...string) string {
	var result string
	for _, name := range names {
		if result = getHeader(request, name); len(result) > 0 {
			return result
		}
		if result = getHeader(request, parameterNameToEnvName(name)); len(result) > 0 {
			return result
		}
	}
	return ""
}

func getHeader(request *http.Request, name string) string {
	var result string
	result = request.Header.Get(name)
	glog.V(4).Infof("get header %s => %s", name, result)
	return result
}

func parameterNameToEnvName(name string) string {
	return strings.Replace(strings.ToUpper(name), "-", "_", -1)
}

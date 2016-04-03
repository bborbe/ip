package handler

import (
	"net/http"
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsHandler(t *testing.T) {
	object := New()
	var expected *http.Handler
	err := AssertThat(object, Implements(expected))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetIpHeaderRemoteAddr(t *testing.T) {
	h := http.Request{Header: http.Header{}}
	h.Header.Add("REMOTE_ADDR", "192.168.1.1")
	ip := getIp(&h)
	if err := AssertThat(ip, Is("192.168.1.1")); err != nil {
		t.Fatal(err)
	}
}

func TestGetIpHeaderForwardedAddr(t *testing.T) {
	h := http.Request{Header: http.Header{}}
	h.Header.Add("HTTP_X_FORWARDED_FOR", "192.168.1.1")
	ip := getIp(&h)
	if err := AssertThat(ip, Is("192.168.1.1")); err != nil {
		t.Fatal(err)
	}
}

func TestGetIpRemoteAddr(t *testing.T) {
	h := http.Request{Header: http.Header{}}
	h.RemoteAddr = "192.168.1.1"
	ip := getIp(&h)
	if err := AssertThat(ip, Is("192.168.1.1")); err != nil {
		t.Fatal(err)
	}
}

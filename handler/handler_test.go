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
	h.Header.Add("X-Remote-Addr", "192.168.1.1")
	ip, err := getIp(&h)
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(ip, Is("192.168.1.1")); err != nil {
		t.Fatal(err)
	}
}

func TestGetIpHeaderRemoteAddrUpper(t *testing.T) {
	h := http.Request{Header: http.Header{}}
	h.Header.Add("X_REMOTE_ADDR", "192.168.1.1")
	ip, err := getIp(&h)
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(ip, Is("192.168.1.1")); err != nil {
		t.Fatal(err)
	}
}

func TestGetIpHeaderForwardedAddr(t *testing.T) {
	h := http.Request{Header: http.Header{}}
	h.Header.Add("X-Forwarded-For", "192.168.1.1")
	ip, err := getIp(&h)
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(ip, Is("192.168.1.1")); err != nil {
		t.Fatal(err)
	}
}

func TestGetIpHeaderForwardedAddrUpper(t *testing.T) {
	h := http.Request{Header: http.Header{}}
	h.Header.Add("X_FORWARDED_FOR", "192.168.1.1")
	ip, err := getIp(&h)
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(ip, Is("192.168.1.1")); err != nil {
		t.Fatal(err)
	}
}

func TestGetIpRemoteAddr(t *testing.T) {
	h := http.Request{Header: http.Header{}}
	h.RemoteAddr = "192.168.1.1"
	ip, err := getIp(&h)
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(ip, Is("192.168.1.1")); err != nil {
		t.Fatal(err)
	}
}

func TestGetIpFail(t *testing.T) {
	h := http.Request{Header: http.Header{}}
	_, err := getIp(&h)
	if err := AssertThat(err, NotNilValue()); err != nil {
		t.Fatal(err)
	}
}

func TestGetIpHeaderForwardedAddrMulti(t *testing.T) {
	h := http.Request{Header: http.Header{}}
	h.Header.Add("X-Forwarded-For", "80.128.80.111, 10.102.95.1")
	ip, err := getIp(&h)
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(ip, Is("80.128.80.111")); err != nil {
		t.Fatal(err)
	}
}

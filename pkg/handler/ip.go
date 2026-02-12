// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ip_requests_total",
			Help: "Total number of IP requests",
		},
		[]string{"status"},
	)
)

// IPHandler returns the client's IP address.
type IPHandler struct{}

// NewIPHandler creates a new IP handler.
func NewIPHandler() *IPHandler {
	return &IPHandler{}
}

func (h *IPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip := getClientIP(r)

	requestsTotal.WithLabelValues("success").Inc()
	glog.V(2).Infof("return ip %s to client", ip)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, ip)
}

// getClientIP extracts the client IP from the request.
// It checks X-Forwarded-For and X-Real-IP headers first (for proxied requests),
// then falls back to RemoteAddr.
func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header (can contain multiple IPs)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// Take the first IP (original client)
		if idx := strings.Index(xff, ","); idx != -1 {
			xff = strings.TrimSpace(xff[:idx])
		}
		if xff != "" {
			return xff
		}
	}

	// Check X-Real-IP header
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Check X-Remote-Addr header (legacy support)
	if xra := r.Header.Get("X-Remote-Addr"); xra != "" {
		return xra
	}

	// Fall back to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// RemoteAddr might not have a port
		return r.RemoteAddr
	}
	return ip
}

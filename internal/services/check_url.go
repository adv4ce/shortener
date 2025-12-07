package services

import (
	"net/url"
	"strings"
)

func IsValidUrl(urlStr string) bool {
	if len(urlStr) > 2048 {
		return false
	}

	parsed, err := url.ParseRequestURI(urlStr)

	if err != nil {
		return false
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return false
	}

	hostname := parsed.Hostname()
	if hostname == "localhost" || hostname == "127.0.0.1" || strings.HasPrefix(hostname, "192.168.") ||
		strings.HasPrefix(hostname, "10.") ||
		strings.HasPrefix(hostname, "172.16.") {
		return false
	}

	return true
}

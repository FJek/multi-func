package util

import (
	"net"
	"net/http"
	"strings"
)

func GetIp(w http.ResponseWriter, r *http.Request) (ip string) {
	ip = ``
	// X-forwarded-for
	xForwardedFor := r.Header.Get(`X-forwarded-for`)
	if xForwardedFor != `` {
		ip = strings.TrimSpace(strings.Split(xForwardedFor, `,`)[0])
		return
	}
	// X-real-Ip
	xRealIp := r.Header.Get(`X-real-Ip`)
	if xRealIp != `` {
		ip = strings.TrimSpace(xRealIp)
		return
	}
	// Remote-Addr
	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return ``
	}
	return
}

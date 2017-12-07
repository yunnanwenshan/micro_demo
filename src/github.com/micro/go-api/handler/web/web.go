// Package web contains the web handler including websocket support
package web

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"

	"github.com/micro/go-api/handler"
	"github.com/micro/go-micro/selector"
)

type web struct {
	options handler.Options

	rp *httputil.ReverseProxy
	dr func(r *http.Request)
}

var (
	re = regexp.MustCompile("^[a-zA-Z0-9]+$")

	BasePathHeader = "X-Micro-Web-Base-Path"
)

func isWebSocket(r *http.Request) bool {
	contains := func(key, val string) bool {
		vv := strings.Split(r.Header.Get(key), ",")
		for _, v := range vv {
			if val == strings.ToLower(strings.TrimSpace(v)) {
				return true
			}
		}
		return false
	}

	if contains("Connection", "upgrade") && contains("Upgrade", "websocket") {
		return true
	}

	return false
}

func director(ns string, sel selector.Selector) func(r *http.Request) {
	return func(r *http.Request) {
		kill := func() {
			r.URL.Host = ""
			r.URL.Path = ""
			r.URL.Scheme = ""
			r.Host = ""
			r.RequestURI = ""
		}

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 2 {
			kill()
			return
		}
		if !re.MatchString(parts[1]) {
			kill()
			return
		}
		next, err := sel.Select(ns + "." + parts[1])
		if err != nil {
			kill()
			return
		}

		s, err := next()
		if err != nil {
			kill()
			return
		}

		r.Header.Set(BasePathHeader, "/"+parts[1])
		r.URL.Host = fmt.Sprintf("%s:%d", s.Address, s.Port)
		r.URL.Path = "/" + strings.Join(parts[2:], "/")
		r.URL.Scheme = "http"
		r.Host = r.URL.Host
	}
}

func (w *web) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	if !isWebSocket(r) {
		// the usual path
		w.rp.ServeHTTP(wr, r)
		return
	}

	// the websocket path
	req := new(http.Request)
	*req = *r
	w.dr(req)
	host := req.URL.Host

	if len(host) == 0 {
		http.Error(wr, "invalid host", 500)
		return
	}

	// connect to the backend host
	conn, err := net.Dial("tcp", host)
	if err != nil {
		http.Error(wr, err.Error(), 500)
		return
	}

	// hijack the connection
	hj, ok := wr.(http.Hijacker)
	if !ok {
		http.Error(wr, "failed to connect", 500)
		return
	}

	nc, _, err := hj.Hijack()
	if err != nil {
		return
	}

	defer nc.Close()
	defer conn.Close()

	if err = req.Write(conn); err != nil {
		return
	}

	errCh := make(chan error, 2)

	cp := func(dst io.Writer, src io.Reader) {
		_, err := io.Copy(dst, src)
		errCh <- err
	}

	go cp(conn, nc)
	go cp(nc, conn)

	<-errCh
}

func (w *web) String() string {
	return "web"
}

func NewHandler(opts ...handler.Option) handler.Handler {
	var options handler.Options

	for _, o := range opts {
		o(&options)
	}

	dr := director(
		options.Namespace,
		options.Service.Client().Options().Selector,
	)

	return &web{
		options: options,
		rp:      &httputil.ReverseProxy{Director: dr},
		dr:      dr,
	}
}

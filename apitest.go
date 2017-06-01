package apitest

import (
	"crypto/tls"
	"net/http/httptest"

	"net/http"
)

type Apitest struct {
	Handler http.Handler      // handler to test
	Server  *httptest.Server  // testing server
	clients chan *http.Client // http clients
}

func New(h http.Handler) *Apitest {

	return NewWithPool(h, 2)
}

func NewWithPool(h http.Handler, n int) *Apitest {

	a := &Apitest{
		Handler: h,
		Server:  httptest.NewServer(h),
		clients: make(chan *http.Client, n),
	}

	for i := 0; i < cap(a.clients); i++ {
		a.clients <- &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				DisableKeepAlives: false,
			},
		}
	}

	return a
}

func (a *Apitest) Destroy() {
	if nil != a.Server {
		a.Server.Close()
		a.Server = nil
	}
}

func (a *Apitest) Request(method, path string) *Request {

	return NewRequest(
		method,
		a.Server.URL+path,
		a,
	)
}

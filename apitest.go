package apitest

import (
	"net/http/httptest"

	"github.com/fulldump/golax"
)

type Apitest struct {
	Api    *golax.Api
	Server *httptest.Server
}

func New(api *golax.Api) *Apitest {
	return &Apitest{
		Api:    api,
		Server: httptest.NewServer(api),
	}
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
		a.Server.URL+a.Api.Prefix+path,
	)
}

package apitest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Request_Host(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Host != "my-host" {
			t.Errorf("Received header should be 'my-host' instead of '%s'", r.Host)
		}

	}))
	defer server.Close()

	a := NewWithBase(server.URL)
	a.Request(http.MethodGet, "/path").WithHost("my-host").Do()

}

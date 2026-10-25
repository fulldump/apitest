package apitest

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
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

func Test_Request_Query(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")
		if name != "John" {
			t.Errorf("Received query 'name' should be 'John' instead of '%s'", name)
		}

	}))
	defer server.Close()

	a := NewWithBase(server.URL)
	a.Request(http.MethodGet, "/path").
		WithQuery("name", "John").
		Do()

}

func Test_Request_Query_Multiple(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		have := r.URL.Query()["color"]
		want := []string{"red", "blue"}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("Received query 'color' should be '%s' instead of '%s'", want, have)
		}

	}))
	defer server.Close()

	a := NewWithBase(server.URL)
	a.Request(http.MethodGet, "/path").
		WithQuery("color", "red").
		WithQuery("color", "blue").
		Do()

}

func Test_Request_Query_Multiple_Combined(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		have := r.URL.Query()
		want := url.Values{
			"color": {"white", "red", "blue"},
			"name":  {"John"},
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("Received query 'color' should be '%s' instead of '%s'", want, have)
		}

	}))
	defer server.Close()

	a := NewWithBase(server.URL)
	a.Request(http.MethodGet, "/path?color=white&name=John").
		WithQuery("color", "red").
		WithQuery("color", "blue").
		Do()

}

// A server may answer before reading the request body. The recorded body must
// still be complete, and recording it must not race with the transport.
func Test_Request_BodyRecordedWhenServerAnswersEarly(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	// Large enough that the transport is still writing it when the answer
	// arrives.
	body := strings.Repeat("x", 8<<20)
	a := NewWithBase(server.URL)
	r := a.Request(http.MethodPost, "/path").WithBodyString(body).Do()
	if r.StatusCode != http.StatusBadRequest {
		t.Fatalf("status %d, want %d", r.StatusCode, http.StatusBadRequest)
	}
	if have := r.BodyRequestString(); have != body {
		t.Fatalf("recorded %d request body bytes, want %d", len(have), len(body))
	}
}

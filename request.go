package apitest

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	http.Request
}

func NewRequest(method, urlStr string) *Request {

	http_request, err := http.NewRequest(method, urlStr, strings.NewReader(""))
	if nil != err {
		panic(err)
	}

	return &Request{*http_request}

}

func (this *Request) WithCredentials(api_key, api_secret string) *Request {

	this.Header.Set("Api-Key", api_key)
	this.Header.Set("Api-Secret", api_secret)

	return this
}

func (this *Request) WithCookie(key, value string) *Request {

	c := &http.Cookie{
		Name:  key,
		Value: value,
	}

	this.AddCookie(c)
	return this
}

func (this *Request) set_body(body io.Reader) {

	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}
	this.Body = rc

	if body != nil {
		switch v := body.(type) {
		case *bytes.Buffer:
			this.ContentLength = int64(v.Len())
		case *bytes.Reader:
			this.ContentLength = int64(v.Len())
		case *strings.Reader:
			this.ContentLength = int64(v.Len())
		}
	}
}

func (this *Request) WithHeader(key, value string) *Request {

	this.Header.Set(key, value)

	return this
}

func (this *Request) WithBodyString(body string) *Request {
	b := strings.NewReader(body)
	this.set_body(b)

	return this
}

func (this *Request) WithBodyJson(o interface{}) *Request {
	bytes, err := json.Marshal(o)
	if nil != err {
		panic(err)
	}

	this.WithBodyString(string(bytes))

	return this
}

func (this *Request) Do() *Response {

	res, err := http.DefaultClient.Do(&this.Request)
	if err != nil {
		panic(err)
	}

	return &Response{*res}
}

package apitest

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	http.Response
	apitest *Apitest
	client  *http.Client
}

func (r *Response) BodyClose() {
	io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
}

func (r *Response) BodyBytes() []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	r.BodyClose()
	return body
}

func (r *Response) BodyString() string {
	return string(r.BodyBytes())
}

func (r *Response) BodyJson() interface{} {
	d := json.NewDecoder(r.Body)
	d.UseNumber()
	var body interface{}
	if err := d.Decode(&body); err != nil {
		panic(err)
	}
	r.BodyClose()
	return body
}

func (r *Response) BodyJsonMap() *map[string]interface{} {
	d := json.NewDecoder(r.Body)
	d.UseNumber()
	body := map[string]interface{}{}
	if err := d.Decode(&body); err != nil {
		panic(err)
	}
	r.BodyClose()
	return &body
}

package apitest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	http.Response
}

func (this *Response) BodyBytes() []byte {
	body, err := ioutil.ReadAll(this.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func (this *Response) BodyString() string {
	return string(this.BodyBytes())
}

func (this *Response) BodyJson() interface{} {
	d := json.NewDecoder(this.Body)
	d.UseNumber()
	var body interface{}
	if err := d.Decode(&body); err != nil {
		panic(err)
	}
	return body
}

func (this *Response) BodyJsonMap() *map[string]interface{} {
	d := json.NewDecoder(this.Body)
	d.UseNumber()
	body := map[string]interface{}{}
	if err := d.Decode(&body); err != nil {
		panic(err)
	}
	return &body
}

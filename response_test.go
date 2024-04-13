package apitest

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestResponse_BodyJsonMap(t *testing.T) {

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Hello World",
		})
	})

	a := NewWithHandler(h)

	req := a.Request("GET", "/").Do()
	message := req.BodyJsonMap()["message"]

	if message != "Hello World" {
		t.Errorf("Message should be 'Hello World'")
	}
}

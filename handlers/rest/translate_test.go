package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luzeduardo/shipping-go/handlers/rest"
)

func TestTranslateAPI(t *testing.T) {
	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/hello", nil)

	handler := http.HandlerFunc(rest.TranslateHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf(`expected status %d but received %d`, http.StatusOK, rr.Code)
	}

	var resp rest.Resp
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp.Language != "english" {
		t.Errorf(`expected language "english" but received %s`, resp.Language)
	}

	if resp.Translation != "hello" {
		t.Errorf(`expected Translation "hello" but received %s`, resp.Translation)
	}
}

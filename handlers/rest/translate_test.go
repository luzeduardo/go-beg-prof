package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luzeduardo/shipping-go/handlers/rest"
	"github.com/luzeduardo/shipping-go/translation"
)

func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint:            "/hello?language=notcoveredlanguage",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
		{
			Endpoint:            "/hello?language=dutch",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}
	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		underTest := rest.NewTranslatorHandler(translation.NewStaticService())
		handler := http.HandlerFunc(underTest.TranslateHandler)

		handler.ServeHTTP(rr, req)

		if rr.Code != test.StatusCode {
			t.Errorf(`expected status %d but received %d`, test.StatusCode, rr.Code)
		}

		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`expected language "%s" but received %s`, test.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`expected Translation "%s" but received %s`, test.ExpectedTranslation, resp.Translation)
		}
	}
}

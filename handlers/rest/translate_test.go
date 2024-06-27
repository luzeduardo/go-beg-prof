package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luzeduardo/shipping-go/handlers/rest"
)

type stubbedService struct{}

func (s *stubbedService) Translate(word string, language string) string {
	if word == "foo" {
		return "baz"
	}
	return ""
}
func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
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
		{
			Endpoint:            "/foo?language=english",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "baz",
		},
	}
	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		underTest := rest.NewTranslatorHandler(&stubbedService{})
		handler := http.HandlerFunc(underTest.TranslateHandler)

		handler.ServeHTTP(rr, req)

		if rr.Code != test.StatusCode {
			t.Errorf(`expected status %d but received %d on %s`, test.StatusCode, rr.Code, test.Endpoint)
		}

		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`expected language "%s" but received %s on %s`, test.ExpectedLanguage, resp.Language, test.Endpoint)
		}

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`expected Translation "%s" but received %s on %s`, test.ExpectedTranslation, resp.Translation, test.Endpoint)
		}
	}
}

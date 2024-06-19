package rest

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

const defaultLanguage = "english"

type Translator interface {
	Translate(word string, language string) string
}

type TranslateHandler struct {
	service Translator
}

func NewTranslatorHandler(service Translator) *TranslateHandler {
	return &TranslateHandler{
		service: service,
	}
}

func (t *TranslateHandler) TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = defaultLanguage
	}

	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translation := t.service.Translate(word, language)

	if translation == "" {
		language = ""
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := Resp{
		Language:    language,
		Translation: translation,
	}

	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}

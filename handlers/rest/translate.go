package rest

import (
	"encoding/json"
	"net/http"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resp := Resp{
		Language:    "English",
		Translation: "Hello",
	}

	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}

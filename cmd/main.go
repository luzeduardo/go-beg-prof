package main

import (
	"log"
	"net/http"

	"github.com/luzeduardo/shipping-go/config"
	"github.com/luzeduardo/shipping-go/handlers/rest"
	"github.com/luzeduardo/shipping-go/translation"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func main() {
	cfg := config.LoadConfiguration()
	addr := cfg.Port

	mux := http.NewServeMux()

	translateService := translation.NewStaticService()
	translateHandler := rest.NewTranslatorHandler(translateService)
	mux.HandleFunc("/translate/hello", translateHandler.TranslateHandler)

	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

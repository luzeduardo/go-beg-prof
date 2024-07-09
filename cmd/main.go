package main

import (
	"log"
	"net/http"

	"github.com/luzeduardo/shipping-go/config"
	"github.com/luzeduardo/shipping-go/handlers"
	"github.com/luzeduardo/shipping-go/handlers/rest"
	"github.com/luzeduardo/shipping-go/translation"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func main() {
	cfg := config.LoadConfiguration()

	var translationService rest.Translator
	translationService = translation.NewStaticService()

	if cfg.LegacyEndpoint != "" {
		log.Printf("creating external translation client: %s", cfg.LegacyEndpoint)
		client := translation.NewHelloClient(cfg.LegacyEndpoint)
		translationService = translation.NewRemoteService(client)
	}

	addr := cfg.Port

	mux := http.NewServeMux()

	// translateService := translation.NewStaticService()
	translateHandler := rest.NewTranslatorHandler(translationService)
	mux.HandleFunc("/translate/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/info", handlers.Info)

	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

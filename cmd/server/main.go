package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jkeddari/aitranslate/handlers"
	"github.com/jkeddari/aitranslate/translator"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("no api key found !")
	}

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":3000"
	}

	t, err := translator.NewTranslator(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	r := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	r.HandleFunc("GET /", handlers.NewHomeHandler().ServeHTTP)
	r.HandleFunc("POST /api/v1/translate", handlers.NewTranslateHander(t).Translate)
	r.HandleFunc("POST /translate", handlers.NewTranslateHander(t).ServeHTTP)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Println("Starting server on", addr)
	log.Fatal(srv.ListenAndServe())
}

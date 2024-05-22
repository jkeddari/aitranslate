package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jkeddari/aitranslate/handlers"
	"github.com/jkeddari/aitranslate/translator"
	"github.com/joho/godotenv"
)

type apiServer struct {
	translator *translator.Translator
}

func NewServer(apiKey, port string) (*apiServer, error) {
	t, err := translator.NewTranslator(apiKey)
	if err != nil {
		return nil, err
	}

	return &apiServer{translator: t}, nil
}

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("no api key found !")
	}

	t, err := translator.NewTranslator(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	r := http.NewServeMux()
	r.HandleFunc("POST /translate", handlers.NewTranslateHander(t).Translate)

	srv := &http.Server{
		Addr:    "localhost:3000",
		Handler: r,
	}

	srv.ListenAndServe()
}

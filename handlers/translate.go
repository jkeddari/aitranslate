package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jkeddari/aitranslate/model"
	"github.com/jkeddari/aitranslate/translator"
)

type TranslateHandler struct {
	translator *translator.Translator
}

func NewTranslateHander(t *translator.Translator) *TranslateHandler {
	return &TranslateHandler{t}
}

func (h *TranslateHandler) Translate(w http.ResponseWriter, r *http.Request) {
	var req model.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	text, err := h.translator.Translate(req.Source, req.Language)
	if err != nil {
		resp := model.Response{
			Success: false,
			Error:   err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := model.Response{
		Success: true,
		Text:    text,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *TranslateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sourceText := r.Form.Get("sourceText")
	language := r.Form.Get("targetLang")

	output, err := h.translator.Translate(sourceText, language)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

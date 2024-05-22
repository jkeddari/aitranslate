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
	json.NewDecoder(r.Body).Decode(&req)

	text, err := h.translator.Translate(req.Source, req.Language)
	if err != nil {
		resp := model.Response{
			Success: false,
			Error:   err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := model.Response{
		Success: true,
		Text:    text,
	}

	json.NewEncoder(w).Encode(resp)
}

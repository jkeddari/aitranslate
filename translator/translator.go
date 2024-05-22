package translator

import (
	"context"
	"errors"
	"slices"

	openai "github.com/sashabaranov/go-openai"
)

var LanguagesList = []string{
	"english",
	"french",
	"spanish",
	"german",
	"italian",
	"portuguese",
	"dutch",
	"russian",
	"japanese",
	"korean",
	"arabic",
	"hindi",
	"bengali",
	"urdu",
	"turkish",
	"polish",
	"swedish",
	"norwegian",
	"danish",
	"finnish",
	"greek",
	"czech",
	"hungarian",
	"romanian",
	"thai",
	"vietnamese",
	"indonesian",
	"malay",
	"hebrew",
	"swahili",
}

type Translator struct {
	client *openai.Client
}

func NewTranslator(apiKey string) (*Translator, error) {
	client := openai.NewClient(apiKey)
	return &Translator{client}, nil
}

func (t *Translator) Translate(source, language string) (string, error) {
	if !slices.Contains(LanguagesList, language) {
		return "", errors.New("language not supported !")
	}

	resp, err := t.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "translate the text in " + language + " : " + source,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

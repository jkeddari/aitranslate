package model

type Request struct {
	Language string `json:"language"`
	Source   string `json:"source"`
}

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Text    string `json:"text,omitempty"`
}

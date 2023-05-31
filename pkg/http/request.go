package http

type Request struct {
	Model string `json:"model"`
	Messages []Message `json:"messages"`
}

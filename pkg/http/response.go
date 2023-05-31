package http

type OpenAIResponse struct {
	ID string `json:"id"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Index int `json:"index"`
	Message Message `json:"message"`
}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

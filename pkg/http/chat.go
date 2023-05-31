package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const URL = "https://api.openai.com/v1/chat/completions"
const model = "gpt-3.5-turbo"
const role = "user"

func SendMessage(from, to, message string) (*OpenAIResponse, error) {

	var messages []Message
	messages = append(messages, Message{Role: role, Content: message})
	messageConstruct := Request{
		Model: model,
		Messages: messages,
	}
	messageBody, err := json.Marshal(messageConstruct)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(messageBody)
	req, err := http.NewRequest(http.MethodPost, URL, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", from))

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("response: %v", res)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected http response")
	}

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response OpenAIResponse
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

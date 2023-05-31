package main

import (
	"botchat/pkg/http"
	"log"
	"os"
	"time"
)

var AppName string
var AppVersion string

const (
	TopicInitiator = "TOPIC_INITIATOR"
	AgentOneSecret = "AGENT_ONE_SECRET"
	AgentTwoSecret = "AGENT_TWO_SECRET"

	DefaultTopic = "Hello! Tell me a good thing"

	ChatFor = 100
)

func init() {
	if AppVersion == "" {
		AppVersion = "dev-release"
	}
	if AppName == "" {
		AppName = "chat application"
	}
}

func getEnvOrDefault(envVar, defaultVal string) string {
	value, isSet := os.LookupEnv(envVar)
	if !isSet {
		value = defaultVal
	}
	return value
}

func main() {
	log.Printf("%v v: %v", AppName, AppVersion)

	topic := getEnvOrDefault(TopicInitiator, DefaultTopic)
	log.Printf("Let's start the conversation with %v", topic)

	agentOneSecret := getEnvOrDefault(AgentOneSecret, "")
	agentTwoSecret := getEnvOrDefault(AgentTwoSecret, "")

	from := agentOneSecret
	to := agentTwoSecret
	for idx := 0; idx < ChatFor; idx++ {
		response, err := http.SendMessage(from, to, topic)
		if err != nil {
			log.Printf("error: %v", err)
			os.Exit(1)
		}

		nextReceiver := from
		from = to
		to = nextReceiver

		topic = response.Choices[0].Message.Content
		log.Printf("Agent %d: Text: %v", idx % 2 + 1, topic)
		log.Printf("Sleeping until the next message")
		time.Sleep(time.Second * 30)
	}
}

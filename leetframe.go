package leetframe

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"leetframe.com/client/pubsublite"
)

type Config struct {
	ProjectID    string
	Zone         string
	TopicID      string
	MessageCount int
}

func Publish() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err.Error())
	}

	mc, err := strconv.Atoi(os.Getenv("MESSAGE_COUNT"))
	if err != nil {
		log.Fatal(err.Error())
	}

	envConfig := Config{
		ProjectID:    os.Getenv("PROJECT_ID"),
		Zone:         os.Getenv("ZONE"),
		TopicID:      os.Getenv("TOPIC_ID"),
		MessageCount: mc,
	}

	fmt.Printf("%v", envConfig)

	pubsublite.Publish(envConfig.ProjectID, envConfig.Zone, envConfig.TopicID, envConfig.MessageCount)

}

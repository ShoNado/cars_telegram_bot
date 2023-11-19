package handleAPI

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	TelegramBotToken string
}

func GetApiToken() string { //read token from file
	file, _ := os.Open("config/config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	return configuration.TelegramBotToken
}

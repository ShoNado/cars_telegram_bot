package carsAvailable

import (
	api "cars_telegram_bot/handleAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

func carsAvailable() {

}

func ShowCarsList() {

}

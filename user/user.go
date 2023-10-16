package user

import (
	api "cars_telegram_bot/handleAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

func HandleMessage(message *tgbotapi.Message) {
	if message.IsCommand() {
		handleCommand(message)
		return
	}
	msg := tgbotapi.NewMessage(message.From.ID, "")

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

func handleCommand(command *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(command.From.ID, "")

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на команду")
		panic(err)
	}
}

func HandleButton(query *tgbotapi.CallbackQuery) {
	msg := tgbotapi.NewMessage(query.From.ID, "")

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на кнопку")
		panic(err)
	}
}

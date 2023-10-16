package user

import (
	api "cars_telegram_bot/handleAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

func HandleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.From.ID, "Ты юзер")
	if _, err := bot.Send(msg); err != nil {
		panic(err) // not correct way handle error, remake!
	}
}

func handleCommand(chatId int64, command string) error {
	return nil
}

func HandleButton(query *tgbotapi.CallbackQuery) {
	msg := tgbotapi.NewMessage(query.From.ID, "Ты админ поздравляю")
	if _, err := bot.Send(msg); err != nil {
		panic(err) // not correct way handle error, remake!
	}
}

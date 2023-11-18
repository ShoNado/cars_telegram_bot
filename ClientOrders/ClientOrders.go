package ClientOrders

import (
	api "cars_telegram_bot/handleAPI"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

func ClientFavorites(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = fmt.Sprintf("Ваше избранное:\n\n В разрабокте")

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

package ClientOrders

import (
	api "cars_telegram_bot/handleAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

func ClientOrders(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Ваши заказы"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

func ClientFavorites(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Ваше избранное"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

func NewOrder(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Создание нового заказа"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

func OrdersList(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Список заказов от клиентов"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

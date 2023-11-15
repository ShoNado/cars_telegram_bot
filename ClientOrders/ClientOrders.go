package ClientOrders

import (
	api "cars_telegram_bot/handleAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

func ClientFavorites(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Ваше избранное"

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

func WarnClient(id int) {

}

package user

import (
	"cars_telegram_bot/carsAvailable"
	"cars_telegram_bot/clientOrders"
	api "cars_telegram_bot/handleAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
	btn1   = "Машины в наличии" //Машины в наличии
	btn2   = "Мои заказы"       //Мои заказы
	btn3   = "Избранное"        //Избранное
	btn4   = "Заказать машину"  //Заказать машину
)

func HandleMessage(message *tgbotapi.Message) {
	if message.IsCommand() {
		handleCommand(message)
		return
	}
	msg := tgbotapi.NewMessage(message.From.ID, "something wrong")

	switch message.Text {
	case btn1:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		carsAvailable.ShowCarsList()

	case btn2:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		clientOrders.ClientOrders()

	case btn3:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		clientOrders.ClientFavorites()

	case btn4:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		clientOrders.NewOrder()

	}
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

func handleCommand(command *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(command.From.ID, "")
	switch command.Command() {
	case "start":
		msg.Text = "Я бот помошник для покупки автомобиля вашей мечты . \n" +
			"Если у вас появтся дополнительные вопросы свяжитесь с нашим менеджером @blyaD1ma"
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(btn1),
				tgbotapi.NewKeyboardButton(btn2),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(btn3),
				tgbotapi.NewKeyboardButton(btn4),
			),
		)

	case "help":
		msg.Text = "0 блять помощи ПОМОГИТЕ"
	case "mycars":
		msg.Text = "work in progress"
	case "order":
		msg.Text = "work in progress"
	default:
		msg.Text = "Неизвестная команда"
	}

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

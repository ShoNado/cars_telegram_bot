package admin

import (
	"cars_telegram_bot/carsAvailable"
	"cars_telegram_bot/clientOrders"
	api "cars_telegram_bot/handleAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _    = tgbotapi.NewBotAPI(api.GetApiToken())
	adminList = []int64{
		362859506, //лиза
		//231043417, //я
		//дима
	}
	btn1 = "Список машин"
	btn2 = "Заявки клиентов"
	btn3 = "Добавить машину"
	btn4 = "Редактировать машину"
)

func CheckForAdmin(ID int64) bool {
	ok := false
	for _, op := range adminList {
		if op == ID {
			ok = true
		}
	}
	return ok
}

func HandleAdminMessage(message *tgbotapi.Message) {
	if message.IsCommand() {
		handleAdminCommand(message)
		return
	}
	msg := tgbotapi.NewMessage(message.From.ID, "work in progress")
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

	default:
		msg.Text = "дефолтное сообщение"
	}
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение админа")
		panic(err) // not correct way handle error, remake!
	}
}

func handleAdminCommand(command *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(command.From.ID, "")
	switch command.Command() {
	case "start":
		msg.Text = "К. \n" +
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
	}
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на команду админа")
		panic(err)
	}
}

func HandleAdminQuery(query *tgbotapi.CallbackQuery) {
	msg := tgbotapi.NewMessage(query.From.ID, "")

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалость ответить на нажатие кнопки админа")
		panic(err) // not correct way handle error, remake!
	}
}

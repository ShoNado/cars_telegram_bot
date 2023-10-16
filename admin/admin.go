package admin

import (
	api "cars_telegram_bot/handleAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _    = tgbotapi.NewBotAPI(api.GetApiToken())
	adminList = []int64{
		362859506, //лиза
		231043417, //я
	}
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
	msg := tgbotapi.NewMessage(message.From.ID, "")

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение админа")
		panic(err) // not correct way handle error, remake!
	}
}

func handleAdminCommand(command *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(command.From.ID, "")

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

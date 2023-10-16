package admin

import (
	api "cars_telegram_bot/handleAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	msg := tgbotapi.NewMessage(message.From.ID, "Ты админ поздравляю")
	if _, err := bot.Send(msg); err != nil {
		panic(err) // not correct way handle error, remake!
	}
}
func HandleAdminQuery(query *tgbotapi.CallbackQuery) {
	msg := tgbotapi.NewMessage(query.From.ID, "Вы нажали кнопочку")
	if _, err := bot.Send(msg); err != nil {
		panic(err) // not correct way handle error, remake!
	}
}

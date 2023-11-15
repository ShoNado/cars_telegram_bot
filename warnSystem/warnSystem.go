package warnSystem

import (
	api "cars_telegram_bot/handleAPI"
	"cars_telegram_bot/usersDB"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	bot, _    = tgbotapi.NewBotAPI(api.GetApiToken())
	adminList = []int64{
		362859506, //лиза
		//231043417, //я
		//314539937, //дима
	}
)

func WarnAdmin(profile usersDB.UserProfile, id int) error {
	var err error
	for _, admin := range adminList {
		msg := tgbotapi.NewMessage(admin, fmt.Sprintf("Получен новый заказ от пользователя %v\n"+
			"Время получения заявки: %v\n"+"Контакты клиента %v %v)",
			profile.UserName, profile.OrderTime, profile.NameFromUser, profile.PhoneNumber))
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(fmt.Sprintf(fmt.Sprintf("Подтвердить принятие заказа %v", id))),
				tgbotapi.NewKeyboardButton(fmt.Sprintf(fmt.Sprintf("Посмотреть подробнее о заявке %v", id))),
			))
		_, err = bot.Send(msg)
	}
	return err
}

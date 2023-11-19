package warnSystem

import (
	api "cars_telegram_bot/handleAPI"
	"cars_telegram_bot/handleUsersDB"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

func WarnAdmin(profile handleUsersDB.UserProfile, id int) error {
	var err error
	for _, admin := range handleUsersDB.GetAdminList() {
		msg := tgbotapi.NewMessage(admin, fmt.Sprintf("Получен новый заказ от пользователя %v\n"+
			"Время получения заявки: %v.%v.%v %v:%v\n"+"Контакты клиента %v %v)",
			profile.UserName, profile.OrderTime.Day(), profile.OrderTime.Month(), profile.OrderTime.Year(),
			profile.OrderTime.Hour(), profile.OrderTime.Minute(), profile.NameFromUser, profile.PhoneNumber))
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(fmt.Sprintf(fmt.Sprintf("Подтвердить принятие заказа %v", id))),
				tgbotapi.NewKeyboardButton(fmt.Sprintf(fmt.Sprintf("Посмотреть подробнее о заявке %v", id))),
			))
		_, err = bot.Send(msg)
	}

	return err
}

func WarnClient(id int, text string) {
	tgID, err := handleUsersDB.GetTgID(id)
	msg := tgbotapi.NewMessage(int64(tgID), fmt.Sprintf("Что-то пошло не так"))
	if err != nil {
		fmt.Println("что-то пошло не так")
	}
	if text == "Заказ" {
		msg.Text = fmt.Sprintf("Ваш заказ подвержден менеджером")
	} else if text == "Работа" {
		msg.Text = fmt.Sprintf("Менеджер работает с вашим заказом")
	}

	bot.Send(msg)
}

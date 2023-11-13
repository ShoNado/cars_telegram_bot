package AddEditDeleteCarDB

import (
	api "cars_telegram_bot/handleAPI"
	"cars_telegram_bot/handleDatabase"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _            = tgbotapi.NewBotAPI(api.GetApiToken())
	AddCar      bool  = false //индикатор перехода в состояние добавления машины
	IDglobal    int64         //запоминаем айди пользователя который хочет добавить машину
	UpdateLocal string
	UpdateHere  int = 1 //счетчик по обновлениям машины
)

func GetUpdates(text string) {
	var car handleDatabase.Car
	msg := tgbotapi.NewMessage(IDglobal, "")
	switch UpdateHere {
	case 1:
		msg.Text = "Добавление новой машины\n Укажите бренд:"
		bot.Send(msg)
		car.Brand = UpdateLocal

	case 2:
		msg.Text = fmt.Sprintf("Вы указали бренд %v:\n Укажите модель:", UpdateLocal)
		bot.Send(msg)
		car.Model = UpdateLocal

	case 10:
		msg.Text = "машина успешно добавлена"
		bot.Send(msg)
		AddCar = false
	}

	if car.IsCompleted == true {
		_, err := handleDatabase.AddNewCar(car)
		if err != nil {
			log.Printf("не удалось добавить новую машину")
		}
	}
}

func NewCar(id int64, msg tgbotapi.MessageConfig) {
	AddCar = true //переходим в режим добавления машины
	IDglobal = id

}

func CorrectCar(msg tgbotapi.MessageConfig) {
	msg.Text = "Редактирование старой машины"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

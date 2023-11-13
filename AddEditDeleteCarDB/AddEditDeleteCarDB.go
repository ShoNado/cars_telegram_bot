package AddEditDeleteCarDB

import (
	api "cars_telegram_bot/handleAPI"
	"cars_telegram_bot/handleDatabase"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _          = tgbotapi.NewBotAPI(api.GetApiToken())
	AddCar     bool = false //индикатор перехода в состояние добавления машины
	UpdateHere int          //счетчик по обновлениям машины
)

func GetUpdates(text string, msg tgbotapi.MessageConfig) {
	var car handleDatabase.Car
	switch UpdateHere {
	case 1:
		car.Brand = text
		if car.Brand != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали бренд: %v\n Укажите модель:", car.Brand)
			bot.Send(msg)
		}

	case 2:
		car.Model = text
		if car.Model != "" {
			UpdateHere += 1
		}
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

func NewCar(msg tgbotapi.MessageConfig) {
	AddCar = true //переходим в режим добавления машины
	UpdateHere = 1
	msg.Text = fmt.Sprintf("Добавление новой машины\n Укажите бренд:")
	bot.Send(msg)
}

func CorrectCar(msg tgbotapi.MessageConfig) {
	msg.Text = "Редактирование старой машины"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

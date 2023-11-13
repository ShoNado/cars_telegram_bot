package AddEditDeleteCarDB

import (
	api "cars_telegram_bot/handleAPI"
	"cars_telegram_bot/handleDatabase"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

func NewCar(msg tgbotapi.MessageConfig) {
	var car handleDatabase.Car
	msg.Text = "Добавление новой машины\n Укажите бренд:"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}

	id, err := handleDatabase.AddNewCar(car)
	if err != nil {
		log.Printf("не удалось добавить новую машину")
	}
	fmt.Print(id)
}

func CorrectCar(msg tgbotapi.MessageConfig) {
	msg.Text = "Редактирование старой машины"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

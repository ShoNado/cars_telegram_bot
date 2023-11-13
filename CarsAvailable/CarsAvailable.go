package CarsAvailable

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

func ShowCarsList(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Доступные машины:\n "
	carlist, err := handleDatabase.ReadAll()
	if err != nil {
		log.Printf("getting car list: %v", err)
	}

	for _, car := range carlist {
		msg.Text += fmt.Sprintf("%v. %v %v %v \n Двигатель: %v %v \n Цвет: %v \n Цена: %v \n Чтобы посмотреть подробнее нажмите /car_"+"%v \n\n",
			car.Id, car.Brand, car.Brand, car.Year, car.Enginevolume, car.Enginetype, car.Color, car.Price, car.Id)
	}

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

func NewCar(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Добавление новой машины"
	var car handleDatabase.Car

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

func CorrectCar(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Редактирование старой машины"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

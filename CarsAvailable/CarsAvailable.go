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

func ShowCarsListAvailable(msg tgbotapi.MessageConfig) {
	msg.Text = "Доступные машины:\n "
	carlist, err := handleDatabase.ReadAll()
	if err != nil {
		log.Printf("getting car list: %v", err)
	}
	count := 1
	for _, car := range carlist {
		if car.StatusBool == true {
			msg.Text += fmt.Sprintf("%v. %v %v %v \n Двигатель: %v %v \n Цвет: %v \n Цена: %v \n Чтобы посмотреть подробнее нажмите /car_"+"%v \n\n",
				count, car.Brand, car.Brand, car.Year, car.Enginevolume, car.Enginetype, car.Color, car.Price, car.Id)
			count += 1
		}
	}

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

func ShowCarsListOnWay(msg tgbotapi.MessageConfig) {
	msg.Text = "Машины в пути:\n "
	carlist, err := handleDatabase.ReadAll()
	if err != nil {
		log.Printf("getting car list: %v", err)
	}
	count := 1
	for _, car := range carlist {
		if car.StatusBool == false {
			msg.Text += fmt.Sprintf("%v. %v %v %v \n Двигатель: %v %v \n Цвет: %v \n Цена: %v \n Чтобы посмотреть подробнее нажмите /car_"+"%v \n\n",
				count, car.Brand, car.Brand, car.Year, car.Enginevolume, car.Enginetype, car.Color, car.Price, car.Id)
			count += 1
		}
	}

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

func CorrectCar(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Редактирование старой машины"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

package CarsAvailable

import (
	api "cars_telegram_bot/handleAPI"
	"cars_telegram_bot/handleDatabase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

func ShowCarsList(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Список машин"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
	carsAvailable()
}

func NewCar(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Добавление новой машины"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

func carsAvailable() {
	handleDatabase.ConnectDB()
}

func CorrectCar(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Редактирование старой машины"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

type Car struct {
	Brand        string
	Model        string
	Country      string
	Year         int
	Status       string
	Enginetype   string
	Enginevolume float64
	Transmission string
	DriveType    string
	Color        string
	Mileage      float64
	//FavoriteNum  int
	Other string
}

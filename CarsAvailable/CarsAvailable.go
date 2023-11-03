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
	Brand        string  `json:"brand,omitempty" :"brand"`
	Model        string  `json:"model,omitempty" :"model"`
	Country      string  `json:"country,omitempty" :"country"`
	Year         int     `json:"year,omitempty" :"year"`
	Status       string  `json:"status,omitempty" :"status"`
	Enginetype   string  `json:"enginetype,omitempty" :"enginetype"`
	Enginevolume float64 `json:"enginevolume,omitempty" :"enginevolume"`
	Transmission string  `json:"transmission,omitempty" :"transmission"`
	DriveType    string  `json:"drive_type,omitempty" :"drive_type"`
	Color        string  `json:"color,omitempty" :"color"`
	Mileage      float64 `json:"mileage,omitempty" :"mileage"`
	//FavoriteNum  int     `json:"favoritenum,omitempty" :"favorite_num"`
	Other string `json:"other,omitempty" :"other"`
}

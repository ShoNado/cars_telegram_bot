package CarsAvailable

import (
	api "cars_telegram_bot/handleAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

type Car struct {
	Brand        string  `json:"brand,omitempty"`
	Model        string  `json:"model,omitempty"`
	Country      string  `json:"country,omitempty"`
	Year         int     `json:"year,omitempty"`
	Status       string  `json:"status,omitempty"`
	Eniginetype  string  `json:"eniginetype,omitempty"`
	Enginevolume float64 `json:"enginevolume,omitempty"`
	Transmission string  `json:"transmission,omitempty"`
	DriveType    string  `json:"drive_type,omitempty"`
	Color        string  `json:"color,omitempty"`
	Mileage      float64 `json:"mileage,omitempty"`
	FavoriteNum  int     `json:"favoritenum,omitempty"`
	Other        string  `json:"other,omitempty"`
}

func ShowCarsList(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Список машин"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}

}

func NewCar(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Добавление новой машины"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

func carsAvailable() {

}

func CorrectCar(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.Text = "Редактирование старой машины"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

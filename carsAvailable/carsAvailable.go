package carsAvailable

import (
	api "cars_telegram_bot/handleAPI"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func ShowCarsList() {

	fmt.Println("список машин")
}

func NewCar() {
	honda := Car{Brand: "Honda", Model: "xr"}
	fmt.Println("добавить машину", honda)
}

func carsAvailable() {
	fmt.Println("вызываю список машин")
}

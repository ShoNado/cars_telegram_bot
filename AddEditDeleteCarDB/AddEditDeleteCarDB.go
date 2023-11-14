package AddEditDeleteCarDB

import (
	api "cars_telegram_bot/handleAPI"
	"cars_telegram_bot/handleDatabase"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
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
			msg.Text = fmt.Sprintf("Вы указали модель: %v\n Укажите страну производитель:", car.Model)
			bot.Send(msg)
		}
	case 3:
		car.Country = text
		if car.Country != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали страну производитель: %v\n Укажите год производства/постановки на первый учет:", car.Country)
			bot.Send(msg)
		}
	case 4:
		car.Year, _ = strconv.Atoi(text)
		if car.Year != 0 {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали год производства/постановки на первый учет: %v\n Укажите статус(доступна/впути/под заказ и тд):", car.Year)
			bot.Send(msg)
		}
	case 5:
		car.Status = text
		if car.Status != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали статус: %v\n Укажите вид двигателя (дизельб/бензин/гибрид/электричка), не объем!:", car.Status)
			bot.Send(msg)
		}
	case 6:
		car.Enginetype = text
		if car.Enginetype != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали вид двигателя: %v\n Укажите объем двигателя:", car.Enginetype)
			bot.Send(msg)
		}
	case 7:
		car.Enginevolume = text
		if car.Enginevolume != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали объем двигателя: %v\n Укажите  мощность (в л.с.):", car.Enginevolume)
			bot.Send(msg)
		}
	case 8:
		car.Horsepower = text
		if car.Horsepower != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали мощность: %v\n Укажите  крутящий момент (в н.м.):", car.Horsepower)
			bot.Send(msg)
		}
	case 9:
		car.Torque = text
		if car.Torque != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали крутящий момент: %v\n Укажите вид трансмиссии (механика/автома/робот/вариатор):", car.Torque)
			bot.Send(msg)
		}
	case 10:
		car.Transmission = text
		if car.Transmission != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали вид трансмисии: %v\n Укажите вид привода (полный/задний/передний/parttime):", car.Transmission)
			bot.Send(msg)
		}
	case 11:
		car.DriveType = text
		if car.DriveType != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали вид привода: %v\n Укажите  цвет:", car.DriveType)
			bot.Send(msg)
		}

	case 12:
		car.Color = text
		if car.Color != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали объем цвет: %v\n Укажите  пробеш (в км):", car.Color)
			bot.Send(msg)
		}

	case 25:
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
	msg.Text = fmt.Sprintf("Добавление новой машины\n Если хотите прервать добавление машины нажмите /stop \n Укажите бренд:")
	bot.Send(msg)
}

func CorrectCar(msg tgbotapi.MessageConfig) {
	msg.Text = "Редактирование старой машины"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

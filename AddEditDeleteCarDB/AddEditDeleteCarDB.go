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
	bot, _     = tgbotapi.NewBotAPI(api.GetApiToken())
	AddCar     = false //индикатор перехода в состояние добавления машины
	UpdateHere int     //счетчик по обновлениям машины
	newcar     handleDatabase.Car
)

func GetUpdates(text string, msg tgbotapi.MessageConfig) {

	switch UpdateHere {
	case 1:
		newcar.Brand = text
		if newcar.Brand != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали бренд: %v\n Укажите модель:", newcar.Brand)
			bot.Send(msg)
		}
	case 2:
		newcar.Model = text
		if newcar.Model != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали модель: %v\n Укажите страну производитель:", newcar.Model)
			bot.Send(msg)
		}
	case 3:
		newcar.Country = text
		if newcar.Country != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали страну производитель: %v\n Укажите год производства/постановки на первый учет:", newcar.Country)
			bot.Send(msg)
		}
	case 4:
		var err error
		newcar.Year, err = strconv.Atoi(text)
		if err == nil {
			if newcar.Year != 0 {
				UpdateHere += 1
				msg.Text = fmt.Sprintf("Вы указали год производства/постановки на первый учет: %v\n Укажите статус(доступна/впути/под заказ и тд):", newcar.Year)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("В пути"),
						tgbotapi.NewKeyboardButton("В наличии"),
					))
				bot.Send(msg)
			}
		}
	case 5:
		newcar.Status = text
		if newcar.Status == "В наличии" {
			newcar.StatusBool = true
		} else if newcar.Status == "В пути" {
			newcar.StatusBool = false
		}
		if newcar.Status == "В наличии" || newcar.Status == "В пути" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали статус: %v\n Укажите вид двигателя (дизельб/бензин/гибрид/электричка), не объем!:", newcar.Status)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
		}
	case 6:
		newcar.Enginetype = text
		if newcar.Enginetype != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали вид двигателя: %v\n Укажите объем двигателя (для электрических машин следует указать 0):", newcar.Enginetype)
			bot.Send(msg)
		}
	case 7:
		newcar.Enginevolume = text
		if newcar.Enginevolume != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали объем двигателя: %v\n Укажите  мощность (в л.с.):", newcar.Enginevolume)
			bot.Send(msg)
		}
	case 8:
		newcar.Horsepower = text
		if newcar.Horsepower != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали мощность: %v\n Укажите  крутящий момент (в н.м.):", newcar.Horsepower)
			bot.Send(msg)
		}
	case 9:
		newcar.Torque = text
		if newcar.Torque != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали крутящий момент: %v\n Укажите вид трансмиссии (механика/автома/робот/вариатор):", newcar.Torque)
			bot.Send(msg)
		}
	case 10:
		newcar.Transmission = text
		if newcar.Transmission != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали вид трансмисии: %v\n Укажите вид привода (полный/задний/передний/parttime):", newcar.Transmission)
			bot.Send(msg)
		}
	case 11:
		newcar.DriveType = text
		if newcar.DriveType != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали вид привода: %v\n Укажите  цвет:", newcar.DriveType)
			bot.Send(msg)
		}
	case 12:
		newcar.Color = text
		if newcar.Color != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали объем цвет: %v\n Укажите  пробег (в км):", newcar.Color)
			bot.Send(msg)
		}
	case 13:
		newcar.Mileage = text
		if newcar.Mileage != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали пробег: %v\n Укажите цену(либо напишите \"по запросу\"):", newcar.Mileage)
			bot.Send(msg)
		}
	case 14:
		newcar.Price = text
		if newcar.Price != "" {
			UpdateHere += 1
			msg.Text = fmt.Sprintf("Вы указали цену: %v\n Укажите примечание (если примечаний нет напишите \"нет\"):", newcar.Price)
			bot.Send(msg)
		}
	case 15:
		newcar.Other = text
		if newcar.Other != "" {
			if newcar.Other == "нет" {
				newcar.Other = ""
				msg.Text = fmt.Sprintf("Примечаний нет \n\n")
			} else {
				msg.Text = fmt.Sprintf("Вы указали примечание: %v\n\n ", newcar.Other)
			}
			UpdateHere += 1
			msg.Text += fmt.Sprintf("Подтвердите верность введеных данных\n Бренд: %v\n Модель: %v\n Страна производитель: %v\n "+
				"Год производства: %v\n Статус доставки: %v\n Тип двигателя: %v\n Объем двигателя: %v\n Мощность: %v\n Крутящий момент: %v\n"+
				"Коробка передач: %v\n Привод: %v\n Цвет: %v\n Пробег: %v\n Цена: %v\n Примечания: %v\n\n Все верно?", newcar.Brand, newcar.Model, newcar.Country,
				newcar.Year, newcar.Status, newcar.Enginetype, newcar.Enginevolume, newcar.Horsepower, newcar.Torque, newcar.Transmission, newcar.DriveType, newcar.Color,
				newcar.Mileage, newcar.Price, newcar.Other)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Да"),
					tgbotapi.NewKeyboardButton("Нет"),
				))
			bot.Send(msg)
		}
	case 16:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		if text == "Да" {
			newcar.IsCompleted = true
			id, err := handleDatabase.AddNewCar(newcar)
			if err != nil {
				msg.Text = fmt.Sprintf("Произошла ошибка при добавлении машины %v", id)
			} else {
				msg.Text = fmt.Sprintf("Машина успешно добавлена:\n /car_%v\n /menu", id)
			}
			bot.Send(msg)
			AddCar = false //выходим из режима добавления машины
		} else if text == "Нет" {
			newcar.IsCompleted = false
			id, err := handleDatabase.AddNewCar(newcar)
			if err != nil {
				msg.Text = fmt.Sprintf("Произошла ошибка при добавлении машины, в которой что-то не так")
			} else {
				msg.Text = fmt.Sprintf("Машина сохранена, неверные параметры можно скорректировать далее:\n /car_%v\n /menu", id)
			}
			bot.Send(msg)
			AddCar = false //выходим из режима добавления машины
		}
	}
}

func NewCar(msg tgbotapi.MessageConfig) {
	AddCar = true //переходим в режим добавления машины
	UpdateHere = 1
	msg.Text = fmt.Sprintf("Добавление новой машины\n Если хотите прервать добавление машины нажмите /stop " +
		"(переход в меню также прервет добавление машины) \n Укажите бренд:")
	bot.Send(msg)
}

func CorrectCar(msg tgbotapi.MessageConfig) { //todo
	msg.Text = "Редактирование старой машины"

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на сообщение")
		panic(err)
	}
}

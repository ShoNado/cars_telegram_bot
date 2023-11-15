package Admin

import (
	add "cars_telegram_bot/AddEditDeleteCarDB"
	"cars_telegram_bot/CarsAvailable"
	"cars_telegram_bot/ClientOrders"
	api "cars_telegram_bot/handleAPI"
	"cars_telegram_bot/handleDatabase"
	"cars_telegram_bot/usersDB"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

var (
	bot, _    = tgbotapi.NewBotAPI(api.GetApiToken())
	adminList = []int64{
		362859506, //лиза
		//231043417, //я
		//314539937, //дима
	}
	btn1 = "Машины в наличии"
	btn2 = "Машины в пути"
	btn3 = "Добавить машину"
	btn4 = "Заявки клиентов"
)

func CheckForAdmin(ID int64) bool {
	ok := false
	for _, op := range adminList {
		if op == ID {
			ok = true
		}
	}
	return ok
}

func HandleAdminMessage(message *tgbotapi.Message) {
	if message.IsCommand() {
		handleAdminCommand(message)
		return
	}
	msg := tgbotapi.NewMessage(message.From.ID, "")
	switch {
	case message.Text == btn1:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		CarsAvailable.ShowCarsListAvailable(msg)

	case message.Text == btn2:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		CarsAvailable.ShowCarsListOnWay(msg)

	case message.Text == btn3:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		add.NewCar(msg)

	case message.Text == btn4:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		ClientOrders.OrdersList(message, msg)
	case strings.HasPrefix(message.Text, "Да, удалить машину № "):
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		id, err := strconv.Atoi(message.Text[38:])
		if err != nil {
			msg.Text = "Что-то не так с айди"
			break
		}
		err = handleDatabase.DeleteCar(id)
		if err != nil {
			msg.Text = "Что-то не так с удалением"
		} else {
			msg.Text = fmt.Sprintf("Вы удалили машину №%v\n/menu", id)
		}
		bot.Send(msg)
	case message.Text == "Отмена":
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		msg.Text = "Удаление отменено"
		bot.Send(msg)
	case strings.HasPrefix(message.Text, "Подтвердить принятие заказа "):
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		id, _ := strconv.Atoi(message.Text[53:])
		ClientOrders.WarnClient(id)
		msg.Text = "Заказ успешно подтвержден"
		bot.Send(msg)
	case strings.HasPrefix(message.Text, "Посмотреть подробнее о заявке номер "):
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		id, _ := strconv.Atoi(message.Text[56:])
		ClientOrders.WarnClient(id)
		profile := usersDB.ShowOrder(id)
		fmt.Println(profile)
		msg.Text = "все гуд"
		bot.Send(msg)

	default:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		if add.AddCar == true {
			add.GetUpdates(message.Text, msg)
		} else {
			msg.Text = "Используйте /menu"
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("не удалость ответить на сообщение админа")
			}
		}

	}

}

func handleAdminCommand(command *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(command.From.ID, "")
	switch {
	case command.Command() == "start" || command.Command() == "menu":
		msg.Text = fmt.Sprintf("Используйте встроенную клавиатуру телеграмма")

		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(btn1),
				tgbotapi.NewKeyboardButton(btn2),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(btn3),
				tgbotapi.NewKeyboardButton(btn4),
			),
		)

	case command.Command() == "stop":
		add.AddCar = false
		msg.Text = fmt.Sprintf("Прерывано добавление машины\n /menu")
	case command.Command() == "help":
		msg.Text = "Список команд: \n" +
			"/start \n" +
			"/help \n" +
			"/stop"
	case strings.HasPrefix(command.Command(), "car_"):
		id, err := strconv.Atoi(command.Command()[4:])
		if err != nil {
			msg.Text = "Что-то не так с командой"
			break
		}
		car, err := handleDatabase.ShowCar(id)
		if err != nil {
			msg.Text = "Не удалось получить информацию о машине"
			break
		}
		msg.Text = fmt.Sprintf("Выбранная вами машина: \n Бренд: %v\n Модель: %v\n Страна производитель: %v\n "+
			"Год производства: %v\n Статус доставки: %v\n Тип двигателя: %v\n Объем двигателя: %v\n Мощность: %v\n Крутящий момент: %v\n"+
			"Коробка передач: %v\n Привод: %v\n Цвет: %v\n Пробег: %v\n Цена: %v\n Примечания: %v\n\n Вернутся в меню: /menu\n\n", car.Brand, car.Model, car.Country,
			car.Year, car.Status, car.Enginetype, car.Enginevolume, car.Horsepower, car.Torque, car.Transmission, car.DriveType, car.Color,
			car.Mileage, car.Price, car.Other)
		msg.Text += fmt.Sprintf("Для удаления машины из базы данных нажмите /delete_%v\n\n", id)
		msg.Text += fmt.Sprintf("Для изменения машины  нажмите /correct_%v\n\n", id)

	case strings.HasPrefix(command.Command(), "delete_"):
		id, err := strconv.Atoi(command.Command()[7:])
		if err != nil {
			msg.Text = "Что-то не так с командой"
			break
		}
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(fmt.Sprintf("Да, удалить машину № %v", id)),
				tgbotapi.NewKeyboardButton("Отмена"),
			))
		msg.Text = fmt.Sprintf("Подтвержите что желаете удалить машину № %v", id)
	case strings.HasPrefix(command.Command(), "correct_"):
		msg.Text = "В разработке"

	default:
		msg.Text = "Нет такой команды"
	}
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на команду админа")
		panic(err)
	}
}

func HandleAdminQuery(query *tgbotapi.CallbackQuery) {
	msg := tgbotapi.NewMessage(query.From.ID, "")

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалость ответить на нажатие кнопки админа")
		panic(err) // not correct way handle error, remake!
	}
}

package Admin

import (
	add "cars_telegram_bot/AddEditDeleteCarDB"
	"cars_telegram_bot/CarsAvailable"
	api "cars_telegram_bot/handleAPI"
	"cars_telegram_bot/handleDatabase"
	"cars_telegram_bot/usersDB"
	"cars_telegram_bot/warnSystem"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())

	btn1 = "Машины в наличии"
	btn2 = "Машины в пути"
	btn3 = "Добавить машину"
	btn4 = "Заявки клиентов"
)

func CheckForAdmin(ID int64) bool {
	ok := false
	for _, op := range warnSystem.AdminList {
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
		profiles := usersDB.ShowAllOrders()
		count := 1
		for _, profile := range profiles {
			msg.Text += fmt.Sprintf("%v. %v (%v) %v\n Тут будет время получения заявки\nУзнать подробнее /order_%v \n\n", count, profile.NameFromUser, profile.UserName, profile.PhoneNumber, profile.Id)
			count += 1
		}
		msg.Text += fmt.Sprintf("/menu")
		bot.Send(msg)
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
		warnSystem.WarnClient(id, "Заказ")
		err := usersDB.AdminSeen(id)
		if err != nil {
			fmt.Println("ошибка в смене статуса Admin seen")
		}
		msg.Text = "Заказ успешно подтвержден"
		bot.Send(msg)
	case strings.HasPrefix(message.Text, "Посмотреть подробнее о заявке "):
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		id, _ := strconv.Atoi(message.Text[56:])
		warnSystem.WarnClient(id, "Заказ")
		err := usersDB.AdminSeen(id)
		if err != nil {
			fmt.Println("ошибка в смене статуса Admin seen")
		}
		profile, _, err := usersDB.ShowOrder(id)
		if err != nil {
			msg.Text = "произошла какая-то ошибка при получении информации о заказе"
		} else {
			msg.Text = fmt.Sprintf("Данные о заказе №%v\nКонтактные данные человека:\n%v(%v) %v\n"+
				"Пожелание клиента по цене:\n%v\nПожелание клиента по бренду/модели:\n%v\nПожелания клиента по двигателю:\n%v\n"+
				"Пожелаения клиенту по коробке передач:\n%v\nПожелания клиенту по цвету:\n%v\nДополнительные пожелания:\n%v\nВремя заказа:\nВ разрабокте\n\n",
				id, profile.NameFromUser, profile.UserName, profile.PhoneNumber, profile.Price, profile.BrandCountryModel, profile.Engine, profile.Transmission,
				profile.Color, profile.Other)
			if profile.IsAdminSaw {
				msg.Text += "Потверждено администратором"
			} else {
				msg.Text += "Не подтверждено администратором"
			}
			if profile.IsInWork {
				msg.Text += fmt.Sprintf("\nВзято в работу")
			} else {
				msg.Text += fmt.Sprintf("\nНе взято в работу")
			}
			msg.Text += fmt.Sprintf("\nНе взято в работу\n\n /menu")
		}
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
	case strings.HasPrefix(command.Command(), "order_"):
		id, _ := strconv.Atoi(command.Command()[6:])
		profile, _, _ := usersDB.ShowOrder(id)
		msg.Text = fmt.Sprintf("Данные о заказе №%v\nКонтактные данные человека:\n%v(%v) %v\n"+
			"Пожелание клиента по цене:\n%v\nПожелание клиента по бренду/модели:\n%v\nПожелания клиента по двигателю:\n%v\n"+
			"Пожелаения клиенту по коробке передач:\n%v\nПожелания клиенту по цвету:\n%v\nДополнительные пожелания:\n%v\nВремя заказа:\nВ разработке\n",
			id, profile.NameFromUser, profile.UserName, profile.PhoneNumber, profile.Price, profile.BrandCountryModel, profile.Engine, profile.Transmission,
			profile.Color, profile.Other)

		if profile.IsAdminSaw {
			msg.Text += "Потверждено администратором"
		} else {
			msg.Text += "Не подтверждено администратором"
		}
		if profile.IsInWork {
			msg.Text += fmt.Sprintf("\nВзято в работу")
		} else {
			msg.Text += fmt.Sprintf("\nНе взято в работу\nВзять в работу /work_%v", id)
		}

	case strings.HasPrefix(command.Command(), "work_"):
		id, _ := strconv.Atoi(command.Command()[5:])
		err := usersDB.AdminGotInWork(id)
		if err != nil {
			fmt.Println("ошибка в смене статуса Admin works")
		}
		warnSystem.WarnClient(id, "Работа")
		msg.Text = fmt.Sprintf("Заказ взят №%v в работу\n/menu ", id)

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

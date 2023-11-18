package user

import (
	"cars_telegram_bot/AddClientInDB"
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
	bot, _  = tgbotapi.NewBotAPI(api.GetApiToken())
	btn1    = "Машины в наличии" //Машины в наличии
	btn2    = "Машины в пути"    //Мои заказы
	btn3    = "Избранное"        //Избранное
	btn4    = "Заказать машину"  //Заказать машину
	manager = "@blyaD1ma"
)

func HandleMessage(message *tgbotapi.Message) {
	if message.IsCommand() {
		handleCommand(message)
		return
	}
	msg := tgbotapi.NewMessage(message.From.ID, "")

	switch {
	case message.Text == btn1:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		CarsAvailable.ShowCarsListAvailable(msg) //передаем туда msg чтобы удалить клавиатуру

	case message.Text == btn2:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		CarsAvailable.ShowCarsListOnWay(msg)

	case message.Text == btn3:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		ClientOrders.ClientFavorites(message, msg)

	case message.Text == btn4:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		AddClientInDB.NewOrder(message, msg)

	default:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		if AddClientInDB.AddOrder == true {
			AddClientInDB.OrderUpdate(message, msg)
		} else {
			msg.Text = "Используйте /menu"
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("не удалость ответить на сообщение админа")
			}
		}

	}

}

func handleCommand(command *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(command.From.ID, "")
	switch {
	case command.Command() == "start":
		AddClientInDB.AddOrder = false
		msg.Text = "Я бот помошник для покупки автомобиля вашей мечты . \n" +
			"Если у вас появтся дополнительные вопросы свяжитесь с нашим менеджером " + manager
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
	case command.Command() == "menu":
		AddClientInDB.AddOrder = false
		msg.Text = "Используйте встроенную клавиатуру телеграмма"
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
	case command.Command() == "help":
		msg.Text = "Список доступных команд: \n" +
			"/start - перезапускает бота и вызывает основное меню \n" +
			"/mycars - показывает ваших избранных авто \n" +
			"/myorder - показывает ваш заказ и его статус \n" +
			"/order - новая заявка на подбор и заказ автомобиля \n" +
			"\nЕсли у вас есть какие-то дополнительные вопросы можете связатся с нашим менеджером " + manager
	case command.Command() == "myorder":
		profile, _ := usersDB.GetClientOrder(int(command.From.ID))
		msg.Text = fmt.Sprintf("Данные о заказе №%v\nКонтактные данные человека:\n%v(%v) %v\n"+
			"Пожелание клиента по цене:\n%v\nПожелание клиента по бренду/модели:\n%v\nПожелания клиента по двигателю:\n%v\n"+
			"Пожелаения клиенту по коробке передач:\n%v\nПожелания клиенту по цвету:\n%v\nДополнительные пожелания:\n%v\nВремя заказа:\nВ разрабокте\n\n",
			profile.Id, profile.NameFromUser, profile.UserName, profile.PhoneNumber, profile.Price, profile.BrandCountryModel, profile.Engine, profile.Transmission,
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
		msg.Text += "Для заказа свяжитесь с менеждером " + manager

	default:
		msg.Text = "Неизвестная команда"
	}

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на команду")
		panic(err)
	}
}

func HandleButton(query *tgbotapi.CallbackQuery) {
	msg := tgbotapi.NewMessage(query.From.ID, "где ты нашел кнопку не понял")

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Не удалось ответить на кнопку")
		panic(err)
	}
}

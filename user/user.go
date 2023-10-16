package user

import (
	"cars_telegram_bot/CarsAvailable"
	"cars_telegram_bot/ClientOrders"
	api "cars_telegram_bot/handleAPI"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
	btn1   = "Машины в наличии" //Машины в наличии
	btn2   = "Мои заказы"       //Мои заказы
	btn3   = "Избранное"        //Избранное
	btn4   = "Заказать машину"  //Заказать машину
)

func HandleMessage(message *tgbotapi.Message) {
	if message.IsCommand() {
		handleCommand(message)
		return
	}
	msg := tgbotapi.NewMessage(message.From.ID, "")

	switch message.Text {
	case btn1:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		CarsAvailable.ShowCarsList(message, msg) //передаем туда msg чтобы удалить клавиатуру

	case btn2:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		ClientOrders.ClientOrders(message, msg)

	case btn3:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		ClientOrders.ClientFavorites(message, msg)

	case btn4:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		ClientOrders.NewOrder(message, msg)

	default:
		msg.Text = "дефолтное сообщение"
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Не удалось ответить на сообщение")
			panic(err)
		}
	}

}

func handleCommand(command *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(command.From.ID, "")
	switch command.Command() {
	case "start":
		msg.Text = "Я бот помошник для покупки автомобиля вашей мечты . \n" +
			"Если у вас появтся дополнительные вопросы свяжитесь с нашим менеджером @blyaD1ma"
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

	case "help":
		msg.Text = "Список доступных команд: \n" +
			"/start - перезапускает бота и вызывает основное меню \n" +
			"/mycars - показывает информацию о вашем автомобиле и его статусе \n" +
			"/order - новая заявка на подбор и заказ автомобиля \n" +
			"\nЕсли у вас есть какие-то дополнительные вопросы можете связатся с нашим менеджером @blyaD1ma"
	case "mycars":
		msg.Text = "work in progress"
	case "order":
		msg.Text = "work in progress"
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

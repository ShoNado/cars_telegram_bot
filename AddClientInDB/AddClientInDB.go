package AddClientInDB

import (
	api "cars_telegram_bot/handleAPI"
	"cars_telegram_bot/usersDB"
	"cars_telegram_bot/warnSystem"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

var (
	bot, _      = tgbotapi.NewBotAPI(api.GetApiToken())
	AddOrder    = false
	UpdateOrder int
	profile     usersDB.UserProfile
	text        string
	manager     = "@blyaD1ma"
)

func NewOrder(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	AddOrder = true
	UpdateOrder = 1
	profile.UserName = message.From.UserName
	profile.NameFromTg = message.From.LastName + " " + message.From.FirstName
	profile.TgID = message.From.ID
	msg.Text = fmt.Sprintf("Оставьте заявку о том какой автомобиль вы бы хоели заказать. "+
		"После подачи заявки наш менеджер подберет автомобиль наиболее подходящий под ваши параметры и свяжется с вами."+
		"Не забудьте оставить свои контакты, чтобы менеджер мог связаться с вами. \n\n"+
		"Для отмены создания заказа нажмите /stop, либо вернитесь в меню /menu. \n\n"+
		"Если у вас остались вопросы свяжиетсь с нашим менеджером ") + manager +
		fmt.Sprintf("\n\n Укажите номер телефона для связи")
	bot.Send(msg)
}

func OrderUpdate(message *tgbotapi.Message, msg tgbotapi.MessageConfig) {
	text = message.Text
	switch UpdateOrder {
	case 1:
		profile.PhoneNumber = text
		if profile.PhoneNumber != "" {
			UpdateOrder += 1
			msg.Text = fmt.Sprintf("Ваш номер телефона: %v\nВерно?", profile.PhoneNumber)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Да, продолжить"),
					tgbotapi.NewKeyboardButton("Нет, изменить"),
				))
			bot.Send(msg)
		}
	case 2:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		if text == "Да, продолжить" {
			msg.Text = "Как можно к вам обращаться?"
			UpdateOrder += 1
			bot.Send(msg)
		} else if text == "Нет, изменить" {
			msg.Text = "Укажите верный номер телефона"
			UpdateOrder = 1
			bot.Send(msg)
		}
	case 3:
		profile.NameFromUser = text
		if profile.NameFromUser != "" {
			UpdateOrder += 1
			msg.Text = fmt.Sprintf("%v, подробно расскажите в одном сообщении,"+
				" о том на какой бюджет вы рассчитываете при заказе автомобиля. ", profile.NameFromUser)
			bot.Send(msg)
		}
	case 4:
		profile.Price = text
		if profile.Price != "" {
			UpdateOrder += 1
			msg.Text = fmt.Sprintf("%v, мы постараемся подобрать автомобиль подходящий под ваш бюджет. "+
				"Расскажите о ваших предпочтениях страны производиеля, бреда и модели. "+
				"Если вы уже выбрали машину мечты, то укажите её тут. "+
				"Также сооветуем сразу сообщить нам насколько вам важна именно конкретная модель или"+
				" вы готовы рассмотреть другие варианты в пределах бюджета", profile.NameFromUser)
			bot.Send(msg)
		}
	case 5:
		profile.BrandCountryModel = text
		if profile.BrandCountryModel != "" {
			UpdateOrder += 1
			msg.Text = fmt.Sprintf("Учтём ваши пожелания. Есть ли у вас предпочтения по двигателю " +
				"(дизель, бензин, гибрид, ээлектричепский), его объему?")
			bot.Send(msg)
		}
	case 6:
		profile.Engine = text
		if profile.Engine != "" {
			UpdateOrder += 1
			msg.Text = fmt.Sprintf("Какую трансмиссию вы хотите (автоматичекска/механическая)?")
			bot.Send(msg)
		}
	case 7:
		profile.Transmission = text
		if profile.Transmission != "" {
			UpdateOrder += 1
			msg.Text = fmt.Sprintf("Есть ли у вас предпочтения по цвету автомобиля или салона?")
			bot.Send(msg)
		}
	case 8:
		profile.Color = text
		if profile.Color != "" {
			UpdateOrder += 1
			msg.Text = fmt.Sprintf("Расскажите о ваших пожеланиях по поводу комплектации " +
				"или каких либо других параметрах автомобиля, которые вы хотели бы уточнить")
			bot.Send(msg)
		}
	case 9:
		profile.Other = text
		if profile.Other != "" {
			UpdateOrder += 1
			msg.Text = fmt.Sprintf("Спасибо за информацию. "+
				"Как только менеджер приступит к подбору автомобиля вы получите уведомление.\n\n"+
				"Если у вас остались вопросы или пожелания свяжитесь с менеджером ") + manager +
				fmt.Sprintf("\n\n Вернутся в меню /menu")

			profile.IsCompleted = true
			profile.OrderTime = time.Now()
			id, err := usersDB.AddNewOrder(profile)
			if err != nil {
				msg.Text += fmt.Sprintf("\nЧто-то пошло не так при добавлении в базу данных.")
			} else {
				msg.Text += fmt.Sprintf("\n\nНомер вашей заявки: %v", id)
			}
			err = warnSystem.WarnAdmin(profile, id)
			if err != nil {
				msg.Text += fmt.Sprintf("\nЧто-то пошло не так при отправке заявке администратору.")
			}
			bot.Send(msg)
		}
	default:
		msg.Text = "Что-то пошло не так"
		bot.Send(msg)
	}
}

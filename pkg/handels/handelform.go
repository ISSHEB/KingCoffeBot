package handels

import (
	"KingCoffe/pkg/tgform"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleForm(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case "О нас":
			tgform.AboutAs(bot, update)
		case "Адрес":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,
				"Кафе : KINGS’ COFFEE \n"+
					"Адрес: Москва, п.Сосенское, ул.Большое Понизовье 10")
			bot.Send(msg)
		case "Связаться с нами":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "🌐  https//kingcoffee \n \n"+
				"✉️  info@kingcoffee.ru \n \n"+
				"📱 +7 915 456 16 01")
			bot.Send(msg)
		}
	} else if update.Message != nil {
		switch update.Message.Command() {
		case "start":
			tgform.Form(bot, update.Message)
		}
	}
}

package handels

import (
	f "KingCoffe/internal/tgform"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleForm(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case "О нас":
			f.AboutAs(bot, update)
		case "Меню":

		case "Связаться с нами":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "🌐  https//kingcoffee \n \n"+
				"✉️  info@kingcoffee.ru \n \n"+
				"📱 +7 915 456 16 01")
			bot.Send(msg)
		}
	} else if update.Message != nil {
		switch update.Message.Command() {
		case "start":
			f.MainForm(bot, update.Message)
		}
	}
}

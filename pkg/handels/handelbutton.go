package handels

import (
	"KingCoffe/pkg/tgform"
	b "KingCoffe/pkg/tginline"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleButton(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	switch update.Message.Text {
	case "Каталог":
		// Your code for case 1
	case "Меню":
		// Your code for case 2
	case "Профель":

	case "Заказы":

	case "Админ панель":
		tgform.AdminPanal(bot, update)
	default:
		b.Button(bot, update.Message)
	}
}

package handels

import (
	"KingCoffe/pkg/tgform"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandelProduct(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case "Управление товаром":
			tgform.ProductForm(bot, update)
		case "Управление пользователем":
			// Your code for case 2
		}
	}
}

package handels

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandelAskProduc(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		switch update.Message.Text {
		case "Добавить Каталог":

		case "Добавить товар":
			//tgform.AddProduct(bot, update.Message)
		case "Редоктировать Каталог":

		case "Редоктировать товар":
			// Your code for case 2
		}
	}
}

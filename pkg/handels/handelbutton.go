package handels

import (
	b "KingCoffe/pkg/tginline"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleButton(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	switch update.Message.Text {
	case "1":
		// Your code for case 1
	case "2":
		// Your code for case 2
	case "3":
		// Your code for case 3
	case "4":
		// Your code for case 4
	case "5":
		// Your code for case 5
	case "6":
		// Your code for case 6
	default:
		b.Button(bot, update.Message)
	}
}

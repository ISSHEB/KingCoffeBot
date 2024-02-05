package tginline

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func MainButton(bot *tgbotapi.BotAPI, m *tgbotapi.Message) {
	buttons := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("Button 1", "button1"),
		tgbotapi.NewInlineKeyboardButtonData("Button 2", "button2"),
		tgbotapi.NewInlineKeyboardButtonData("Button 3", "button3"),
	}

	rows := [][]tgbotapi.InlineKeyboardButton{
		buttons,
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(rows)

	message := tgbotapi.NewMessage(m.Chat.ID, "This is a message with inline buttons.")
	message.ReplyMarkup = keyboard

	bot.Send(message)

}

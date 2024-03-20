package tginline

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"log"
)

func Button(bot *tgbotapi.BotAPI, m *tgbotapi.Message) {
	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Каталог"),
		), tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Меню"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Профель"),
			tgbotapi.NewKeyboardButton("Заказы"),
		), tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Админ панель"),
		),
	)

	msg := tgbotapi.NewMessage(m.Chat.ID, "панель открыта 🎛")
	msg.ReplyMarkup = numericKeyboard

	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

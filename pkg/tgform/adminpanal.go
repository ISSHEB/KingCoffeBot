package tgform

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func AdminPanal(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	form := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Управление товаром", "Управление товаром"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Управление пользователем", "Управление пользователем"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите что хотите сделать")
	msg.ReplyMarkup = form

	bot.Send(msg)
}

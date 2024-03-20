package tgform

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func ProductForm(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	form := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить Каталог", "Добавить Каталог"),
			tgbotapi.NewInlineKeyboardButtonData("Добавить товар", "Добавить товар"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Редоктировать Каталог", "Редоктировать Каталог"),
			tgbotapi.NewInlineKeyboardButtonData("Редоктировать товар", "Редоктировать товар"),
		),
	)

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Продукт")
	msg.ReplyMarkup = form

	bot.Send(msg)
}

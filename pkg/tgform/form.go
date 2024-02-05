package tgform

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

func Form(bot *tgbotapi.BotAPI, m *tgbotapi.Message) {
	// Создайте макет формы
	form := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("О нас", "О нас"),
			tgbotapi.NewInlineKeyboardButtonData("Адрес", "Адрес"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Связаться с нами", "Связаться с нами"),
		),
	)

	// Отправьте форму пользователю
	file, _ := os.Open("./photo/photo_2024-02-04_18-45-52.jpg")
	reader := tgbotapi.FileReader{Name: "./photo/photo_2024-02-04_18-45-52.jpg", Reader: file}

	msg := tgbotapi.NewPhoto(m.Chat.ID, reader)
	msg.Caption = "Добро пожаловать!👋" +
		" \n Я бот-помощник. 🤖" +
		"\n Могу рассказать вам о нашем кофе и помочь сделать заказ. 🗯" +
		"\n Просто зайдите в меню и выберите себе кофе или закуску. ☕️"
	msg.ReplyMarkup = form
	bot.Send(msg)

}

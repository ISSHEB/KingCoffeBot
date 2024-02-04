package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Создайте бота
	bot, err := tgbotapi.NewBotAPI("6739867707:AAHLH2uTo13WuuWNDjjxPiblWrBFNrNGPxs")
	if err != nil {
		panic(err)
	}

	// Зарегистрируйте обработчик формы
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		handleForm(bot, update)

	}

}

func mainForm(bot *tgbotapi.BotAPI, m *tgbotapi.Message) {
	// Создайте макет формы
	form := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("О нас", "О нас"),
			tgbotapi.NewInlineKeyboardButtonData("Меню", "Меню"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Адрес", "Адрес"),
		),
	)

	// Отправьте форму пользователю
	msg := tgbotapi.NewMessage(m.Chat.ID, "Пожалуйста, выберите опцию:")
	msg.ReplyMarkup = form
	bot.Send(msg)

}

func handleForm(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case "О нас":
			text := "О нас :\n\n" +
				"Компания Kings Coffee была founded in 1983 году одним человеком, благодаря амбициям которого, маленькая кофейная мельница стала ядром того, что в наши дни известно как крупнейшая сеть обжарочных и мельничных заводов в государстве Кувейт.\n" +
				"Мы рады сообщить, что компания Kings Coffee официально вышла на рынок России в 2023 году и начала прямые поставки в Российскую Федерацию.\n" +
				"Наш кофе проходит специальные исследования, а также обработку и обжарку в Государстве Кувейт. Мы используем новейшее оборудование и технологии для измельчения и помола кофе.\n" +
				"Расфасовка и упаковка кофе происходит исключительно в Кувейте по всем стандартам и санитарным нормам государства Кувейт и Российской Федерации.\n" +
				"Попробуйте настоящий арабский кофе, ощутите незабываемый аромат и почувствуйте вкус, с которым\n" +
				"Вам не захочется расставаться"
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
			bot.Send(msg)
		case "Меню":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Это наша информация о нас.")
			bot.Send(msg)
		case "Адрес":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Это наша информация о нас.")
			bot.Send(msg)
		}
	} else if update.Message != nil {
		switch update.Message.Command() {
		case "start":
			mainForm(bot, update.Message)
		}
	}
}

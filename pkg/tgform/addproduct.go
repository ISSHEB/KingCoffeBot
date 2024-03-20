package tgform

//import (
//	"fmt"
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//	"gopkg.in/yaml.v2"
//)
//
//type Product struct {
//	Name        string `yaml:"name"`
//	Price       string `yaml:"price"`
//	Description string `yaml:"description"`
//}
//type Items []Product
//
//func (i *Items) Add(Name string, Price string, Description string) {
//	// Создание нового экземпляра структуры item
//	item := Product{
//		Name:        Name,
//		Price:       Price,
//		Description: Description,
//	}
//
//	// Добавление нового экземпляра структуры item в массив t
//	*i = append(*i, item)
//}
//
//func (i *Items) AddProduct(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
//
//	// Создаем новый экземпляр структуры Product
//	product := &Product{}
//
//	// Функция для обработки ответов пользователя
//	handler := func(text string) {
//		switch {
//		case text == "Отмена":
//			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добавление товара отменено.")
//			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
//			bot.Send(msg)
//		default:
//			switch product.Name {
//			case "":
//				product.Name = text
//				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Напишите цену товара:")
//				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{
//					tgbotapi.NewKeyboardButton("Отмена"),
//				})
//				bot.Send(msg)
//			case product.Price:
//				product.Price = text
//				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Напишите описание товара:")
//				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{
//					tgbotapi.NewKeyboardButton("Отмена"),
//				})
//				bot.Send(msg)
//			case product.Description:
//				product.Description = text
//				data, _ := yaml.Marshal(product)
//				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Товар добавлен:\n\n%s", data))
//				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
//				bot.Send(msg)
//			}
//		}
//	}
//
//	// Отправляем сообщение пользователю
//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Напишите название товара:")
//	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{
//		tgbotapi.NewKeyboardButton("Отмена"),
//	})
//	bot.Send(msg)
//
//	// Ожидаем ответ пользователя
//	for {
//		update := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
//			Timeout: 60,
//		})
//
//		for update := range update {
//			if update.Message == nil { // ignore callback queries
//				continue
//			}
//
//			handler(update.Message.Text)
//		}
//	}
//}

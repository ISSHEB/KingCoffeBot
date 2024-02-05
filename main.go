package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	h "KingCoffe/pkg/handels"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	token := os.Getenv("TOKEN")
	// Создайте бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	fmt.Println("bot runnig!")

	// Зарегистрируйте обработчик формы
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		h.HandleForm(bot, update)
		h.HandleButton(bot, update)

	}

}

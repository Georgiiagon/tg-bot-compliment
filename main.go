package main

import (
	"log"
	"os"
	"tg-bot-compliment/api"
	"tg-bot-compliment/helpers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	helpers.LoadEnv()

	sendTelegramMessage()
}

func sendTelegramMessage() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.Text = api.GetRandomCompliment()

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

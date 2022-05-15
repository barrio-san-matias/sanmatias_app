package handler

import (
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joeshaw/envdecode"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("incoming request")

	var cfg struct {
		TelegramToken string `env:"TELEGRAM_TOKEN,required"`
	}
	if err := envdecode.StrictDecode(&cfg); err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	update, err := bot.HandleUpdate(r)
	if err != nil {
		log.Fatal(err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Hi there %s", update.Message.From.UserName))
	msg.ReplyToMessageID = update.Message.MessageID

	_, err = bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, "OK")
}

package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joeshaw/envdecode"
	redis "github.com/redis/go-redis/v9"
)

func TelegramHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("incoming request")

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Fatalf(">>> REDIS ERROR: %+v", err)
	}

	var cfg struct {
		TelegramToken string `env:"TELEGRAM_TOKEN,required"`
		KVUrl         string `env:"KV_URL,required"`
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

	fmt.Printf(">>>>>>> UPDATE: %+v\n", update)
	if update.Message.Contact != nil {
		fmt.Printf(">>>>>>> CONTACT: %+v\n", update.Message.Contact.PhoneNumber)
	}

	// si empieza con "lote <numero>" devuelve la ubicacion,
	// sino tira la consulta a dialogflow
	if update.Message.Text == "/lote 636" {
		msg := tgbotapi.NewVenue(
			update.Message.Chat.ID,
			"lote 636",
			"(tocar el mapa para abrir y navegar)",
			-34.3507263,
			-58.7637032,
		)

		msg.ReplyToMessageID = update.Message.MessageID

		_, err = bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		msg := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf(">>> sender: %+v\n, update:%+v", update.Message.From, update),
		)

		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.KeyboardButton{
					Text:           "compartir numero?",
					RequestContact: true,
				}))

		_, err = bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}

	}

	fmt.Fprint(w, "OK")
}

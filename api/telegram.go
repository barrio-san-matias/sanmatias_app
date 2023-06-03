package handler

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joeshaw/envdecode"
	redis "github.com/redis/go-redis/v9"
)

type User struct {
	Id       string `redis:"id"`
	Phone    string `redis:"phone"`
	Name     string `redis:"name"`
	LastName string `redis:"last_name"`
	Username string `redis:"username"`

	CertifiedTime time.Time `redis:"certified_time"`

	CreateTime time.Time `redis:"create_time"`
	UpdateTime time.Time `redis:"update_time"`
}

func TelegramHandler(w http.ResponseWriter, r *http.Request) {
	var cfg struct {
		TelegramToken string `env:"TELEGRAM_TOKEN,required"`
		KVUrl         string `env:"KV_URL,required"`
	}

	if err := envdecode.StrictDecode(&cfg); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	kv, err := getRedisClient(cfg.KVUrl)
	if err != nil {
		log.Fatalf("cannot get kv client: %v", err)
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

	fmt.Fprint(w, "OK")

	u, err := upsertUser(ctx, kv, update)
	if err != nil {
		panic(err)
	}

	var msg tgbotapi.MessageConfig
	if u == nil || u.Phone == "" {
		msg = getContactRequestMessage(update)
		_, err = bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if u != nil && u.CertifiedTime.IsZero() {
		msg = getAccessDeniedMessage(update)
		_, err = bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// handle bot request, user is authorized:
	msg = getResponseMessage(update, "TO-DO")
	_, err = bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func getContactRequestMessage(update *tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Bienvenido al Bot de San Matias. Solo vecinos certificados.")

	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.KeyboardButton{
				Text:           "compartir numero?",
				RequestContact: true,
			}))

	return msg
}

func getAccessDeniedMessage(update *tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"⛔️ Su numero de telefono no esta registrado como un vecino de San Matias.")

	return msg
}

func getResponseMessage(update *tgbotapi.Update, res string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		res)

	return msg
}

func getRedisClient(u string) (*redis.Client, error) {
	opt, err := redis.ParseURL(u)
	if err != nil {
		return nil, err
	}

	opt.TLSConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rdb := redis.NewClient(opt)

	return rdb, nil
}

func upsertUser(ctx context.Context, kv *redis.Client, update *tgbotapi.Update) (*User, error) {
	if update == nil {
		return nil, nil
	}

	spew.Dump(update)

	id := fmt.Sprint(update.Message.From.ID)

	var user *User = &User{}
	// Scan all fields into the model.
	if err := kv.HGetAll(ctx, id).Scan(&user); err != nil {
		return nil, err
	}

	if user == nil {
		user = &User{
			Id: id,
		}
	}

	if update.Message.Contact != nil {
		user.Phone = update.Message.Contact.PhoneNumber
	}

	user.UpdateTime = time.Now()

	// save user
	kv.HSet(ctx, id, user)

	return user, nil
}

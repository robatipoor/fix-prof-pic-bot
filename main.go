package main

import (
	"bytes"
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

var port string
var token string
var appURL string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	appURL = os.Getenv("APP_URL")
	if appURL == "" {
		log.Fatal("Application URL not set")
	}
	token = os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatalln("TELEGRAM_TOKEN not set !")
	}
}

func main() {

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: appURL},
	}
	setings := tb.Settings{
		Token:  token,
		Poller: webhook,
	}
	bot, err := tb.NewBot(setings)
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle(tb.OnPhoto, func(m *tb.Message) {
		b, err := fixSizeImage(getFile(m.Photo.FileID))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("ok fixed size image")
		ph := &tb.Photo{File: tb.FromReader(bytes.NewBuffer(b))}
		bot.Send(m.Sender, ph)
	})

	log.Println("Bot Start ...")
	bot.Start()
}

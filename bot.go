package main

import (
	"bytes"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func handelBot() *tb.Bot {

	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalln(err)
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
	
	return bot
}

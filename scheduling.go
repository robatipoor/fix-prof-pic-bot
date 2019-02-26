package main

import (
	"log"

	"github.com/jasonlvhit/gocron"
)

func scheduledTask() {

	s := gocron.NewScheduler()
	s.Every(30).Minutes().Do(func() {
		s, err := get("https://fixtelegramprofilepicturebot.herokuapp.com/")
		if err != nil {
			log.Println(err)
		}
		log.Println(s)
	})
	log.Println("Start Scheduled Task")
	<-s.Start()
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var token string
var address string

func init() {
	token = os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatalln("Port Not Set !")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("Set Port Number = ", port)
	}
	address = "0.0.0.0:" + port
}

func main() {

	bot := handelBot()
	go bot.Start()
	http.HandleFunc("/", handler)
	log.Println("Start Server : ", address)
	log.Fatal(http.ListenAndServe(address, nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Start Bot !")
}

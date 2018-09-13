package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Syfaro/telegram-bot-api"
)

const (
	ChatID      = 1 //Insert here your chat id
	Token       = "Your telegram bot token"
	OkResponse  = "Ok."
	BadResponse = "Bad."
	URL         = "Url to your remote server"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized %s", bot.Self.UserName)

	var lastResponse int
	for {
		var msg tgbotapi.MessageConfig

		//TODO: Use WS instead of HTTP
		resp, _ := http.Get(URL)

		log.Printf("Server status code is: %v \n %v \n", resp.StatusCode, resp.Body)

		if resp.StatusCode != lastResponse {
			if resp.StatusCode != 200 {
				msg = tgbotapi.NewMessage(ChatID, BadResponse)
			} else {
				msg = tgbotapi.NewMessage(ChatID, OkResponse)
			}
			bot.Send(msg)
		}
		lastResponse = resp.StatusCode
		time.Sleep(30 * time.Second)
	}
}

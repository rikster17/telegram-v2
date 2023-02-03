package main

import (
	"fmt"
	"log"
	"net/http"
	"string"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5161306872:AAFR2360W3NpwUnaa3IBrNiEjUfmT57BUH0")
	if err != nil {
			log.Panic(err)
	}
}

bot.Debug = true
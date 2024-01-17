package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type wallet map[string]float64

var db = map[int64]wallet{}

func main() {
	bot, err := tgbotapi.NewBotAPI("6670379625:AAH44qoKmV7AzRiXlq-A2kqeSKvPs2MnoM4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}

		command := strings.Split(update.Message.Text, " ")
		command[0] = strings.ToUpper(command[0])
		switch command[0] {
		case "ADD":
			ADD(command, bot, &update)
		case "SUB":
			SUB(command, bot, &update)
		case "DEL":
			DEL(command, bot, &update)
		case "SHOW":
			SHOW(bot, &update)
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Command has`t been found"))
		}

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, command[0])
		//msg.ReplyToMessageID = update.Message.MessageID
		//bot.Send(msg)
	}
}

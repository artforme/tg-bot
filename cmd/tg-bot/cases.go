package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func ADD(command []string, bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if len(command) != 3 {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "invalid command"))
	} else {
		res, err1 := strconv.ParseFloat(command[2], 64)
		if err1 != nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err1.Error()))
			return
		}

		if _, ok := db[update.Message.Chat.ID]; !ok {
			db[update.Message.Chat.ID] = wallet{}
		}
		db[update.Message.Chat.ID][command[1]] += res
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "successfully ADD"))
	}
}

func SUB(command []string, bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if len(command) != 3 {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "invalid command"))
	} else {
		res, err1 := strconv.ParseFloat(command[2], 64)
		if err1 != nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err1.Error()))
			return
		}

		if _, ok := db[update.Message.Chat.ID]; !ok {
			db[update.Message.Chat.ID] = wallet{}
		}
		if db[update.Message.Chat.ID][command[1]] < res {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "you have less values in the balance that you are going to get"))
			return
		}
		db[update.Message.Chat.ID][command[1]] -= res
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "successfully SUB"))
	}
}

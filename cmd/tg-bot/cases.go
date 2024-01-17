package main

import (
	"fmt"
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
		str := fmt.Sprintf("Successefully add; current balance is %f", db[update.Message.Chat.ID][command[1]])
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, str))
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
		str := fmt.Sprintf("Successefully sub; current balance is %f", db[update.Message.Chat.ID][command[1]])
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, str))
	}
}

func DEL(command []string, bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if len(command) != 2 {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "invalid command"))
	} else {
		delete(db[update.Message.Chat.ID], command[1])
		str := fmt.Sprintf("Balance %s was successufully deleted", command[1])
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, str))
	}
}

func SHOW(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if len(db[update.Message.Chat.ID]) == 0 {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Your wallet is empty"))
		return
	}
	msg := "Your wallet contains:\n"

	for key, value := range db[update.Message.Chat.ID] {
		msg += fmt.Sprintf("%s: %f\n", key, value)
	}
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
}

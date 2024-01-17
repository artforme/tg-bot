package main

import (
	"encoding/json"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
)

type bnResp struct {
	Price float64 `json:"price,string"`
	Code  int64   `json:"code"`
}

func getPrice(symbol string, bot *tgbotapi.BotAPI, update *tgbotapi.Update) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.binance.com/api/v3/avgPrice?symbol=%sRUB", symbol))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var jsonResp bnResp

	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return 0, err
	}
	if jsonResp.Code != 0 {
		err = errors.New("invalid symbol")
	}
	return jsonResp.Price, nil
}

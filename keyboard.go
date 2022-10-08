package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("点我"),
		// tgbotapi.NewKeyboardButton("2"),
		// tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

func handleKeyboard(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	switch update.Message.Text {
	case "open":
		msg.ReplyMarkup = numericKeyboard
	case "close":
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	default:
		return false
	}

	if _, err := bot.Send(msg); err != nil {
		log.Println(msg)
		return false
	}
	return true
}

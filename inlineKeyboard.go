package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func handleInlineKeyboard(update tgbotapi.Update) bool {
	if update.Message != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// If the message was open, add a copy of our numeric keyboard.
		switch update.Message.Text {
		case "open_inline":
			msg.ReplyMarkup = numericInlineKeyboard
		}

		// Send the message.
		if _, err := bot.Send(msg); err != nil {
			Log(err)
			return false
		}
	} else if update.CallbackQuery != nil {
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
		// callback.URL = "https://baidu.com"
		if _, err := bot.Request(callback); err != nil {
			Log(err)
			return false
		}

		remoteMsg := tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

		// And finally, send a message containing the data received.

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("4", "4"),
				tgbotapi.NewInlineKeyboardButtonData("5", "5"),
				tgbotapi.NewInlineKeyboardButtonData("6", "6"),
			),
		)

		go bot.Send(remoteMsg)
		go bot.Send(msg)
	}

	return true
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	Token string `json:"token"`
}

var bot *tgbotapi.BotAPI

func main() {
	config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Panic(err)
	}

	var c Config

	if err := json.Unmarshal(config, &c); err != nil {
		log.Panic(err)
	}

	bot, err = tgbotapi.NewBotAPI(c.Token)
	if err != nil {
		panic(err)
	}

	// bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 10

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		// Extract the command from the Message.
		if handleCommand(update) {
			Log(1232)
			continue
		}
		if handleKeyboard(update) {
			Log(3245)
			continue
		}
		if handleInlineKeyboard(update) {
			Log(345345)
			continue
		}
		Log(45645645)

		handleMessage(update)

	}

}

func handleMessage(update tgbotapi.Update) {
	if update.Message == nil {
		PrintJson(update)
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// We'll also say that this message is a reply to the previous message.
	// For any other specifications than Chat ID or Text, you'll need to
	// set fields on the `MessageConfig`.
	msg.ReplyToMessageID = update.Message.MessageID

	// Okay, we're sending our message off! We don't care about the message
	// we just sent, so we'll discard it.
	if _, err := bot.Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		panic(err)
	}
}

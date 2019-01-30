package commands

import (
	"log"
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

// ProcessCommands process commands
func ProcessCommands(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	chatID := message.Chat.ID
	command := extractCommand(message.Command())

	switch command {

	case "help":
		sendMsg(bot, chatID, "Here will be help text")

	case "start":
		sendMsg(bot, chatID, "Hi there! This bot helps you to monitor prices for a goods and notify you when price drops.")

	case "add":
		sendMsg(bot, chatID, "Ok, now send me the link of an item you want me to watch:")

	case "list":
		sendMsg(bot, chatID, "Here are list of all the saved items to watch")
	}
}

// properly extracts command from the input string, removing all unnecessary parts
// please refer to unit tests for details
func extractCommand(rawCommand string) string {

	command := rawCommand

	// remove slash if necessary
	if rawCommand[0] == '/' {
		command = command[1:]
	}

	// if command contains the name of our bot, remote it
	command = strings.Split(command, "@")[0]
	command = strings.Split(command, " ")[0]

	return command
}

// simply send a message to bot in Markdown format
func sendMsg(bot *tgbotapi.BotAPI, chatID int64, textMarkdown string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, textMarkdown)
	msg.ParseMode = "Markdown"
	msg.DisableWebPagePreview = true

	// send the message
	resp, err := bot.Send(msg)
	if err != nil {
		log.Println("bot.Send:", err, resp)
		return resp, err
	}

	return resp, err
}

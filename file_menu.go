package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getFileMenu(filename string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Reload", getFileButtonData(filename)),
			tgbotapi.NewInlineKeyboardButtonData("Switch to anothre file", getCancelButtonData()),
		),
	)
}

func getFileMessage(data string) (string, tgbotapi.InlineKeyboardMarkup) {
	if !strings.HasPrefix(data, "f:") {
		log.Fatalf("bad data %s", data)
	}

	filename := data[2:]

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to load content from filename %s \n\n %s", filename, err)
	}

	return string(content), getFileMenu(filename)
}

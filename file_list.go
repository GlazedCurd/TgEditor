package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getFilesListKeyboard() tgbotapi.InlineKeyboardMarkup {
	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	mkup := make([][]tgbotapi.InlineKeyboardButton, 0, len(files))

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := file.Name()
		mkup = append(
			mkup,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(filename, getFileButtonData(filename)),
			),
		)
	}
	return tgbotapi.NewInlineKeyboardMarkup(mkup...)

}

func getFilesListMessage() (string, tgbotapi.InlineKeyboardMarkup) {
	return "choose file for rendering:", getFilesListKeyboard()
}

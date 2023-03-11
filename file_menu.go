package main

import (
	"log"
	"os"
	"path/filepath"
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
		// По построению мы не должны сюда не попадать никогда
		log.Fatalf("bad command %s", data)
	}

	log.Printf("Processing command %s", data)

	filename := data[2:]

	currentPath, err := os.Getwd()
	if err != nil {
		log.Printf("Failed to get current working directory")
		return "Internal error", getFileMenu(filename)
	}

	fileAbsPath, err := filepath.Abs(filename)
	if err != nil {
		log.Printf("Failed to get current file abs path, %s", err)
		return "Failed to get abs filepath", getFileMenu(filename)
	}

	relFileName, err := filepath.Rel(currentPath, fileAbsPath)
	if err != nil {
		log.Printf("Failed to get current working directory, %s", err)
		return "Failed to get rel filepath", getFileMenu(filename)
	}

	if !filepath.IsLocal(relFileName) {
		return "Non local filepath", getFileMenu(filename)
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("failed to load content from filename %s \n\n %s", filename, err)
		return "Failed to get file content", getFileMenu(filename)
	}

	return string(content), getFileMenu(filename)
}

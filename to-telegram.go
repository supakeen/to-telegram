package main

import (
    "os"

	"github.com/joho/godotenv"
)


type ApplicationConfiguration struct {
	TelegramChatID   string
	TelegramBotToken string
}

func init() {
	godotenv.Load()
}

func main() {
	conf := ApplicationConfiguration{
		TelegramChatID:   os.Getenv("TELEGRAM_CHAT_ID"),
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}
}

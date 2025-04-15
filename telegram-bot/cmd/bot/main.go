package main

import (
	"log"

	"telegram-bot/internal/bot"
	"telegram-bot/internal/config"
)

func main() {
	cfg := config.Load()
	
	bot, err := bot.New(cfg)
	if err != nil {
		log.Panic(err)
	}
	
	bot.Start()
}
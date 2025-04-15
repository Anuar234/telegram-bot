package bot

import (
	"log"

	"telegram-bot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot wraps the Telegram bot
type Bot struct {
	API    *tgbotapi.BotAPI
	Config *config.Config
}

// New creates a new bot instance
func New(config *config.Config) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		return nil, err
	}

	botAPI.Debug = config.Debug
	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	return &Bot{
		API:    botAPI,
		Config: config,
	}, nil
}

// Start begins polling for updates
func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.API.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.HandleCommand(update)
			continue
		}

		b.HandleMessage(update)
	}
}
package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleMessage processes regular messages
func (b *Bot) HandleMessage(update tgbotapi.Update) {
	switch update.Message.Text {
	case "Help":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Available commands:\n"+
			"/start - Start the bot\n"+
			"/help - Show this help message\n"+
			"/keyboard - Show keyboard options\n"+
			"/echo [text] - Echo back the text")
		b.API.Send(msg)
	case "About":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "This is a simple Telegram bot created with Go!")
		b.API.Send(msg)
	default:
		// Echo all other text messages
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You said: "+update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		b.API.Send(msg)
	}
	
}
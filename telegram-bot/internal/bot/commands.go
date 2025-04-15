package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleCommand processes bot commands
func (b *Bot) HandleCommand(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	
	command := update.Message.Command()
	
	switch command {
	case "start":
		msg.Text = "Hello! I'm your Telegram bot. Try /help to see what I can do."
	case "help":
		msg.Text = "Available commands:\n" +
			"/start - Start the bot\n" +
			"/help - Show this help message\n" +
			"/keyboard - Show keyboard options\n" +
			"/echo [text] - Echo back the text"
	case "keyboard":
		b.ShowMainKeyboard(update.Message.Chat.ID)
		return
	case "echo":
		args := update.Message.CommandArguments()
		if args == "" {
			msg.Text = "Please provide some text to echo!"
		} else {
			msg.Text = args
		}
	default:
		msg.Text = "Unknown command. Try /help to see available commands."
	}

	if _, err := b.API.Send(msg); err != nil {
		log.Println(err)
	}
}

// ShowMainKeyboard displays the main keyboard to the user
func (b *Bot) ShowMainKeyboard(chatID int64) {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Help"),
			tgbotapi.NewKeyboardButton("About"),
		),
	)
	
	msg := tgbotapi.NewMessage(chatID, "Choose an option:")
	msg.ReplyMarkup = keyboard
	
	if _, err := b.API.Send(msg); err != nil {
		log.Println(err)
	}
}

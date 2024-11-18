package main

import (
	"github.com/joho/godotenv"
	tb "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func main() {
	// Get your bot token from environment variables or replace it directly
	err := godotenv.Load()
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Configure the bot with a poller for long polling
	bot, err := tb.NewBot(tb.Settings{
		Token:  botToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalf("Failed to create Telegram bot: %v", err)
	}
	webAppButton := tb.Btn{
		Text: "Open Mini App",
		WebApp: &tb.WebApp{
			URL: "https://f-bot.click", // Replace with the actual URL of your mini application
		},
	}

	// Create a reply markup with the Web App button
	inlineMenu := &tb.ReplyMarkup{}
	inlineMenu.Inline(
		inlineMenu.Row(webAppButton),
	)

	// Handle /start command to display the Web App button
	bot.Handle("/start", func(c tb.Context) error {
		return c.Send("Click the button below to open the mini application within Telegram:", inlineMenu)
	})

	// Start the bot to listen for messages
	log.Println("Starting bot with long polling...")
	bot.Start()
}

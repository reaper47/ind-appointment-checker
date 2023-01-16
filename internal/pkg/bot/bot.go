package bot

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var bot config

type config struct {
	URL string
}

// Init initializes and verifies the bot.
func Init() error {
	chatID := os.Getenv("TELEGRAM_CHATID")

	_, err := strconv.Atoi(chatID[1:])
	if err != nil {
		return fmt.Errorf("chatID must be an integer")
	}

	botID := os.Getenv("TELEGRAM_BOTID")
	parts := strings.Split(botID, ":")
	if len(parts) != 2 {
		return fmt.Errorf("botID must be separated by a colon (:)")
	}

	_, err = strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("left part of botID must be an integer")
	}

	bot = config{
		URL: "https://api.telegram.org/bot" + botID + "/sendMessage?chat_id=" + chatID + "&text=",
	}
	return nil
}

// SendMessage sends a Telegram message to the bot chat.
func SendMessage(text string) error {
	_, err := http.Get(bot.URL + strings.ReplaceAll(text, " ", "+"))
	if err != nil {
		return fmt.Errorf("could not send message: %s", err)
	}

	log.Printf("sent message: %s", strings.ReplaceAll(text, "\n", " "))
	return nil
}

// Clear resets the bot's configuration.
func Clear() {
	bot = config{}
}

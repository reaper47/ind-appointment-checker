package bot

import (
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	bot  config
	once sync.Once
)

type config struct {
	URL string
}

func (c config) message(text string) {
	_, err := http.Get(c.URL + strings.ReplaceAll(text, " ", "+"))
	if err != nil {
		log.Println("could not send message: ", err)
	}
}

// SendMessage sends a Telegram message to the bot chat.
func SendMessage(text string) {
	once.Do(func() {
		bot = config{
			URL: "https://api.telegram.org/bot" +
				os.Getenv("TELEGRAM_BOTID") +
				"/sendMessage?chat_id=" +
				os.Getenv("TELEGRAM_CHATID") + "&text=",
		}
	})
	bot.message(text)
}

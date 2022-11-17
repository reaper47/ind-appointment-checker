package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/reaper47/ind-appointment-checker/internal/pkg/bot"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/client"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/config"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/constants"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/jobs"

	"github.com/joho/godotenv"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		log.Fatalf("could not get working dir: %q", err)
	}

	err = godotenv.Load(filepath.Dir(exe) + "/.env")
	if err != nil {
		log.Fatalf("error loading .env file: %q", err)
	}
	config.Init()

	err = bot.Init()
	if err != nil {
		log.Fatalf("error initializing bot: %q", err)
	}

	var isReceived string
	bot.SendMessage("Hello, your bot is set up correctly.")
	for {
		fmt.Print("Did your bot receive a test message? [Y/n]: ")
		fmt.Scanln(&isReceived)

		if isReceived == "" {
			break
		}

		firstLetter := []rune(strings.ToLower(isReceived))[0]
		if firstLetter == 'y' {
			break
		}

		if firstLetter == 'n' {
			log.Println("Program aborted. Please verify your bot and retry.")
			log.Fatalln("If the problem persists, please send a message to macpoule@gmail.com with your.env file.")
		}
	}
	fmt.Printf("The IND appointment checker is running hot. Enjoy :-)\n\n")

	c := client.NewClient(constants.BaseURL)
	jobs.ScheduleCronJobs(c)
}

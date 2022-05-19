package main

import (
	"github.com/reaper47/ind-appointment-checker/internal/pkg/client"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/config"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/constants"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/jobs"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal("could not get working dir:", err)
	}

	err = godotenv.Load(filepath.Dir(exe) + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	config.Init()

	c := client.NewClient(constants.BaseURL)
	jobs.ScheduleCronJobs(c)
}

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/reaper47/ind-appointment-checker/internal/jobs"
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

	jobs.ScheduleCronJobs()
}

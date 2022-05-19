package config

import (
	"os"
	"time"
)

// Config is the configuration singleton for use throughout the application.
var Config config

type config struct {
	StartDate                       time.Time
	CurrAppointmentBiometrics       time.Time
	CurrAppointmentResidenceSticker time.Time
	CurrAppointmentResidenceCard    time.Time
}

// Init initializes the Config struct with the environment variables from the .env file.
func Init() {
	startDate, err := time.Parse("02/01/2006", os.Getenv("IND_START_DATE"))
	if err != nil {
		startDate = time.Now()
	}

	biometrics, err := time.Parse("02/01/2006", os.Getenv("IND_CURRENT_APPOINTMENT_BIOMETRICS"))
	if err != nil {
		biometrics = time.Now().Add(1460 * time.Hour)
	}

	sticker, err := time.Parse("02/01/2006", os.Getenv("IND_CURRENT_APPOINTMENT_RESIDENCE_STICKER"))
	if err != nil {
		biometrics = time.Now().Add(1460 * time.Hour)
	}

	card, err := time.Parse("02/01/2006", os.Getenv("IND_CURRENT_APPOINTMENT_RESIDENCE_CARD"))
	if err != nil {
		biometrics = time.Now().Add(1460 * time.Hour)
	}

	Config = config{
		StartDate:                       startDate,
		CurrAppointmentBiometrics:       biometrics,
		CurrAppointmentResidenceSticker: sticker,
		CurrAppointmentResidenceCard:    card,
	}
}

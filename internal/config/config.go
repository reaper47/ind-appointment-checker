package config

import (
	"os"
	"sync"
	"time"
)

var (
	c    config
	once sync.Once
)

type config struct {
	StartDate                       time.Time
	CurrAppointmentBiometrics       time.Time
	CurrAppointmentResidenceSticker time.Time
}

func Config() config {
	once.Do(func() {
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

		c = config{
			StartDate:                       startDate,
			CurrAppointmentBiometrics:       biometrics,
			CurrAppointmentResidenceSticker: sticker,
		}
	})
	return c
}

package jobs

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/reaper47/ind-appointment-checker/internal/appointments"
	"github.com/reaper47/ind-appointment-checker/internal/bot"
	"github.com/reaper47/ind-appointment-checker/internal/config"
	"github.com/reaper47/ind-appointment-checker/internal/models"
	"github.com/reaper47/ind-appointment-checker/internal/repository"
)

// ScheduleCronJobs schedules cron jobs for the app. It starts the following jobs:
//
// - watchAppointments: warch for earlier appointments every 10m
func ScheduleCronJobs() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(10).Minutes().Do(func() {
		watchAppointments()
	})
	s.StartBlocking()
}

func watchAppointments() {
	checkAvailabilities(appointments.Biometrics(), config.Config().CurrAppointmentBiometrics, true)
	checkAvailabilities(appointments.ResidenceSticker(), config.Config().CurrAppointmentResidenceSticker, false)
}

func checkAvailabilities(availabilities []models.Availabilities, currAppointmentDate time.Time, isBiometrics bool) {
	for _, avail := range availabilities {
		checkDates(avail.Data, avail.City, currAppointmentDate, isBiometrics)
	}
}

func checkDates(appointments []models.Availability, city models.City, currDate time.Time, isBiometrics bool) {
	for _, appointment := range appointments {
		date, err := time.Parse("2006-01-02", appointment.Date)
		if err != nil {
			log.Printf("could not parse %q for %q: %s", appointment.Date, city, err)
			continue
		}
		startTime, _ := time.Parse("15:04", appointment.StartTime)

		var containsDate bool
		if isBiometrics {
			containsDate = repository.Repo().ContainsBiometricsDate(city, date)
		} else {
			containsDate = repository.Repo().ContainsResidenceStickerDate(city, date)
		}

		if !containsDate && date.After(config.Config().StartDate) && date.Before(currDate) {
			if isBiometrics {
				repository.Repo().AddBiometricDate(city, date)
			} else {
				repository.Repo().AddResidenceStickerDate(city, date)
			}

			key := "VAA"
			name := "residence sticker"
			if isBiometrics {
				key = "BIO"
				name = "biometrics"
			}

			text := fmt.Sprintf("%s:\nAn earlier %s appointment is available on %s at %s.\nBook an appointment now: https://oap.ind.nl/oap/en/#/%s", city, name, date.Format("02 Jan 2006"), startTime.Format("15:04"), key)
			text = strings.ReplaceAll(text, "/#", "/%23")
			bot.SendMessage(strings.ReplaceAll(text, "\n", "%0A"))
			break
		}
	}
}

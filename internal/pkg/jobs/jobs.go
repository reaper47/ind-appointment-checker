package jobs

import (
	"fmt"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/appointments"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/bot"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/client"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/config"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/constants"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/models"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/repository"
	"log"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
)

type allURLs struct {
	biometrics       []models.URL
	residenceSticker []models.URL
	residenceCard    []models.URL
}

// ScheduleCronJobs schedules cron jobs for the app. It starts the following jobs:
//
// - watchAppointments: watch for earlier appointments every 10m
func ScheduleCronJobs(c client.Client) {
	s := gocron.NewScheduler(time.UTC)

	urls := allURLs{
		biometrics:       appointments.Biometrics(),
		residenceSticker: appointments.ResidenceSticker(),
		residenceCard:    appointments.ResidenceCard(),
	}
	_, _ = s.Every(10).Minutes().Do(func() {
		watchAppointments(c, urls)
	})

	s.StartBlocking()
}

func watchAppointments(c client.Client, urls allURLs) {
	xa := appointments.Process(c, urls.biometrics)
	checkAvailabilities(xa, config.Config.CurrAppointmentBiometrics, constants.ProductKeyBiometrics)

	xa = appointments.Process(c, urls.residenceSticker)
	checkAvailabilities(xa, config.Config.CurrAppointmentResidenceSticker, constants.ProductKeyResidenceSticker)

	xa = appointments.Process(c, urls.residenceCard)
	checkAvailabilities(xa, config.Config.CurrAppointmentResidenceCard, constants.ProductKeyResidenceCard)
}

func checkAvailabilities(availabilities []models.Availabilities, currAppointmentDate time.Time, productKey string) {
	for _, avail := range availabilities {
		checkDates(avail.Data, avail.City, currAppointmentDate, productKey)
	}
}

func checkDates(availabilities []models.Availability, city models.City, currDate time.Time, productKey string) {
	for _, availability := range availabilities {
		date, err := time.Parse("2006-01-02", availability.Date)
		if err != nil {
			log.Printf("could not parse %q for %q: %s", availability.Date, city, err)
			continue
		}
		startTime, _ := time.Parse("15:04", availability.StartTime)

		var containsDate bool
		switch productKey {
		case constants.ProductKeyBiometrics:
			containsDate = repository.Repo().ContainsBiometricsDate(city, date)
		case constants.ProductKeyResidenceSticker:
			containsDate = repository.Repo().ContainsResidenceStickerDate(city, date)
		case constants.ProductKeyResidenceCard:
			containsDate = repository.Repo().ContainsResidenceCardDate(city, date)
		}

		if !containsDate && date.After(config.Config.StartDate) && date.Before(currDate) {
			switch productKey {
			case constants.ProductKeyBiometrics:
				repository.Repo().AddBiometricDate(city, date)
			case constants.ProductKeyResidenceSticker:
				repository.Repo().AddResidenceStickerDate(city, date)
			case constants.ProductKeyResidenceCard:
				repository.Repo().AddResidenceCardDate(city, date)
			}

			var name string
			switch productKey {
			case constants.ProductKeyBiometrics:
				name = "biometrics"
			case constants.ProductKeyResidenceSticker:
				name = "residence sticker"
			case constants.ProductKeyResidenceCard:
				name = "residence card"
			}

			text := fmt.Sprintf("%s:\nAn earlier %s appointment is available on %s at %s.\nBook an appointment now: https://oap.ind.nl/oap/en/#/%s", city, name, date.Format("02 Jan 2006"), startTime.Format("15:04"), productKey)
			text = strings.ReplaceAll(text, "/#", "/%23")
			bot.SendMessage(strings.ReplaceAll(text, "\n", "%0A"))
			break
		}
	}
}

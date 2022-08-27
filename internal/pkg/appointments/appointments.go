package appointments

import (
	"github.com/reaper47/ind-appointment-checker/internal/pkg/client"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/config"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/constants"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/models"
	"golang.org/x/exp/slices"
	"log"
	"time"
)

// Biometrics prepares all URLs for the biometrics appointment.
func Biometrics() []models.URL {
	return makeURLs(constants.ProductKeyBiometrics, models.BiometricCities, config.Config.Persons)
}

// ResidenceSticker prepares all URLs for the residence sticker appointment.
func ResidenceSticker() []models.URL {
	return makeURLs(constants.ProductKeyResidenceSticker, models.ResidenceStickerCities, config.Config.Persons)
}

// ResidenceCard prepares all URLs for the residence card collection appointment.
func ResidenceCard() []models.URL {
	return makeURLs(constants.ProductKeyResidenceCard, models.ResidenceCardCities, config.Config.Persons)
}

func makeURLs(productKey string, cities map[string]models.City, persons int) []models.URL {
	if len(config.Config.Cities) > 0 {
		cities = config.Config.Cities
	}

	xr := make([]models.URL, len(cities))
	i := 0
	for _, city := range cities {
		xr[i] = models.NewURL(city, productKey, persons)
		i++
	}
	return xr
}

// Process fetches availabilities for every URL.
// A delay of 200ms is set between HTTP calls to ensure a light load on IND's servers.
func Process(c client.Client, urls []models.URL) []models.Availabilities {
	var xa []models.Availabilities
	for _, u := range urls {
		a, err := u.Process(c)
		if err != nil {
			log.Printf("error processing %s: %s", u.Endpoint, err)
			continue
		}
		xa = append(xa, a)
		time.Sleep(200 * time.Millisecond)
	}

	for _, availabilities := range xa {
		slices.SortFunc(availabilities.Data, func(a models.Availability, b models.Availability) bool {
			aDate, err := time.Parse("2006-01-02", a.Date)
			if err != nil {
				return false
			}

			bDate, err := time.Parse("2006-01-02", b.Date)
			if err != nil {
				return false
			}

			return aDate.Before(bDate)
		})
	}

	return xa
}

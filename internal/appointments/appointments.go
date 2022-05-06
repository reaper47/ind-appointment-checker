package appointments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/reaper47/ind-appointment-checker/internal/models"
)

// Biometrics fetches all availabilities for the biometrics appointment.
func Biometrics() []models.Availabilities {
	urls := getURLs(true)
	return processURLs(urls)
}

// ResidenceSticker fetches all availabilities for the residence sticker appointment.
func ResidenceSticker() []models.Availabilities {
	urls := getURLs(false)
	return processURLs(urls)
}

func getURLs(isBiometrics bool) []models.URL {
	cities := models.ResidenceStickerCities
	if isBiometrics {
		cities = models.BiometricCities
	}

	var urls []models.URL
	for abbr, city := range cities {
		urls = append(urls, makeURL(abbr, city, isBiometrics))
	}
	return urls
}

func makeURL(cityAbbr string, city models.City, isBiometrics bool) models.URL {
	productKey := "VAA"
	if isBiometrics {
		productKey = "BIO"
	}
	return models.URL{
		City: city,
		Url:  fmt.Sprintf("https://oap.ind.nl/oap/api/desks/%s/slots/?productKey=%s&persons=1", cityAbbr, productKey),
	}
}

func processURLs(urls []models.URL) []models.Availabilities {
	var availabilities []models.Availabilities
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(u models.URL) {
			defer wg.Done()

			res, err := http.Get(u.Url)
			if err != nil {
				fmt.Printf("could not process %s: %s", u, err)
				return
			}

			xb, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Printf("could not read body for %s: %s", u, err)
				return
			}
			xb = bytes.ReplaceAll(xb, []byte(")]}',"), []byte(""))

			var availability models.Availabilities
			err = json.Unmarshal(xb, &availability)
			if err != nil {
				log.Fatalln(err)
			}
			availability.City = u.City

			availabilities = append(availabilities, availability)
		}(url)
	}

	wg.Wait()
	return availabilities
}

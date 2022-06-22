package repository

import (
	"sync"
	"time"

	"github.com/reaper47/ind-appointment-checker/internal/pkg/constants"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/models"
)

var (
	repo *repository
	once sync.Once
)

type repository struct {
	sentBiometricsDates       map[models.City][]time.Time
	sentResidenceStickerDates map[models.City][]time.Time
	sentResidenceCardDates    map[models.City][]time.Time
}

// Repo initialises and returns the application's in-memory repository.
func Repo() *repository {
	once.Do(func() {
		repo = &repository{
			sentBiometricsDates:       make(map[models.City][]time.Time),
			sentResidenceStickerDates: make(map[models.City][]time.Time),
			sentResidenceCardDates:    make(map[models.City][]time.Time),
		}
	})
	return repo
}

// AddDate adds an appointment date for the given city and product type.
// The goal is to store dates already sent to the user via Telegram to avoid
// sending the same appointment twice.
func (r *repository) AddDate(productKey string, city models.City, date time.Time) {
	switch productKey {
	case constants.ProductKeyBiometrics:
		r.sentBiometricsDates[city] = append(r.sentBiometricsDates[city], date)
	case constants.ProductKeyResidenceSticker:
		r.sentResidenceStickerDates[city] = append(r.sentResidenceStickerDates[city], date)
	case constants.ProductKeyResidenceCard:
		r.sentResidenceCardDates[city] = append(r.sentResidenceCardDates[city], date)
	}
}

// ContainsBiometricsDate checks whether the appointment for the
// product type lies in the repository.
func (r *repository) ContainsDate(productKey string, city models.City, date time.Time) bool {
	switch productKey {
	case constants.ProductKeyBiometrics:
		return findDate(r.sentBiometricsDates, city, date)
	case constants.ProductKeyResidenceSticker:
		return findDate(r.sentResidenceStickerDates, city, date)
	case constants.ProductKeyResidenceCard:
		return findDate(r.sentResidenceCardDates, city, date)
	default:
		return false
	}
}

func findDate(m map[models.City][]time.Time, city models.City, date time.Time) bool {
	dates, ok := m[city]
	if !ok {
		return false
	}

	for _, d := range dates {
		if d == date {
			return true
		}
	}
	return false
}

package repository

import (
	"sync"
	"time"

	"github.com/reaper47/ind-appointment-checker/internal/models"
)

var (
	repo *repository
	once sync.Once
)

// Repo initialises and returns the application's in-memory repository.
func Repo() *repository {
	once.Do(func() {
		repo = &repository{
			sentBiometricsDates:       make(map[models.City][]time.Time),
			sentResidenceStickerDates: make(map[models.City][]time.Time),
		}
	})
	return repo
}

type repository struct {
	sentBiometricsDates       map[models.City][]time.Time
	sentResidenceStickerDates map[models.City][]time.Time
}

// AddBiometricDate adds a biometrics appointment date for the given city.
// The goal is to store dates already sent to the user via Telegram to avoid
// sending the same appointment twice.
func (r *repository) AddBiometricDate(city models.City, date time.Time) {
	r.sentBiometricsDates[city] = append(r.sentBiometricsDates[city], date)
}

// AddResidenceStickerDate adds a biometrics appointment date for the given city.
// The goal is to store dates already sent to the user via Telegram to avoid
// sending the same appointment twice.
func (r *repository) AddResidenceStickerDate(city models.City, date time.Time) {
	r.sentResidenceStickerDates[city] = append(r.sentResidenceStickerDates[city], date)
}

// ContainsBiometricsDate checks whether the biometrics appointment
// lies in the repository.
func (r *repository) ContainsBiometricsDate(city models.City, date time.Time) bool {
	return findDate(r.sentBiometricsDates, city, date)
}

// ContainsResidenceStickerDate checks whether the residence sticker
// apppintment lies in the repository.
func (r *repository) ContainsResidenceStickerDate(city models.City, date time.Time) bool {
	return findDate(r.sentResidenceStickerDates, city, date)
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

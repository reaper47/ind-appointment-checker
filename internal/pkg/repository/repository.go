package repository

import (
	"github.com/reaper47/ind-appointment-checker/internal/pkg/models"
	"sync"
	"time"
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
			sentResidenceCardDates:    make(map[models.City][]time.Time),
		}
	})
	return repo
}

type repository struct {
	sentBiometricsDates       map[models.City][]time.Time
	sentResidenceStickerDates map[models.City][]time.Time
	sentResidenceCardDates    map[models.City][]time.Time
}

// AddBiometricDate adds a biometrics appointment date for the given city.
// The goal is to store dates already sent to the user via Telegram to avoid
// sending the same appointment twice.
func (r *repository) AddBiometricDate(city models.City, date time.Time) {
	r.sentBiometricsDates[city] = append(r.sentBiometricsDates[city], date)
}

// AddResidenceStickerDate adds a residence sticker appointment date for the given city.
// The goal is to store dates already sent to the user via Telegram to avoid
// sending the same appointment twice.
func (r *repository) AddResidenceStickerDate(city models.City, date time.Time) {
	r.sentResidenceStickerDates[city] = append(r.sentResidenceStickerDates[city], date)
}

// AddResidenceCardDate adds a residence card collection appointment date for the given city.
// The goal is to store dates already sent to the user via Telegram to avoid
// sending the same appointment twice.
func (r *repository) AddResidenceCardDate(city models.City, date time.Time) {
	r.sentResidenceCardDates[city] = append(r.sentResidenceCardDates[city], date)
}

// ContainsBiometricsDate checks whether the biometrics appointment
// lies in the repository.
func (r *repository) ContainsBiometricsDate(city models.City, date time.Time) bool {
	return findDate(r.sentBiometricsDates, city, date)
}

// ContainsResidenceStickerDate checks whether the residence sticker
// appointment lies in the repository.
func (r *repository) ContainsResidenceStickerDate(city models.City, date time.Time) bool {
	return findDate(r.sentResidenceStickerDates, city, date)
}

// ContainsResidenceCardDate checks whether the residence sticker
// appointment lies in the repository.
func (r *repository) ContainsResidenceCardDate(city models.City, date time.Time) bool {
	return findDate(r.sentResidenceCardDates, city, date)
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

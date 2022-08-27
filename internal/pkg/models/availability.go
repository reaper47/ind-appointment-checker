package models

import "golang.org/x/exp/slices"

// Availabilities is a representation of the response
// received from IND when after selecting a city.
//
// The City and Persons fields have been added for convenience to
// associate the availabilities to the city and the number of people.
type Availabilities struct {
	City    City
	Persons int
	Status  string         `json:"status"`
	Data    []Availability `json:"data"`
}

// Equal verifies the Availabilities structs are identical.
func (a Availabilities) Equal(other Availabilities) bool {
	return a.City == other.City && a.Persons == other.Persons && a.Status == other.Status && slices.Equal(a.Data, other.Data)
}

// Availability is a representation of the response.data field of
// the response received from IND when after selecting a city.
type Availability struct {
	Key       string `json:"key"`
	Date      string `json:"date"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Parts     int    `json:"parts"`
}

// Equal verifies the Availability structs are identical.
func (a Availability) Equal(other Availability) bool {
	return a == other
}

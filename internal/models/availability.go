package models

// Availabilities is a representation of the response
// received from IND when after selecting a city.
//
// The City field has been added for convenience to
// associate the availabilities to the city.
type Availabilities struct {
	City   City
	Status string         `json:"status"`
	Data   []Availability `json:"data"`
}

// Availability is a representation of the response.data field of the
// the response received from IND when after selecting a city.
type Availability struct {
	Key       string `json:"key"`
	Date      string `json:"date"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Parts     int    `json:"parts"`
}

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/client"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/constants"
	"io/ioutil"
	"net/http"
)

// URL holds a URL and the name of the city the URL is associated with.
type URL struct {
	City       City
	Endpoint   string
	ProductKey string
	Persons    int
}

// Equal verifies the URLs are identical.
func (u URL) Equal(other URL) bool {
	return u.City == other.City && u.Endpoint == other.Endpoint && u.Persons == other.Persons
}

// Process fetches appointments for the URL struct.
func (u URL) Process(c client.Client) (Availabilities, error) {
	res, err := c.MakeRequest(u.Endpoint)
	if err != nil {
		return Availabilities{}, fmt.Errorf("could not process %s: %s", u.Endpoint, err)
	}

	if res.StatusCode != http.StatusOK {
		return Availabilities{}, fmt.Errorf("got status %d (%s) instead of 200", res.StatusCode, http.StatusText(res.StatusCode))
	}

	xb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Availabilities{}, fmt.Errorf("could not read body for %s: %s", u.Endpoint, err)
	}
	xb = bytes.ReplaceAll(xb, []byte(")]}',"), []byte(""))

	var availability Availabilities
	err = json.Unmarshal(xb, &availability)
	if err != nil {
		return Availabilities{}, fmt.Errorf("unable to unmarshal json for %s: %s", u.Endpoint, err)
	}
	availability.City = u.City
	availability.Persons = u.Persons
	return availability, nil
}

// NewURL creates a Req model from the input City and product key.
func NewURL(city City, productKey string, persons int) URL {
	if persons <= 0 || persons > constants.MaxPersons {
		persons = 1
	}

	return URL{
		City:       city,
		Endpoint:   fmt.Sprintf("/%s/slots/?productKey=%s&persons=%d", city.Abbrev(), productKey, persons),
		ProductKey: productKey,
		Persons:    persons,
	}
}

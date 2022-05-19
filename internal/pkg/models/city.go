package models

// City is an abstraction over string for the name of a city.
type City string

// Abbrev abbreviates the city to its initials.
func (c City) Abbrev() string {
	switch c.String() {
	case Amsterdam.String():
		return "AM"
	case TheHague.String():
		return "DH"
	case Rotterdam.String():
		return "RO"
	case Zwolle.String():
		return "ZW"
	case DenBosch.String():
		return "DB"
	case Haarlem.String():
		return "6b425ff9f87de136a36b813cccf26e23"
	case Utrecht.String():
		return "fa24ccf0acbc76a7793765937eaee440"
	case ExpatEnschede.String():
		return "3535aca0fb9a2e8e8015f768fb3fa69d"
	case ExpatRotterdam.String():
		return "f0ef3c8f0973875936329d713a68c5f3"
	default:
		return c.String()
	}
}

// String provides the string representation of the City.
func (c City) String() string {
	return string(c)
}

// Constants for various cities in The Netherlands.
const (
	Amsterdam      = City("Amsterdam")
	TheHague       = City("The Hague")
	Rotterdam      = City("Rotterdam")
	Zwolle         = City("Zwolle")
	DenBosch       = City("Den Bosch")
	Haarlem        = City("IND Haarlem")
	Utrecht        = City("Expat center Utrecht")
	ExpatRotterdam = City("Expat center Rotterdam")
	ExpatEnschede  = City("Expat center Enschede")
)

// BiometricCities is the list of cities that offer appointments for biometrics.
var BiometricCities = map[string]City{
	"AM":                               Amsterdam,
	"DH":                               TheHague,
	"RO":                               Rotterdam,
	"ZW":                               Zwolle,
	"DB":                               DenBosch,
	"6b425ff9f87de136a36b813cccf26e23": Haarlem,
	"fa24ccf0acbc76a7793765937eaee440": Utrecht,
	"3535aca0fb9a2e8e8015f768fb3fa69d": ExpatEnschede,
	"f0ef3c8f0973875936329d713a68c5f3": ExpatRotterdam,
}

// ResidenceStickerCities is the list of cities that offer
// appointments for the residence sticker.
var ResidenceStickerCities = map[string]City{
	"AM": Amsterdam,
	"DH": TheHague,
	"RO": Rotterdam,
	"ZW": Zwolle,
	"DB": DenBosch,
}

// ResidenceCardCities is the list of cities that offer
// appointments to collect the residence card.
var ResidenceCardCities = map[string]City{
	"AM": Amsterdam,
	"DH": TheHague,
	"RO": Rotterdam,
	"ZW": Zwolle,
	"DB": DenBosch,
}

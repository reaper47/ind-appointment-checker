package models

// City is an abstraction over string for the name of a city.
type City string

// Abbrev abbreviates the city to its initials.
func (c City) Abbrev() string {
	switch c.String() {
	case INDAmsterdam.String():
		return "AM"
	case INDDenHaag.String():
		return "DH"
	case INDZwolle.String():
		return "ZW"
	case INDDenBosch.String():
		return "DB"
	case INDRotterdam.String():
		return "RO"
	case INDHaarlem.String():
		return "6b425ff9f87de136a36b813cccf26e23"
	case ExpatGroningen.String():
		return "0c127eb6d9fe1ced413d2112305e75f6"
	case ExpatMaastricht.String():
		return "6c5280823686521552efe85094e607cf"
	case ExpatWageningen.String():
		return "b084907207cfeea941cd9698821fd894"
	case ExpatEindhoven.String():
		return "0588ef4088c08f53294eb60bab55c81e"
	case ExpatDenHaag.String():
		return "5e325f444aeb56bb0270a61b4a0403eb"
	case ExpatRotterdam.String():
		return "f0ef3c8f0973875936329d713a68c5f3"
	case ExpatEnschede.String():
		return "3535aca0fb9a2e8e8015f768fb3fa69d"
	case ExpatUtrecht.String():
		return "fa24ccf0acbc76a7793765937eaee440"
	case ExpatAmsterdam.String():
		return "284b189314071dcd571df5bb262a31db"
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
	INDAmsterdam    = City("IND Amsterdam")
	INDDenHaag      = City("IND Den Haag")
	INDZwolle       = City("IND Zwolle")
	INDDenBosch     = City("IND Den Bosch")
	INDRotterdam    = City("IND Rotterdam")
	INDHaarlem      = City("IND Haarlem")
	ExpatGroningen  = City("Expatcenter Groningen")
	ExpatMaastricht = City("Expatcenter Maastricht")
	ExpatWageningen = City("Expatcenter Wageningen")
	ExpatEindhoven  = City("Expatcenter Eidhoven")
	ExpatDenHaag    = City("Expatcenter Den Haag")
	ExpatRotterdam  = City("Expatcenter Rotterdam")
	ExpatEnschede   = City("Expatcenter Enschede")
	ExpatUtrecht    = City("Expatcenter Utrecht")
	ExpatAmsterdam  = City("Expatcenter Amsterdam")
)

// BiometricCities is the list of cities that offer appointments for biometrics.
var BiometricCities = map[string]City{
	"AM":                               INDAmsterdam,
	"DH":                               INDDenHaag,
	"ZW":                               INDZwolle,
	"DB":                               INDDenBosch,
	"RO":                               INDRotterdam,
	"6b425ff9f87de136a36b813cccf26e23": INDHaarlem,
	"0c127eb6d9fe1ced413d2112305e75f6": ExpatGroningen,
	"6c5280823686521552efe85094e607cf": ExpatMaastricht,
	"b084907207cfeea941cd9698821fd894": ExpatWageningen,
	"0588ef4088c08f53294eb60bab55c81e": ExpatEindhoven,
	"5e325f444aeb56bb0270a61b4a0403eb": ExpatDenHaag,
	"f0ef3c8f0973875936329d713a68c5f3": ExpatRotterdam,
	"3535aca0fb9a2e8e8015f768fb3fa69d": ExpatEnschede,
	"fa24ccf0acbc76a7793765937eaee440": ExpatUtrecht,
	"284b189314071dcd571df5bb262a31db": ExpatAmsterdam,
}

// ResidenceStickerCities is the list of cities that offer
// appointments for the residence sticker.
var ResidenceStickerCities = map[string]City{
	"AM": INDAmsterdam,
	"DH": INDDenHaag,
	"ZW": INDZwolle,
	"DB": INDDenBosch,
}

// ResidenceCardCities is the list of cities that offer
// appointments to collect the residence card.
var ResidenceCardCities = map[string]City{
	"AM": INDAmsterdam,
	"DH": INDDenHaag,
	"ZW": INDZwolle,
	"DB": INDDenBosch,
}

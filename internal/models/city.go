package models

type City string

func (c City) String() string {
	return string(c)
}

const (
	Amsterdam = City("Amsterdam")
	TheHague  = City("The Hague")
	Rotterdam = City("Rotterdam")
	Zwolle    = City("Zwolle")
	DenBosch  = City("Den Bosch")
	Haarlem   = City("IND Haarlem")
	Utrecht   = City("Expatcenter Utrecht")
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

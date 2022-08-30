package constants

// Constants for the various IND product keys.
const (
	ProductKeyBiometrics       = "BIO"
	ProductKeyResidenceSticker = "VAA"
	ProductKeyResidenceCard    = "DOC"
)

// BaseURL is the base URL for IND's API.
const BaseURL = "https://oap.ind.nl/oap/api/desks"

// MaxPersons is the maximum number of persons for which an appointment can be
// requested.
const MaxPersons = 6

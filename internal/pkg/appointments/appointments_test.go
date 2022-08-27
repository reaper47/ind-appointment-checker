package appointments_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/reaper47/ind-appointment-checker/internal/pkg/appointments"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/client"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/config"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/models"
	"golang.org/x/exp/slices"
)

func TestBiometrics(t *testing.T) {
	t.Run("all cities", func(t *testing.T) {
		got := appointments.Biometrics()
		want := []models.URL{
			{
				City:       models.INDAmsterdam,
				Endpoint:   "/AM/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.INDDenHaag,
				Endpoint:   "/DH/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.INDZwolle,
				Endpoint:   "/ZW/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.INDDenBosch,
				Endpoint:   "/DB/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.INDRotterdam,
				Endpoint:   "/RO/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.INDHaarlem,
				Endpoint:   "/6b425ff9f87de136a36b813cccf26e23/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.ExpatGroningen,
				Endpoint:   "/0c127eb6d9fe1ced413d2112305e75f6/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.ExpatMaastricht,
				Endpoint:   "/6c5280823686521552efe85094e607cf/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.ExpatWageningen,
				Endpoint:   "/b084907207cfeea941cd9698821fd894/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.ExpatEindhoven,
				Endpoint:   "/0588ef4088c08f53294eb60bab55c81e/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.ExpatDenHaag,
				Endpoint:   "/5e325f444aeb56bb0270a61b4a0403eb/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.ExpatRotterdam,
				Endpoint:   "/f0ef3c8f0973875936329d713a68c5f3/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.ExpatEnschede,
				Endpoint:   "/3535aca0fb9a2e8e8015f768fb3fa69d/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.ExpatUtrecht,
				Endpoint:   "/fa24ccf0acbc76a7793765937eaee440/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.ExpatAmsterdam,
				Endpoint:   "/284b189314071dcd571df5bb262a31db/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
		}
		assertURLs(t, got, want)
	})

	t.Run("select cities", func(t *testing.T) {
		_ = os.Setenv("TARGET_CITIES", "IND Amsterdam, IND Den Haag,Expatcenter Enschede")
		defer func() {
			_ = os.Unsetenv("TARGET_CITIES")
		}()
		config.Init()

		got := appointments.Biometrics()
		want := []models.URL{
			{
				City:       models.INDAmsterdam,
				Endpoint:   "/AM/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.INDDenHaag,
				Endpoint:   "/DH/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
			{
				City:       models.ExpatEnschede,
				Endpoint:   "/3535aca0fb9a2e8e8015f768fb3fa69d/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
				Persons:    1,
			},
		}
		assertURLs(t, got, want)
	})
}

func TestResidenceSticker(t *testing.T) {
	t.Run("all cities", func(t *testing.T) {
		config.Init()

		got := appointments.ResidenceSticker()
		want := []models.URL{
			{
				City:       models.INDAmsterdam,
				Endpoint:   "/AM/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
				Persons:    1,
			},
			{
				City:       models.INDDenHaag,
				Endpoint:   "/DH/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
				Persons:    1,
			},
			{
				City:       models.INDZwolle,
				Endpoint:   "/ZW/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
				Persons:    1,
			},
			{
				City:       models.INDDenBosch,
				Endpoint:   "/DB/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
				Persons:    1,
			},
		}

		assertURLs(t, got, want)
	})

	t.Run("selected cities", func(t *testing.T) {
		_ = os.Setenv("TARGET_CITIES", "IND Amsterdam, IND Zwolle")
		defer func() {
			_ = os.Unsetenv("TARGET_CITIES")
		}()
		config.Init()

		got := appointments.ResidenceSticker()
		want := []models.URL{
			{
				City:       models.INDAmsterdam,
				Endpoint:   "/AM/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
				Persons:    1,
			},
			{
				City:       models.INDZwolle,
				Endpoint:   "/ZW/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
				Persons:    1,
			},
		}

		assertURLs(t, got, want)
	})
}

func TestResidence(t *testing.T) {
	t.Run("all cities", func(t *testing.T) {
		config.Init()

		got := appointments.ResidenceCard()
		want := []models.URL{
			{
				City:       models.INDAmsterdam,
				Endpoint:   "/AM/slots/?productKey=DOC&persons=1",
				ProductKey: "DOC",
				Persons:    1,
			},
			{
				City:       models.INDDenHaag,
				Endpoint:   "/DH/slots/?productKey=DOC&persons=1",
				ProductKey: "DOC",
				Persons:    1,
			},
			{
				City:       models.INDZwolle,
				Endpoint:   "/ZW/slots/?productKey=DOC&persons=1",
				ProductKey: "DOC",
				Persons:    1,
			},
			{
				City:       models.INDDenBosch,
				Endpoint:   "/DB/slots/?productKey=DOC&persons=1",
				ProductKey: "DOC",
				Persons:    1,
			},
		}

		assertURLs(t, got, want)
	})

	t.Run("selected cities", func(t *testing.T) {
		_ = os.Setenv("TARGET_CITIES", "IND Amsterdam, IND Zwolle")
		defer func() {
			_ = os.Unsetenv("TARGET_CITIES")
		}()
		config.Init()

		got := appointments.ResidenceCard()
		want := []models.URL{
			{
				City:       models.INDAmsterdam,
				Endpoint:   "/AM/slots/?productKey=DOC&persons=1",
				ProductKey: "VAA",
				Persons:    1,
			},
			{
				City:       models.INDZwolle,
				Endpoint:   "/ZW/slots/?productKey=DOC&persons=1",
				ProductKey: "DOC",
				Persons:    1,
			},
		}

		assertURLs(t, got, want)
	})
}

func TestPersons(t *testing.T) {
	t.Run("all cities", func(t *testing.T) {
		config.Init()

		got := appointments.ResidenceSticker()
		want := []models.URL{
			{
				City:       models.INDAmsterdam,
				Endpoint:   "/AM/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
				Persons:    1,
			},
			{
				City:       models.INDDenHaag,
				Endpoint:   "/DH/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
				Persons:    1,
			},
			{
				City:       models.INDZwolle,
				Endpoint:   "/ZW/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
				Persons:    1,
			},
			{
				City:       models.INDDenBosch,
				Endpoint:   "/DB/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
				Persons:    1,
			},
		}

		assertURLs(t, got, want)
	})

	t.Run("selected cities", func(t *testing.T) {
		_ = os.Setenv("TARGET_CITIES", "IND Amsterdam, IND Zwolle")
		_ = os.Setenv("TOTAL_PERSONS", "4")
		defer func() {
			_ = os.Unsetenv("TARGET_CITIES")
			_ = os.Unsetenv("TOTAL_PERSONS")
		}()
		config.Init()

		got := appointments.ResidenceSticker()
		want := []models.URL{
			{
				City:       models.INDAmsterdam,
				Endpoint:   "/AM/slots/?productKey=VAA&persons=4",
				ProductKey: "VAA",
				Persons:    4,
			},
			{
				City:       models.INDZwolle,
				Endpoint:   "/ZW/slots/?productKey=VAA&persons=4",
				ProductKey: "VAA",
				Persons:    4,
			},
		}

		assertURLs(t, got, want)
	})
}

func TestProcess(t *testing.T) {
	const productKey = "VAA"
	const persons = 6
	want := models.Availabilities{
		City:    models.INDAmsterdam,
		Persons: persons,
		Status:  http.StatusText(http.StatusOK),
		Data: []models.Availability{
			{
				Key:       productKey,
				Date:      "2022-02-06",
				StartTime: "9:15",
				EndTime:   "9:30",
				Parts:     1,
			},
		},
	}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xb, _ := json.Marshal(want)
		_, _ = w.Write(xb)
	}))
	defer svr.Close()
	c := client.NewClient(svr.URL)

	xu := []models.URL{models.NewURL(models.INDAmsterdam, productKey, persons)}
	got := appointments.Process(c, xu)

	if !got[0].Equal(want) {
		t.Fatalf("got %#v but want %#v", got[0], want)
	}
}

func assertURLs(t *testing.T, got, want []models.URL) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("got length %d but want %d", len(got), len(want))
	}

	slices.SortFunc(got, func(a models.URL, b models.URL) bool {
		return a.City < b.City
	})
	slices.SortFunc(want, func(a models.URL, b models.URL) bool {
		return a.City < b.City
	})

	for i := 0; i < len(got); i++ {
		if !got[i].Equal(want[i]) {
			t.Fatalf("got %#v but want %#v", got[i], want[i])
		}
	}
}

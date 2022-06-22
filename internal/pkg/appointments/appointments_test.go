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
				City:       models.Amsterdam,
				Endpoint:   "/AM/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
			},
			{
				City:       models.TheHague,
				Endpoint:   "/DH/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
			},
			{
				City:       models.Rotterdam,
				Endpoint:   "/RO/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
			},
			{
				City:       models.Zwolle,
				Endpoint:   "/ZW/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
			},
			{
				City:       models.DenBosch,
				Endpoint:   "/DB/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
			},
			{
				City:       models.Haarlem,
				Endpoint:   "/6b425ff9f87de136a36b813cccf26e23/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
			},
			{
				City:       models.Utrecht,
				Endpoint:   "/fa24ccf0acbc76a7793765937eaee440/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
			},
			{
				City:       models.ExpatEnschede,
				Endpoint:   "/3535aca0fb9a2e8e8015f768fb3fa69d/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
			},
			{
				City:       models.ExpatRotterdam,
				Endpoint:   "/f0ef3c8f0973875936329d713a68c5f3/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
			},
		}
		assertURLs(t, got, want)
	})

	t.Run("select cities", func(t *testing.T) {
		_ = os.Setenv("TARGET_CITIES", "Amsterdam, The Hague,Expat center Enschede")
		defer func() {
			_ = os.Unsetenv("TARGET_CITIES")
		}()
		config.Init()

		got := appointments.Biometrics()
		want := []models.URL{
			{
				City:       models.Amsterdam,
				Endpoint:   "/AM/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
			},
			{
				City:       models.TheHague,
				Endpoint:   "/DH/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
			},
			{
				City:       models.ExpatEnschede,
				Endpoint:   "/3535aca0fb9a2e8e8015f768fb3fa69d/slots/?productKey=BIO&persons=1",
				ProductKey: "BIO",
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
				City:       models.Amsterdam,
				Endpoint:   "/AM/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
			},
			{
				City:       models.TheHague,
				Endpoint:   "/DH/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
			},
			{
				City:       models.Rotterdam,
				Endpoint:   "/RO/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
			},
			{
				City:       models.Zwolle,
				Endpoint:   "/ZW/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
			},
			{
				City:       models.DenBosch,
				Endpoint:   "/DB/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
			},
		}

		assertURLs(t, got, want)
	})

	t.Run("selected cities", func(t *testing.T) {
		_ = os.Setenv("TARGET_CITIES", "Amsterdam, Rotterdam")
		defer func() {
			_ = os.Unsetenv("TARGET_CITIES")
		}()
		config.Init()

		got := appointments.ResidenceSticker()
		want := []models.URL{
			{
				City:       models.Amsterdam,
				Endpoint:   "/AM/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
			},
			{
				City:       models.Rotterdam,
				Endpoint:   "/RO/slots/?productKey=VAA&persons=1",
				ProductKey: "VAA",
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
				City:       models.Amsterdam,
				Endpoint:   "/AM/slots/?productKey=DOC&persons=1",
				ProductKey: "DOC",
			},
			{
				City:       models.TheHague,
				Endpoint:   "/DH/slots/?productKey=DOC&persons=1",
				ProductKey: "DOC",
			},
			{
				City:       models.Rotterdam,
				Endpoint:   "/RO/slots/?productKey=DOC&persons=1",
				ProductKey: "DOC",
			},
			{
				City:       models.Zwolle,
				Endpoint:   "/ZW/slots/?productKey=DOC&persons=1",
				ProductKey: "DOC",
			},
			{
				City:       models.DenBosch,
				Endpoint:   "/DB/slots/?productKey=DOC&persons=1",
				ProductKey: "DOC",
			},
		}

		assertURLs(t, got, want)
	})

	t.Run("selected cities", func(t *testing.T) {
		_ = os.Setenv("TARGET_CITIES", "Amsterdam, Zwolle")
		defer func() {
			_ = os.Unsetenv("TARGET_CITIES")
		}()
		config.Init()

		got := appointments.ResidenceCard()
		want := []models.URL{
			{
				City:       models.Amsterdam,
				Endpoint:   "/AM/slots/?productKey=DOC&persons=1",
				ProductKey: "VAA",
			},
			{
				City:       models.Zwolle,
				Endpoint:   "/ZW/slots/?productKey=DOC&persons=1",
				ProductKey: "DOC",
			},
		}

		assertURLs(t, got, want)
	})
}

func TestProcess(t *testing.T) {
	const productKey = "VAA"
	want := models.Availabilities{
		City:   models.Amsterdam,
		Status: http.StatusText(http.StatusOK),
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

	xu := []models.URL{models.NewURL(models.Amsterdam, productKey)}
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

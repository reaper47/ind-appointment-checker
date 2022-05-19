package appointments_test

import (
	"encoding/json"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/appointments"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/client"
	"github.com/reaper47/ind-appointment-checker/internal/pkg/models"
	"golang.org/x/exp/slices"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBiometrics(t *testing.T) {
	got := appointments.Biometrics()
	want := []models.URL{
		{models.Amsterdam, "/AM/slots/?productKey=BIO&persons=1", "BIO"},
		{models.TheHague, "/DH/slots/?productKey=BIO&persons=1", "BIO"},
		{models.Rotterdam, "/RO/slots/?productKey=BIO&persons=1", "BIO"},
		{models.Zwolle, "/ZW/slots/?productKey=BIO&persons=1", "BIO"},
		{models.DenBosch, "/DB/slots/?productKey=BIO&persons=1", "BIO"},
		{models.Haarlem, "/6b425ff9f87de136a36b813cccf26e23/slots/?productKey=BIO&persons=1", "BIO"},
		{models.Utrecht, "/fa24ccf0acbc76a7793765937eaee440/slots/?productKey=BIO&persons=1", "BIO"},
		{models.ExpatEnschede, "/3535aca0fb9a2e8e8015f768fb3fa69d/slots/?productKey=BIO&persons=1", "BIO"},
		{models.ExpatRotterdam, "/f0ef3c8f0973875936329d713a68c5f3/slots/?productKey=BIO&persons=1", "BIO"},
	}
	assertURLs(t, got, want)
}

func TestResidenceSticker(t *testing.T) {
	got := appointments.ResidenceSticker()
	want := []models.URL{
		{models.Amsterdam, "/AM/slots/?productKey=VAA&persons=1", "VAA"},
		{models.TheHague, "/DH/slots/?productKey=VAA&persons=1", "VAA"},
		{models.Rotterdam, "/RO/slots/?productKey=VAA&persons=1", "VAA"},
		{models.Zwolle, "/ZW/slots/?productKey=VAA&persons=1", "VAA"},
		{models.DenBosch, "/DB/slots/?productKey=VAA&persons=1", "VAA"},
	}
	assertURLs(t, got, want)
}

func TestResidence(t *testing.T) {
	got := appointments.ResidenceCard()
	want := []models.URL{
		{models.Amsterdam, "/AM/slots/?productKey=DOC&persons=1", "DOC"},
		{models.TheHague, "/DH/slots/?productKey=DOC&persons=1", "DOC"},
		{models.Rotterdam, "/RO/slots/?productKey=DOC&persons=1", "DOC"},
		{models.Zwolle, "/ZW/slots/?productKey=DOC&persons=1", "DOC"},
		{models.DenBosch, "/DB/slots/?productKey=DOC&persons=1", "DOC"},
	}
	assertURLs(t, got, want)
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

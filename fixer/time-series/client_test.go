package time_series

import (
	"os"
	"testing"

	"github.com/sample/go-webservices/fixer"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	apiUrl := "http://data.fixer.io/api"
	apiKey := ""

	fixer.Setup(apiUrl, apiKey)
	fixer.SetDebug(true)
}

func TestLatest(t *testing.T) {
	defer gock.Off()

	gock.New("http://data.fixer.io/api").
		Get("/timeseries").
		Reply(200).
		File("./responses/time_series.json")


	startDate := "2012-05-01"
	endDate := "2012-05-03"
	base := ""
	toCurrencies := []string{"USD", "AUD", "CAD"}

	resp, err := Get(startDate, endDate, base, toCurrencies)

	assert.Nil(t, err)
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, true, resp.Success, "success field should be true")
	assert.NotNil(t, resp.Rates, "symbols field should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

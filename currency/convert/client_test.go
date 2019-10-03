package convert

import (
	"os"
	"testing"

	"github.com/sample/go-webservices/currency"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	apiUrl := ""
	apiKey := ""

	currency.Setup(apiUrl, apiKey)
	currency.SetDebug(true)
}

func TestConvert(t *testing.T) {
	defer gock.Off()

	gock.New("").
		Get("/currency/convert").
		Reply(200).
		File("./responses/convert.json")

	from := "GBP"
	to := "JPY"
	amount := float64(25)

	resp, err := Convert(from, to, amount)

	assert.Nil(t, err)
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

package rates

import (
	"fmt"
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

func TestListRates(t *testing.T) {
	defer gock.Off()

	skip := 0
	limit := 10
	orderBy := ""

	id := "EUR"

	gock.New("").
		Get("/currencies/"+id+"/rates").
		MatchParam("skip", fmt.Sprintf("%v", skip)).
		MatchParam("limit", fmt.Sprintf("%v", limit)).
		MatchParam("order_by", fmt.Sprintf("%s", orderBy)).
		Reply(200).
		File("./responses/rates.json")

	params := &currency.PaginationParams{}
	params.Skip = skip
	params.Limit = limit

	resp, err := ListRates(id, params)

	assert.Nil(t, err)
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, limit, resp.Limit, "limit should be equal")
	assert.NotNil(t, resp.Rates, "symbols field should not be nil")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

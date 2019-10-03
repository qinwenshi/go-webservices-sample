package config

import (
	"os"
	"testing"
	"time"

	"github.com/sample/go-webservices/bemyguest2"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"net/url"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	apiUrl := ""
	apiKey := ""

	bemyguest2.Setup(apiUrl, apiKey)
	bemyguest2.SetDebug(true)
}

func TestGetConfig(t *testing.T) {
	defer gock.Off()

	gock.New("").
		Get("/v2/config").
		Reply(200).
		File("./responses/config.json")

	resp, err := Get()
	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")
	assert.Equal(t, len(resp.Data.CurrencyData.Data), 1, "should only have one currency")
	assert.Equal(t, resp.Data.CurrencyData.Data[0].Code, "SGD", "should be SGD")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetConfigWithUnauthorized(t *testing.T) {
	defer gock.Off()

	gock.New("").
		Get("/v2/config").
		Reply(401).
		File("./responses/error_unauthorized.json")

	resp, err := Get()
	assert.NotNil(t, err, "should be error")
	assert.Nil(t, resp, "response should be nil")
	assert.IsType(t, bemyguest2.ErrorResp{}, err, "should receive bemyguest2 error response")
	content := err.(bemyguest2.ErrorResp).Content
	assert.Equal(t, 401, content.HTTPCode)
	assert.Equal(t, "Unauthorized", content.Message)
	assert.Equal(t, "GEN-UNAUTHORIZED", content.Code)
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetConfigWithTimeout(t *testing.T) {
	defer gock.Off()

	ttl := time.Millisecond * 10
	bemyguest2.SetClientTimeout(ttl)
	defer bemyguest2.SetClientTimeout(bemyguest2.DefaultTTL)

	// gock is not able to trigger timeout
	// we try to send the real http request here
	gock.New("").
		Get("/v2/config").
		EnableNetworking()

	resp, err := Get()
	assert.NotNil(t, err, "should have timeout error")
	assert.Nil(t, resp, "response should be nil")
	assert.IsType(t, &url.Error{}, err, "should be timeout error")
	assert.Equal(t, true, err.(*url.Error).Timeout(), "should be timeout")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

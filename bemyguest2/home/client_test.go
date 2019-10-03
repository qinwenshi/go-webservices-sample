package home

import (
	"os"
	"testing"
	"time"
	"net/url"

	"github.com/sample/go-webservices/bemyguest2"

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

	bemyguest2.Setup(apiUrl, apiKey)
	bemyguest2.SetDebug(true)
}

func TestGetHome(t *testing.T) {
	defer gock.Off()

	gock.New("").
		Get("/v2/").
		Reply(200).
		File("./responses/home.json")

	resp, err := Get()
	assert.Nil(t, err, "should not have error")
	assert.NotNil(t, resp, "response should not be nil")

	assert.Equal(t, "BeMyGuest Distribution API v2", resp.Name)
	assert.Equal(t, "/v2", resp.BaseURL)
	assert.Equal(t, "UTF-8", resp.Encoding)
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

func TestGetHomeWithUnauthorized(t *testing.T) {
	defer gock.Off()

	gock.New("").
		Get("/v2/").
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

func TestGetHomeWithTimeout(t *testing.T) {
	defer gock.Off()

	ttl := time.Millisecond * 10
	bemyguest2.SetClientTimeout(ttl)
	defer bemyguest2.SetClientTimeout(bemyguest2.DefaultTTL)

	// gock is not able to trigger timeout
	// we try to send the real http request here
	gock.New("").
		Get("/v2/").
		EnableNetworking()

	resp, err := Get()
	assert.NotNil(t, err, "should have timeout error")
	assert.Nil(t, resp, "response should be nil")
	assert.IsType(t, &url.Error{}, err, "should be timeout error")
	assert.Equal(t, true, err.(*url.Error).Timeout(), "should be timeout")
	assert.Equal(t, true, gock.IsDone(), "should not have pending mocks")
}

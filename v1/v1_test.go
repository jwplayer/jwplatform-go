package v1

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_BuildParams(t *testing.T) {
	client := NewClient("API_KEY", "API_SECRET")

	// set URL params
	params := url.Values{}
	params.Set("video_key", "VIDEO_KEY")

	rawQuery := client.buildParams(params).Encode()

	m, err := url.ParseQuery(rawQuery)

	assert.Nil(t, err)
	assert.Equal(t, []string{"json"}, m["api_format"])
	assert.Equal(t, []string{"API_KEY"}, m["api_key"])
	assert.Equal(t, []string{"VIDEO_KEY"}, m["video_key"])
	assert.Contains(t, m, "api_nonce")
	assert.Contains(t, m, "api_signature")
	assert.Contains(t, m, "api_timestamp")
}

func TestClient_MakeRequest(t *testing.T) {
	defer gock.Off() // flush pending mocks after test execution

	gock.New("https://api.jwplatform.com").
		Get("/v1/videos/show").
		MatchParam("video_key", "VIDEO_KEY").
		Reply(200).
		JSON(map[string]string{"status": "ok"})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := NewClient("API_KEY", "API_SECRET")

	// set URL params
	params := url.Values{}
	params.Set("video_key", "VIDEO_KEY") // some video key, e.g. gIRtMhYM

	// declare an empty interface
	var result map[string]interface{}

	client.MakeRequest(ctx, http.MethodGet, "/videos/show/", params, &result)

	assert.Equal(t, "ok", result["status"])
}

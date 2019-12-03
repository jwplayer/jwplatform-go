package jwplatform

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_BuildParams(t *testing.T) {
	client := NewClient("API_KEY", "API_SECRET")

	// set URL params
	params := url.Values{}
	params.Set("video_key", "VIDEO_KEY")

	rawQuery := client.buildParams(params).Encode()

	m, err := url.ParseQuery(rawQuery)

	assert.Nil(t, err)
	assert.Equal(t, m["api_format"], []string{"json"})
	assert.Equal(t, m["api_key"], []string{"API_KEY"})
	assert.Equal(t, m["video_key"], []string{"VIDEO_KEY"})
	assert.Contains(t, m, "api_nonce")
	assert.Contains(t, m, "api_signature")
	assert.Contains(t, m, "api_timestamp")
}

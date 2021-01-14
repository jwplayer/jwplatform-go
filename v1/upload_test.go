package v1

import (
	"context"
	"net/url"
	"testing"

	filet "github.com/Flaque/filet"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_Upload(t *testing.T) {
	defer gock.Off()       // flush pending mocks after test execution
	defer filet.CleanUp(t) // clean up temporary files

	gock.New("https://api.jwplatform.com").
		Post("/v1/videos/create").
		Reply(200).
		JSON([]byte(`{"status": "ok", "media": {"type": "video", "key": "8TjWBMMX"}, "link": {"path": "/v1/videos/upload", "query": {"token": "86180b3b11f750a5598a16e5a4e852416ba9f09f78d", "key": "8TjWBMMX"}, "protocol": "http", "address": "upload.jwplatform.com"}, "rate_limit": {"reset": 1575813660, "limit": 60, "remaining": 59}}`))

	// create a temporary file with no parent dir
	file := filet.TmpFile(t, "", "")

	gock.New("https://upload.jwplatform.com").
		Post("/v1/videos/upload").
		File(file.Name()).
		Reply(200).
		JSON(map[string]string{"status": "ok"})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// set URL params
	params := url.Values{}
	params.Set("title", "Your video title")
	params.Set("description", "Your video description")

	client := NewClient("API_KEY", "API_SECRET")

	// declare an empty interface
	var result map[string]interface{}

	// upload video usind direct upload method
	err := client.Upload(ctx, file.Name(), params, &result)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "ok", result["status"])
}

package jwplatform

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestListUploadParts(t *testing.T) {
	defer gock.Off()

	uploadID := "abcdefgh"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/uploads/%s", uploadID)

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200)

	testClient := New(mockAuthToken)
	err := testClient.Uploads.ListUploadParts(uploadID)
	assert.Equal(t, nil, err)
}

func TestCompleteUpload(t *testing.T) {
	defer gock.Off()

	uploadID := "abcdefgh"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/uploads/%s/complete", uploadID)

	gock.New("https://api.jwplayer.com").
		Put(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(202)

	testClient := New(mockAuthToken)
	err := testClient.Uploads.CompleteUpload(uploadID)
	assert.Equal(t, nil, err)
}

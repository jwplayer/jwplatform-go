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
	mockPartsResponse := map[string]interface{}{
		"page_length": 10,
		"page":        1,
		"parts":       []map[string]string{{"etag": "test-etag"}, {"upload_link": "test-link"}},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockPartsResponse)

	testClient := New(mockAuthToken)
	uploadParts, err := testClient.Uploads.ListUploadParts(uploadID)
	assert.Equal(t, nil, err)

	assert.Equal(t, "test-etag", uploadParts.Parts[0].ETag)
	assert.Equal(t, "", uploadParts.Parts[0].UploadLink)
	assert.Equal(t, "", uploadParts.Parts[1].ETag)
	assert.Equal(t, "test-link", uploadParts.Parts[1].UploadLink)
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

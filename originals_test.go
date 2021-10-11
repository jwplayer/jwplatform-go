package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetOriginal(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	originalID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/originals/%s", siteID, mediaID, originalID)
	mockProtectionRuleResp := map[string]string{"id": originalID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockProtectionRuleResp)

	testClient := New(mockAuthToken)
	original, err := testClient.Originals.Get(siteID, mediaID, originalID)
	assert.Equal(t, originalID, original.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteOriginal(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	originalID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/originals/%s", siteID, mediaID, originalID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.Originals.Delete(siteID, mediaID, originalID)
	assert.Equal(t, nil, err)
}

func TestCreateOriginal(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	originalID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/originals", siteID, mediaID)
	mockProtectionRuleResp := map[string]string{"id": originalID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockProtectionRuleResp)

	testClient := New(mockAuthToken)
	OriginalCreateRequest := &OriginalCreateRequest{
		Metadata: OriginalMetadata{},
		Upload:   OriginalUploadRequest{},
	}
	original, err := testClient.Originals.Create(siteID, mediaID, OriginalCreateRequest)
	assert.Equal(t, originalID, original.ID)
	assert.Equal(t, nil, err)
}

func TestListOriginals(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	originalID := "aamkdcaz"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/originals", siteID, mediaID)
	mockPlayerBiddingResponse := map[string]interface{}{
		"page_length": pageLength,
		"page":        page,
		"originals":   []map[string]string{{"id": originalID}},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		MatchParam("page", strconv.Itoa(page)).
		MatchParam("page_length", strconv.Itoa(pageLength)).
		Reply(200).
		JSON(mockPlayerBiddingResponse)

	testClient := New(mockAuthToken)
	params := &QueryParams{PageLength: pageLength, Page: page}
	originalsResponse, err := testClient.Originals.List(siteID, mediaID, params)
	assert.Equal(t, page, originalsResponse.Page)
	assert.Equal(t, pageLength, originalsResponse.PageLength)
	assert.Equal(t, originalID, originalsResponse.Originals[0].ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalOriginal(t *testing.T) {
	originalData := map[string]interface{}{
		"id":               "abZqokMz",
		"type":             "original",
		"error_message":    nil,
		"size":             124,
		"md5":              "somemd5",
		"status":           "ready",
		"container_format": "format",
		"includes": map[string]interface{}{
			"video_track": map[string]interface{}{
				"abqds9lq": map[string]int{
					"height": 120,
					"width":  240,
				},
			},
		},
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"metadata": map[string]interface{}{
			"language":      "English",
			"language_code": "en",
			"name":          "Original",
			"type":          "primary",
		},
	}

	bytes, err := json.Marshal(&originalData)
	assert.NoError(t, err)

	var original OriginalResource
	err = json.Unmarshal(bytes, &original)
	assert.NoError(t, err)

	assert.Equal(t, "abZqokMz", original.ID)
	assert.Equal(t, "original", original.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", original.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", original.LastModified)

	assert.Equal(t, "", original.ErrorMessage)
	assert.Equal(t, 124, original.Size)
	assert.Equal(t, "ready", original.Status)
	assert.Equal(t, "somemd5", original.MD5)
	assert.Equal(t, "format", original.ContainerFormat)

	assert.Equal(t, "English", original.Metadata.Language)
	assert.Equal(t, "en", original.Metadata.LanguageCode)
	assert.Equal(t, "Original", original.Metadata.Name)
	assert.Equal(t, "primary", original.Metadata.Type)
}

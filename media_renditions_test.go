package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetMediaRendition(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	renditionID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/media_renditions/%s", siteID, mediaID, renditionID)
	mockProtectionRuleResp := map[string]string{"id": renditionID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockProtectionRuleResp)

	testClient := New(mockAuthToken)
	mediaRendition, err := testClient.MediaRenditions.Get(siteID, mediaID, renditionID)
	assert.Equal(t, renditionID, mediaRendition.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteMediaRendition(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	renditionID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/media_renditions/%s", siteID, mediaID, renditionID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.MediaRenditions.Delete(siteID, mediaID, renditionID)
	assert.Equal(t, nil, err)
}

func TestCreateMediaRendition(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	renditionID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/media_renditions", siteID, mediaID)
	mockProtectionRuleResp := map[string]string{"id": renditionID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockProtectionRuleResp)

	testClient := New(mockAuthToken)
	templateID := "sometemplate"
	mediaRendition, err := testClient.MediaRenditions.Create(siteID, mediaID, templateID)
	assert.Equal(t, renditionID, mediaRendition.ID)
	assert.Equal(t, nil, err)
}

func TestListMediaRenditions(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	renditionID := "aamkdcaz"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/media_renditions", siteID, mediaID)
	mockPlayerBiddingResponse := map[string]interface{}{
		"page_length":      pageLength,
		"page":             page,
		"media_renditions": []map[string]string{{"id": renditionID}},
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
	mediaRenditionsResponse, err := testClient.MediaRenditions.List(siteID, mediaID, params)
	assert.Equal(t, page, mediaRenditionsResponse.Page)
	assert.Equal(t, pageLength, mediaRenditionsResponse.PageLength)
	assert.Equal(t, renditionID, mediaRenditionsResponse.MediaRenditions[0].ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalMediaRendition(t *testing.T) {
	mediaRenditionData := map[string]interface{}{
		"id":            "abZqokMz",
		"type":          "media_rendition",
		"delivery_url":  "https://cdn.jwplayer.com/videos/gMwHLgJj-88MXImLx.mp4",
		"error_message": nil,
		"filesize":      124,
		"height":        240,
		"width":         320,
		"status":        "ready",
		"media_type":    "video",
		"relationships": map[string]interface{}{
			"media_rendition_template": map[string]string{
				"id": "abqds9lq",
			},
		},
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"metadata":      map[string]interface{}{},
	}

	bytes, err := json.Marshal(&mediaRenditionData)
	assert.NoError(t, err)

	var mediaRendition MediaRenditionResource
	err = json.Unmarshal(bytes, &mediaRendition)
	assert.NoError(t, err)

	assert.Equal(t, "abZqokMz", mediaRendition.ID)
	assert.Equal(t, "media_rendition", mediaRendition.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", mediaRendition.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", mediaRendition.LastModified)

	assert.Equal(t, "https://cdn.jwplayer.com/videos/gMwHLgJj-88MXImLx.mp4", mediaRendition.DeliveryURL)
	assert.Equal(t, "", mediaRendition.ErrorMessage)
	assert.Equal(t, 124, mediaRendition.FileSize)
	assert.Equal(t, 240, mediaRendition.Height)
	assert.Equal(t, 320, mediaRendition.Width)
	assert.Equal(t, "ready", mediaRendition.Status)
	assert.Equal(t, "video", mediaRendition.MediaType)
	// assert.Equal(t, "abqds9lq", mediaRendition.Relationships["media_rendition_template"]["id"])
}

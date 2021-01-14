package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetMedia(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s", siteID, mediaID)
	mockMediaResponse := map[string]string{"id": mediaID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockMediaResponse)

	testClient := New(mockAuthToken)
	media, err := testClient.Media.Get(siteID, mediaID)
	assert.Equal(t, mediaID, media.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteMedia(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s", siteID, mediaID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.Media.Delete(siteID, mediaID)
	assert.Equal(t, nil, err)
}

func TestCreateMedia(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media", siteID)
	mockMediaResponse := map[string]string{"id": mediaID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockMediaResponse)

	testClient := New(mockAuthToken)
	newMedia := &MediaMetadata{Title: "Another test video"}
	media, err := testClient.Media.Create(siteID, newMedia)
	assert.Equal(t, mediaID, media.ID)
	assert.Equal(t, nil, err)
}

func TestUpdateMedia(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s", siteID, mediaID)
	mockMediaResponse := map[string]string{"id": mediaID}

	gock.New("https://api.jwplayer.com").
		Patch(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockMediaResponse)

	testClient := New(mockAuthToken)
	updateMetadata := &MediaMetadata{Title: "Updating another test media"}
	media, err := testClient.Media.Update(siteID, mediaID, updateMetadata)
	assert.Equal(t, mediaID, media.ID)
	assert.Equal(t, nil, err)
}

func TestListMedia(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "mnbvcxkj"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/media", siteID)
	mockMediaResponse := map[string]interface{}{
		"page_length": pageLength,
		"page":        page,
		"media":       []map[string]string{{"id": mediaID}},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		MatchParam("page", strconv.Itoa(page)).
		MatchParam("page_length", strconv.Itoa(pageLength)).
		Reply(200).
		JSON(mockMediaResponse)

	testClient := New(mockAuthToken)
	params := &QueryParams{PageLength: pageLength, Page: page}
	mediaMediaResponse, err := testClient.Media.List(siteID, params)
	assert.Equal(t, page, mediaMediaResponse.Page)
	assert.Equal(t, pageLength, mediaMediaResponse.PageLength)
	assert.Equal(t, mediaID, mediaMediaResponse.Media[0].ID)
	assert.Equal(t, nil, err)
}

func TestReuploadMedia(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/reupload", siteID, mediaID)
	mockMediaResponse := map[string]string{"id": mediaID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockMediaResponse)

	testClient := New(mockAuthToken)
	upload := Upload{}
	media, _ := testClient.Media.Reupload(siteID, mediaID, &upload)
	assert.Equal(t, mediaID, media.ID)
}

func TestUnmarshalMedia(t *testing.T) {
	mediaData := map[string]interface{}{
		"id":             "abZqokMz",
		"type":           "media",
		"created":        "2019-09-25T15:29:11.042095+00:00",
		"last_modified":  "2019-09-25T15:29:11.042095+00:00",
		"error_message":  "",
		"status":         "ready",
		"external_id":    "abc",
		"mime_type":      "video/mp4",
		"duration":       360.00,
		"media_type":     "video",
		"hosting_type":   "hosted",
		"trim_in_point":  "00:00:33",
		"trim_out_point": "00:01:32.121",
		"source_url":     "http://media.com",
		"metadata": map[string]interface{}{
			"title":              "media",
			"description":        "Describes a media",
			"author":             "Steven Spielberg",
			"permalink":          "permalink.com",
			"category":           "Food",
			"publish_start_date": "2019-09-25T15:29:11.042095+00:00",
			"publish_end_date":   "2019-09-25T15:29:11.042095+00:00",
			"tags":               []string{"tag_a", "tag_b"},
			"custom_params": map[string]string{
				"key": "value",
			},
		},
	}

	bytes, err := json.Marshal(&mediaData)
	assert.NoError(t, err)

	var media MediaResource
	err = json.Unmarshal(bytes, &media)
	assert.NoError(t, err)

	assert.Equal(t, "abZqokMz", media.ID)
	assert.Equal(t, "media", media.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", media.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", media.LastModified)
	assert.Equal(t, "ready", media.Status)
	assert.Equal(t, "", media.ErrorMessage)
	assert.Equal(t, "abc", media.ExternalID)
	assert.Equal(t, "video/mp4", media.MimeType)
	assert.Equal(t, 360.00, media.Duration)
	assert.Equal(t, "video", media.MediaType)
	assert.Equal(t, "http://media.com", media.SourceURL)
	assert.Equal(t, "hosted", media.HostingType)
	assert.Equal(t, "00:00:33", media.TrimInPoint)
	assert.Equal(t, "00:01:32.121", media.TrimOutPoint)

	assert.Equal(t, "media", media.Metadata.Title)
	assert.Equal(t, "Describes a media", media.Metadata.Description)
	assert.Equal(t, "Steven Spielberg", media.Metadata.Author)
	assert.Equal(t, "permalink.com", media.Metadata.Permalink)
	assert.Equal(t, "Food", media.Metadata.Category)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", media.Metadata.PublishStartDate)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", media.Metadata.PublishEndDate)
	assert.Equal(t, []string{"tag_a", "tag_b"}, media.Metadata.Tags)
	assert.Equal(t, map[string]string{"key": "value"}, media.Metadata.CustomParams)
}

package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetThumnbnail(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	thumbnailID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/thumbnails/%s", siteID, thumbnailID)
	mockThumnbnailResponse := map[string]string{"id": thumbnailID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockThumnbnailResponse)

	testClient := New(mockAuthToken)
	thumbnail, err := testClient.Thumbnails.Get(siteID, thumbnailID)
	assert.Equal(t, thumbnailID, thumbnail.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteThumnbnail(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	thumbnailID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/thumbnails/%s", siteID, thumbnailID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.Thumbnails.Delete(siteID, thumbnailID)
	assert.Equal(t, nil, err)
}

func TestCreateThumnbnail(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	thumbnailID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/thumbnails", siteID)
	mockThumnbnailResponse := map[string]string{"id": thumbnailID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockThumnbnailResponse)

	testClient := New(mockAuthToken)
	thumbRelations := &ThumbnailRelationships{
		[]ThumbnailMedia{
			ThumbnailMedia{
				MediaID: "abcdefgh",
			},
		},
	}
	newThumnbnail := &CreateThumbnailRequest{
		Relationships: *thumbRelations,
		Upload: ThumbnailUpload{
			Method: "direct",
		},
	}
	thumbnail, err := testClient.Thumbnails.Create(siteID, newThumnbnail)
	assert.Equal(t, thumbnailID, thumbnail.ID)
	assert.Equal(t, nil, err)
}

func TestUpdateThumnbnail(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	thumbnailID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/thumbnails/%s", siteID, thumbnailID)
	mockThumnbnailResponse := map[string]string{"id": thumbnailID}

	gock.New("https://api.jwplayer.com").
		Patch(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockThumnbnailResponse)

	testClient := New(mockAuthToken)
	thumbRelations := &ThumbnailRelationships{
		[]ThumbnailMedia{
			ThumbnailMedia{
				MediaID:  "abcdefgh",
				IsPoster: true,
			},
		},
	}
	updateMetadata := &UpdateThumbnailRequest{Relationships: *thumbRelations}
	thumbnail, err := testClient.Thumbnails.Update(siteID, thumbnailID, updateMetadata)
	assert.Equal(t, thumbnailID, thumbnail.ID)
	assert.Equal(t, nil, err)
}

func TestListThumbnails(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	thumbnailID := "mnbvcxkj"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/thumbnails", siteID)
	mockThumnbnailResponse := map[string]interface{}{
		"page_length": pageLength,
		"page":        page,
		"thumbnail":   []map[string]string{{"id": thumbnailID}},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		MatchParam("page", strconv.Itoa(page)).
		MatchParam("page_length", strconv.Itoa(pageLength)).
		Reply(200).
		JSON(mockThumnbnailResponse)

	testClient := New(mockAuthToken)
	params := &QueryParams{PageLength: pageLength, Page: page}
	thumbnailThumnbnailResponse, err := testClient.Thumbnails.List(siteID, params)
	assert.Equal(t, page, thumbnailThumnbnailResponse.Page)
	assert.Equal(t, pageLength, thumbnailThumnbnailResponse.PageLength)
	assert.Equal(t, thumbnailID, thumbnailThumnbnailResponse.Thumbnails[0].ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalThumnbnail(t *testing.T) {
	thumbnailData := map[string]interface{}{
		"id":             "abZqokMz",
		"type":           "thumbnail",
		"created":        "2019-09-25T15:29:11.042095+00:00",
		"last_modified":  "2019-09-25T15:29:11.042095+00:00",
		"delivery_url":   "thumbnaildelivery.com",
		"thumbnail_type": "static",
		"source_type":    "custom_upload",
		"metadata":       map[string]interface{}{},
		"relationships": map[string]interface{}{
			"media": []map[string]interface{}{
				map[string]interface{}{
					"id":        "media_id",
					"is_poster": true,
				},
			},
		},
	}

	bytes, err := json.Marshal(&thumbnailData)
	assert.NoError(t, err)

	var thumbnail ThumbnailResource
	err = json.Unmarshal(bytes, &thumbnail)
	assert.NoError(t, err)

	assert.Equal(t, "abZqokMz", thumbnail.ID)
	assert.Equal(t, "thumbnail", thumbnail.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", thumbnail.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", thumbnail.LastModified)
	assert.Equal(t, "thumbnaildelivery.com", thumbnail.DeliveryURL)
	assert.Equal(t, "static", thumbnail.ThumbnailType)
	assert.Equal(t, "custom_upload", thumbnail.SourceType)
	assert.Equal(t, "media_id", thumbnail.Relationships.Media[0].MediaID)
	assert.Equal(t, true, thumbnail.Relationships.Media[0].IsPoster)
}

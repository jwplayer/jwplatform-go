package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetTextTrack(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	trackID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks/%s", siteID, mediaID, trackID)
	mockTextTrackResp := map[string]string{"id": trackID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockTextTrackResp)

	testClient := New(mockAuthToken)
	textTrack, err := testClient.TextTracks.Get(siteID, mediaID, trackID)
	assert.Equal(t, trackID, textTrack.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteTextTrack(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	trackID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks/%s", siteID, mediaID, trackID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.TextTracks.Delete(siteID, mediaID, trackID)
	assert.Equal(t, nil, err)
}

func TestCreateTextTrack(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	trackID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks", siteID, mediaID)
	mockTextTrackResp := map[string]string{"id": trackID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockTextTrackResp)

	testClient := New(mockAuthToken)
	trackCreateRequest := &TextTrackCreateRequest{}
	textTrack, err := testClient.TextTracks.Create(siteID, mediaID, trackCreateRequest)
	assert.Equal(t, trackID, textTrack.ID)
	assert.Equal(t, nil, err)
}

func TestListTextTracks(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	trackID := "aamkdcaz"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks", siteID, mediaID)
	mockPlayerBiddingResponse := map[string]interface{}{
		"page_length": pageLength,
		"page":        page,
		"text_tracks": []map[string]string{{"id": trackID}},
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
	textTracksResponse, err := testClient.TextTracks.List(siteID, mediaID, params)
	assert.Equal(t, page, textTracksResponse.Page)
	assert.Equal(t, pageLength, textTracksResponse.PageLength)
	assert.Equal(t, trackID, textTracksResponse.TextTracks[0].ID)
	assert.Equal(t, nil, err)
}

func TestPublishTextTrack(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	trackID := "19zsquao"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks/%s/publish", siteID, mediaID, trackID)

	gock.New("https://api.jwplayer.com").
		Put(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(202)

	testClient := New(mockAuthToken)
	err := testClient.TextTracks.Publish(siteID, mediaID, trackID)
	assert.Equal(t, nil, err)
}

func TestUnpublishTextTrack(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "iqhg8ysz"
	trackID := "19zsquao"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks/%s/unpublish", siteID, mediaID, trackID)

	gock.New("https://api.jwplayer.com").
		Put(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(202)

	testClient := New(mockAuthToken)
	err := testClient.TextTracks.Unpublish(siteID, mediaID, trackID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalTextTrack(t *testing.T) {
	textTrackData := map[string]interface{}{
		"id":            "abZqokMz",
		"type":          "text_track",
		"track_kind":    "captions",
		"delivery_url":  "some delivery url",
		"error_message": "error",
		"status":        "draft",
		"relationships": map[string]interface{}{},
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"metadata": map[string]interface{}{
			"label":    "my track",
			"srclang":  "es",
			"position": 2,
		},
	}

	bytes, err := json.Marshal(&textTrackData)
	assert.NoError(t, err)

	var textTrack TextTrackResource
	err = json.Unmarshal(bytes, &textTrack)
	assert.NoError(t, err)

	assert.Equal(t, "abZqokMz", textTrack.ID)
	assert.Equal(t, "text_track", textTrack.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", textTrack.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", textTrack.LastModified)

	assert.Equal(t, "some delivery url", textTrack.DeliveryURL)
	assert.Equal(t, "error", textTrack.ErrorMessage)
	assert.Equal(t, "captions", textTrack.TrackKind)
	assert.Equal(t, "draft", textTrack.Status)

	assert.Equal(t, "my track", textTrack.Metadata.Label)
	assert.Equal(t, "es", textTrack.Metadata.Srclang)
	assert.Equal(t, 2, textTrack.Metadata.Position)
}

func TestUnmarshalTextTrackCreate(t *testing.T) {
	textTrackData := map[string]interface{}{
		"id":            "abZqokMz",
		"type":          "text_track",
		"track_kind":    "captions",
		"delivery_url":  "some delivery url",
		"error_message": nil,
		"status":        "created",
		"upload_link":   "an upload link",
		"relationships": map[string]interface{}{},
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"metadata": map[string]interface{}{
			"label":    "my track",
			"srclang":  "es",
			"position": 2,
		},
	}

	bytes, err := json.Marshal(&textTrackData)
	assert.NoError(t, err)

	var textTrackCreateResponse TextTrackCreateResponse
	err = json.Unmarshal(bytes, &textTrackCreateResponse)
	assert.NoError(t, err)

	assert.Equal(t, "an upload link", textTrackCreateResponse.UploadLink)
}

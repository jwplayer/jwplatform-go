package jwplatform

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetEvents(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	channelID := "mnbvcxkj"
	eventID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/channels/%s/events/%s", siteID, channelID, eventID)
	mockEventResponse := map[string]string{"id": eventID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockEventResponse)

	testClient := New(mockAuthToken)
	event, err := testClient.Channels.Events.Get(siteID, channelID, eventID)
	assert.Equal(t, eventID, event.ID)
	assert.Equal(t, nil, err)
}

func TestListEvents(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	channelID := "mnbvcxkj"
	eventID := "aqpozjv2a"
	page := 2
	pageLength := 4
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/channels/%s/events", siteID, channelID)
	mockEventsResponse := map[string]interface{}{
		"page_length": pageLength,
		"page":        page,
		"events":      []map[string]string{{"id": eventID}},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockEventsResponse)

	testClient := New(mockAuthToken)
	params := &QueryParams{PageLength: pageLength, Page: page}
	eventsListResponse, err := testClient.Channels.Events.List(siteID, channelID, params)
	assert.Equal(t, page, eventsListResponse.Page)
	assert.Equal(t, pageLength, eventsListResponse.PageLength)
	assert.Equal(t, eventID, eventsListResponse.Events[0].ID)
	assert.Equal(t, nil, err)
}

func TestRequestMaster(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	channelID := "mnbvcxkj"
	eventID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/channels/%s/events/%s/request_master", siteID, channelID, eventID)

	gock.New("https://api.jwplayer.com").
		Put(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.Channels.Events.RequestMaster(siteID, channelID, eventID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalEvent(t *testing.T) {
	eventData := map[string]interface{}{
		"id":            "tNvzjo7S",
		"type":          "event",
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"media_id":      "a_media_id",
		"status":        "idle",
		"master_access": map[string]string{
			"status":     "available",
			"expiration": "2022-11-11T07:50:00+00:00",
		},
	}

	bytes, err := json.Marshal(&eventData)
	assert.NoError(t, err)

	var event EventResource
	err = json.Unmarshal(bytes, &event)
	assert.NoError(t, err)

	assert.Equal(t, "tNvzjo7S", event.ID)
	assert.Equal(t, "event", event.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", event.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", event.LastModified)
	assert.Equal(t, "idle", event.Status)
	assert.Equal(t, "a_media_id", event.MediaID)
	assert.Equal(t, "available", event.MasterAccess.Status)
	assert.Equal(t, "2022-11-11T07:50:00+00:00", event.MasterAccess.Expiration)
}

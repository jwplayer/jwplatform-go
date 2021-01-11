package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetChannel(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	channelID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/channels/%s", siteID, channelID)
	mockChannelResponse := map[string]string{"id": channelID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockChannelResponse)

	testClient := New(mockAuthToken)
	channel, err := testClient.Channels.Get(siteID, channelID)
	assert.Equal(t, channelID, channel.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteChannel(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	channelID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/channels/%s", siteID, channelID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.Channels.Delete(siteID, channelID)
	assert.Equal(t, nil, err)
}

func TestCreateChannel(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	channelID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/channels", siteID)
	mockChannelResponse := map[string]string{"id": channelID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockChannelResponse)

	testClient := New(mockAuthToken)
	newChannel := &ChannelCreateMetadata{Title: "Another test video"}
	channel, err := testClient.Channels.Create(siteID, newChannel)
	assert.Equal(t, channelID, channel.ID)
	assert.Equal(t, nil, err)
}

func TestUpdateChannel(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	channelID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/channels/%s", siteID, channelID)
	mockChannelResponse := map[string]string{"id": channelID}

	gock.New("https://api.jwplayer.com").
		Patch(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockChannelResponse)

	testClient := New(mockAuthToken)
	updateMetadata := &ChannelMetadata{Title: "Updating another test media"}
	channel, err := testClient.Channels.Update(siteID, channelID, updateMetadata)
	assert.Equal(t, channelID, channel.ID)
	assert.Equal(t, nil, err)
}

func TestListChannel(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	channelID := "aamkdcaz"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/channels", siteID)
	mockChannelsResponse := map[string]interface{}{
		"page_length": pageLength,
		"page":        page,
		"channels":    []map[string]string{{"id": channelID}},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		MatchParam("page", strconv.Itoa(page)).
		MatchParam("page_length", strconv.Itoa(pageLength)).
		Reply(200).
		JSON(mockChannelsResponse)

	testClient := New(mockAuthToken)
	params := &QueryParams{PageLength: pageLength, Page: page}
	channelsResponse, err := testClient.Channels.List(siteID, params)
	assert.Equal(t, page, channelsResponse.Page)
	assert.Equal(t, pageLength, channelsResponse.PageLength)
	assert.Equal(t, channelID, channelsResponse.Channels[0].ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalChannel(t *testing.T) {
	channelData := map[string]interface{}{
		"id":               "apmzKzSf",
		"type":             "channel",
		"created":          "2019-09-25T15:29:11.042095+00:00",
		"last_modified":    "2019-09-25T15:29:11.042095+00:00",
		"status":           "idle",
		"reconnect_window": 60,
		"stream_key":       "a_stream_key",
		"metadata": map[string]interface{}{
			"title": "my channel",
			"dvr":   "on",
			"simulcast_targets": []interface{}{
				map[string]string{
					"stream_key": "b_stream_key",
					"stream_url": "http://streamurl.com",
					"title":      "target title",
				},
			},
			"tags": []string{"tag_a", "tag_b"},
			"custom_params": map[string]string{
				"key": "value",
			},
		},
		"latency": "default",
		"recent_events": []interface{}{
			map[string]string{
				"media_id": "a_media_id",
				"status":   "idle",
			},
		},
	}

	bytes, err := json.Marshal(&channelData)
	assert.NoError(t, err)

	var channelResource ChannelResource
	err = json.Unmarshal(bytes, &channelResource)
	assert.NoError(t, err)

	assert.Equal(t, "apmzKzSf", channelResource.ID)
	assert.Equal(t, "channel", channelResource.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", channelResource.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", channelResource.LastModified)

	assert.Equal(t, "default", channelResource.Latency)
	assert.Equal(t, "idle", channelResource.Status)
	assert.Equal(t, "a_stream_key", channelResource.StreamKey)
	assert.Equal(t, 60, channelResource.ReconnectWindow)

	assert.Equal(t, "a_media_id", channelResource.RecentEvents[0].MediaID)
	assert.Equal(t, "idle", channelResource.RecentEvents[0].Status)

	assert.Equal(t, "my channel", channelResource.Metadata.Title)
	assert.Equal(t, "on", channelResource.Metadata.Dvr)
	assert.Equal(t, []string{"tag_a", "tag_b"}, channelResource.Metadata.Tags)
	assert.Equal(t, map[string]string{"key": "value"}, channelResource.Metadata.CustomParams)

	assert.Equal(t, "b_stream_key", channelResource.Metadata.SimulcastTargets[0].StreamKey)
	assert.Equal(t, "http://streamurl.com", channelResource.Metadata.SimulcastTargets[0].StreamURL)
	assert.Equal(t, "target title", channelResource.Metadata.SimulcastTargets[0].Title)
}

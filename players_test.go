package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetPlayer(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	playerID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/players/%s", siteID, playerID)
	mockPlayerResponse := map[string]string{"id": playerID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockPlayerResponse)

	testClient := New(mockAuthToken)
	player, err := testClient.Players.Get(siteID, playerID)
	assert.Equal(t, playerID, player.ID)
	assert.Equal(t, nil, err)
}

func TestDeletePlayer(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	playerID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/players/%s", siteID, playerID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.Players.Delete(siteID, playerID)
	assert.Equal(t, nil, err)
}

func TestCreatePlayer(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	playerID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/players", siteID)
	mockPlayerResponse := map[string]string{"id": playerID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockPlayerResponse)

	testClient := New(mockAuthToken)
	newPlayer := &PlayerMetadata{Name: "Test Player"}
	player, err := testClient.Players.Create(siteID, newPlayer, nil)
	assert.Equal(t, playerID, player.ID)
	assert.Equal(t, nil, err)
}

func TestCreatePlayerWithRelationships(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	playerID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/players", siteID)
	mockPlayerResponse := map[string]string{"id": playerID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockPlayerResponse)

	testClient := New(mockAuthToken)
	newPlayer := &PlayerMetadata{Name: "Test Player"}
	relationships := &PlayerRelationships{AdScheduleID: "adscheduleid", RecommendationsPlaylistID: "recsplaylistID"}
	player, err := testClient.Players.Create(siteID, newPlayer, relationships)
	assert.Equal(t, playerID, player.ID)
	assert.Equal(t, nil, err)
}

func TestUpdatePlayer(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	playerID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/players/%s", siteID, playerID)
	mockPlayerResponse := map[string]string{"id": playerID}

	gock.New("https://api.jwplayer.com").
		Patch(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockPlayerResponse)

	testClient := New(mockAuthToken)
	updateMetadata := &PlayerMetadata{Name: "My Updated DRM Policy"}
	player, err := testClient.Players.Update(siteID, playerID, updateMetadata)
	assert.Equal(t, playerID, player.ID)
	assert.Equal(t, nil, err)
}

func TestListPlayer(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	playerID := "aamkdcaz"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/players", siteID)
	mockPlayersResponse := map[string]interface{}{
		"page_length": pageLength,
		"page":        page,
		"players":     []map[string]string{{"id": playerID}},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		MatchParam("page", strconv.Itoa(page)).
		MatchParam("page_length", strconv.Itoa(pageLength)).
		Reply(200).
		JSON(mockPlayersResponse)

	testClient := New(mockAuthToken)
	params := &QueryParams{PageLength: pageLength, Page: page}
	playersResponse, err := testClient.Players.List(siteID, params)
	assert.Equal(t, page, playersResponse.Page)
	assert.Equal(t, pageLength, playersResponse.PageLength)
	assert.Equal(t, playerID, playersResponse.Players[0].ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalPlayer(t *testing.T) {
	playerData := map[string]interface{}{
		"id":            "OiUUoa90",
		"type":          "drm_policy",
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"metadata": map[string]interface{}{
			"name":            "test player",
			"release_channel": "production",
			"setup_config": map[string]interface{}{
				"autostart": "true",
			},
			"custom_params": map[string]string{
				"custom": "param",
			},
		},
	}

	bytes, err := json.Marshal(&playerData)
	assert.NoError(t, err)

	var playerResource PlayerResource
	err = json.Unmarshal(bytes, &playerResource)
	assert.NoError(t, err)

	assert.Equal(t, "OiUUoa90", playerResource.ID)
	assert.Equal(t, "drm_policy", playerResource.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", playerResource.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", playerResource.LastModified)

	assert.Equal(t, "test player", playerResource.Metadata.Name)
	assert.Equal(t, "production", playerResource.Metadata.ReleaseChannel)
	assert.Equal(t, map[string]interface{}{"autostart": "true"}, playerResource.Metadata.SetupConfig)
	assert.Equal(t, map[string]string{"custom": "param"}, playerResource.Metadata.CustomParams)
}

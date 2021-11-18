package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetPlayerBidding(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	playerBiddingConfigID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/advertising/player_bidding_configs/%s", siteID, playerBiddingConfigID)
	mockPlayerBiddingResponse := map[string]string{"id": playerBiddingConfigID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockPlayerBiddingResponse)

	testClient := New(mockAuthToken)
	playerBiddingConfig, err := testClient.PlayerBidding.Get(siteID, playerBiddingConfigID)
	assert.Equal(t, playerBiddingConfigID, playerBiddingConfig.ID)
	assert.Equal(t, nil, err)
}

func TestDeletePlayerBidding(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	playerBiddingConfigID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/advertising/player_bidding_configs/%s", siteID, playerBiddingConfigID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.PlayerBidding.Delete(siteID, playerBiddingConfigID)
	assert.Equal(t, nil, err)
}

func TestCreatePlayerBidding(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	playerBiddingConfigID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/advertising/player_bidding_configs", siteID)
	mockPlayerBiddingResponse := map[string]string{"id": playerBiddingConfigID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockPlayerBiddingResponse)

	testClient := New(mockAuthToken)
	newPlayerBidding := &PlayerBiddingConfigurationMetadata{}
	playerBiddingConfig, err := testClient.PlayerBidding.Create(siteID, newPlayerBidding)
	assert.Equal(t, playerBiddingConfigID, playerBiddingConfig.ID)
	assert.Equal(t, nil, err)
}

func TestUpdatePlayerBidding(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	playerBiddingConfigID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/advertising/player_bidding_configs/%s", siteID, playerBiddingConfigID)
	mockPlayerBiddingResponse := map[string]string{"id": playerBiddingConfigID}

	gock.New("https://api.jwplayer.com").
		Patch(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockPlayerBiddingResponse)

	testClient := New(mockAuthToken)
	updateMetadata := &PlayerBiddingConfigurationMetadata{}
	playerBiddingConfig, err := testClient.PlayerBidding.Update(siteID, playerBiddingConfigID, updateMetadata)
	assert.Equal(t, playerBiddingConfigID, playerBiddingConfig.ID)
	assert.Equal(t, nil, err)
}

func TestListPlayerBidding(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	playerBiddingConfigID := "aamkdcaz"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/advertising/player_bidding_configs", siteID)
	mockPlayerBiddingResponse := map[string]interface{}{
		"page_length":            pageLength,
		"page":                   page,
		"player_bidding_configs": []map[string]string{{"id": playerBiddingConfigID}},
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
	playerBiddingConfigsResp, err := testClient.PlayerBidding.List(siteID, params)
	assert.Equal(t, page, playerBiddingConfigsResp.Page)
	assert.Equal(t, pageLength, playerBiddingConfigsResp.PageLength)
	assert.Equal(t, playerBiddingConfigID, playerBiddingConfigsResp.PlayerBiddingConfigs[0].ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalPlayerBiddingConfig(t *testing.T) {
	playerBiddingConfigData := map[string]interface{}{
		"id":            "abZqokMz",
		"type":          "player_bidding_config",
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"metadata": map[string]interface{}{
			"bids": map[string]interface{}{
				"settings": map[string]interface{}{
					"bidTimeout":             1,
					"floorPriceCents":        50,
					"mediationLayerAdServer": "dfp",
					"buckets": []interface{}{
						map[string]float64{
							"min":       5.00,
							"max":       5.50,
							"increment": 0.10,
						},
					},
					"consentManagement": map[string]interface{}{
						"gdpr": map[string]interface{}{
							"allowAuctionWithoutConsent": nil,
							"cmpApi":                     "iab",
							"defaultGdprScope":           false,
							"rules": []interface{}{
								map[string]interface{}{
									"enforcePurpose": true,
									"enforceVendor":  true,
									"purpose":        "storage",
								},
								map[string]interface{}{
									"enforcePurpose": true,
									"enforceVendor":  true,
									"purpose":        "basicAds",
									"vendorExceptions": []string{
										"MediaGrid",
										"SpotX",
									},
								},
								map[string]interface{}{
									"enforcePurpose": true,
									"enforceVendor":  true,
									"purpose":        "measurement",
								},
							},
							"timeout": 10000,
						},
						"usp": map[string]interface{}{
							"cmpApi":  "iab",
							"timeout": 10000,
						},
					},
				},
				"bidders": []interface{}{
					map[string]interface{}{
						"name":  "name",
						"id":    "id_a",
						"pubid": "pubid_b",
						"custom_params": map[string]string{
							"key": "value",
						},
					},
				},
			},
		},
	}

	bytes, err := json.Marshal(&playerBiddingConfigData)
	assert.NoError(t, err)

	var playerBiddingConfig PlayerBiddingConfigurationResource
	err = json.Unmarshal(bytes, &playerBiddingConfig)
	assert.NoError(t, err)

	assert.Equal(t, "abZqokMz", playerBiddingConfig.ID)
	assert.Equal(t, "player_bidding_config", playerBiddingConfig.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", playerBiddingConfig.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", playerBiddingConfig.LastModified)

	assert.Equal(t, 1, playerBiddingConfig.Metadata.Bids.Settings.BidTimeout)
	assert.Equal(t, 50, playerBiddingConfig.Metadata.Bids.Settings.FloorPriceCents)
	assert.Equal(t, "dfp", playerBiddingConfig.Metadata.Bids.Settings.MediationLayerAdServer)

	assert.Equal(t, 5.00, playerBiddingConfig.Metadata.Bids.Settings.Buckets[0].Min)
	assert.Equal(t, 5.50, playerBiddingConfig.Metadata.Bids.Settings.Buckets[0].Max)
	assert.Equal(t, 0.10, playerBiddingConfig.Metadata.Bids.Settings.Buckets[0].Increment)

	assert.Equal(t, "name", playerBiddingConfig.Metadata.Bids.Bidders[0].Name)
	assert.Equal(t, "id_a", playerBiddingConfig.Metadata.Bids.Bidders[0].ID)
	assert.Equal(t, "pubid_b", playerBiddingConfig.Metadata.Bids.Bidders[0].PubID)
	assert.Equal(t, map[string]string{"key": "value"}, playerBiddingConfig.Metadata.Bids.Bidders[0].CustomParams)
}

package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetMediaProtectionRule(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	protectionRuleID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media_protection_rules/%s", siteID, protectionRuleID)
	mockProtectionRuleResp := map[string]string{"id": protectionRuleID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockProtectionRuleResp)

	testClient := New(mockAuthToken)
	mediaProtectionRule, err := testClient.MediaProtectionRules.Get(siteID, protectionRuleID)
	assert.Equal(t, protectionRuleID, mediaProtectionRule.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteMediaProtectionRule(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	protectionRuleID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media_protection_rules/%s", siteID, protectionRuleID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.MediaProtectionRules.Delete(siteID, protectionRuleID)
	assert.Equal(t, nil, err)
}

func TestCreateMediaProtectionRule(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	protectionRuleID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media_protection_rules", siteID)
	mockProtectionRuleResp := map[string]string{"id": protectionRuleID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockProtectionRuleResp)

	testClient := New(mockAuthToken)
	newMediaProtectionRule := &MediaProtectionRuleMetadata{}
	mediaProtectionRule, err := testClient.MediaProtectionRules.Create(siteID, newMediaProtectionRule)
	assert.Equal(t, protectionRuleID, mediaProtectionRule.ID)
	assert.Equal(t, nil, err)
}

func TestUpdateMediaProtectionRule(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	protectionRuleID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/media_protection_rules/%s", siteID, protectionRuleID)
	mockPlayerBiddingResponse := map[string]string{"id": protectionRuleID}

	gock.New("https://api.jwplayer.com").
		Patch(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockPlayerBiddingResponse)

	testClient := New(mockAuthToken)
	updateMetadata := &MediaProtectionRuleMetadata{}
	mediaProtectionRule, err := testClient.MediaProtectionRules.Update(siteID, protectionRuleID, updateMetadata)
	assert.Equal(t, protectionRuleID, mediaProtectionRule.ID)
	assert.Equal(t, nil, err)
}

func TestListMediaProtectionRules(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	protectionRuleID := "aamkdcaz"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/media_protection_rules", siteID)
	mockPlayerBiddingResponse := map[string]interface{}{
		"page_length":            pageLength,
		"page":                   page,
		"media_protection_rules": []map[string]string{{"id": protectionRuleID}},
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
	mediaProtectionRulesResponse, err := testClient.MediaProtectionRules.List(siteID, params)
	assert.Equal(t, page, mediaProtectionRulesResponse.Page)
	assert.Equal(t, pageLength, mediaProtectionRulesResponse.PageLength)
	assert.Equal(t, protectionRuleID, mediaProtectionRulesResponse.MediaProtectionRules[0].ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalMediaProtectionRule(t *testing.T) {
	mediaProtectionRuleData := map[string]interface{}{
		"id":            "abZqokMz",
		"type":          "media_protection_rule",
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"metadata": map[string]interface{}{
			"name":      "test rule",
			"rule_type": "allow",
			"countries": []string{"US", "RU"},
		},
	}

	bytes, err := json.Marshal(&mediaProtectionRuleData)
	assert.NoError(t, err)

	var mediaProtectionRule MediaProtectionRuleResource
	err = json.Unmarshal(bytes, &mediaProtectionRule)
	assert.NoError(t, err)

	assert.Equal(t, "abZqokMz", mediaProtectionRule.ID)
	assert.Equal(t, "media_protection_rule", mediaProtectionRule.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", mediaProtectionRule.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", mediaProtectionRule.LastModified)

	assert.Equal(t, []string{"US", "RU"}, mediaProtectionRule.Metadata.Countries)
	assert.Equal(t, "allow", mediaProtectionRule.Metadata.RuleType)
	assert.Equal(t, "test rule", mediaProtectionRule.Metadata.Name)
}

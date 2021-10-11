package jwplatform

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetSiteProtectionRule(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	protectionRuleID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/site_protection_rule", siteID)
	mockProtectionRuleResp := map[string]string{"id": protectionRuleID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockProtectionRuleResp)

	testClient := New(mockAuthToken)
	siteProtectionRule, err := testClient.SiteProtectionRule.Get(siteID)
	assert.Equal(t, protectionRuleID, siteProtectionRule.ID)
	assert.Equal(t, nil, err)
}

func TestUpdateSiteProtectionRule(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	protectionRuleID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/site_protection_rule", siteID)
	mockPlayerBiddingResponse := map[string]string{"id": protectionRuleID}

	gock.New("https://api.jwplayer.com").
		Patch(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockPlayerBiddingResponse)

	testClient := New(mockAuthToken)
	updateMetadata := &SiteProtectionRuleMetadata{}
	siteProtectionRule, err := testClient.SiteProtectionRule.Update(siteID, updateMetadata)
	assert.Equal(t, protectionRuleID, siteProtectionRule.ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalSiteProtectionRule(t *testing.T) {
	siteProtectionRuleData := map[string]interface{}{
		"id":            "abZqokMz",
		"type":          "site_protection_rule",
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"metadata": map[string]interface{}{
			"rule_type": "allow",
			"countries": []string{"US", "RU"},
		},
	}

	bytes, err := json.Marshal(&siteProtectionRuleData)
	assert.NoError(t, err)

	var siteProtectionRule SiteProtectionRuleResource
	err = json.Unmarshal(bytes, &siteProtectionRule)
	assert.NoError(t, err)

	assert.Equal(t, "abZqokMz", siteProtectionRule.ID)
	assert.Equal(t, "site_protection_rule", siteProtectionRule.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", siteProtectionRule.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", siteProtectionRule.LastModified)

	assert.Equal(t, []string{"US", "RU"}, siteProtectionRule.Metadata.Countries)
	assert.Equal(t, "allow", siteProtectionRule.Metadata.RuleType)
}

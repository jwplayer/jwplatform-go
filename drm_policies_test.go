package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetDRMPolicy(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	drmPolicyID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/drm_policies/%s", siteID, drmPolicyID)
	mockDRMPolicyResponse := map[string]string{"id": drmPolicyID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockDRMPolicyResponse)

	testClient := New(mockAuthToken)
	drmPolicy, err := testClient.DRMPolicies.Get(siteID, drmPolicyID)
	assert.Equal(t, drmPolicyID, drmPolicy.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteDRMPolicy(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	drmPolicyID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/drm_policies/%s", siteID, drmPolicyID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.DRMPolicies.Delete(siteID, drmPolicyID)
	assert.Equal(t, nil, err)
}

func TestCreateDRMPolicy(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	drmPolicyID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/drm_policies", siteID)
	mockDRMPolicyResponse := map[string]string{"id": drmPolicyID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockDRMPolicyResponse)

	testClient := New(mockAuthToken)
	newDRMPolicy := &DRMPolicyMetadata{Name: "My DRM Policy"}
	drmPolicy, err := testClient.DRMPolicies.Create(siteID, newDRMPolicy)
	assert.Equal(t, drmPolicyID, drmPolicy.ID)
	assert.Equal(t, nil, err)
}

func TestUpdateDRMPolicy(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	drmPolicyID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/drm_policies/%s", siteID, drmPolicyID)
	mockDRMPolicyResponse := map[string]string{"id": drmPolicyID}

	gock.New("https://api.jwplayer.com").
		Patch(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockDRMPolicyResponse)

	testClient := New(mockAuthToken)
	updateMetadata := &DRMPolicyMetadata{Name: "My Updated DRM Policy"}
	drmPolicy, err := testClient.DRMPolicies.Update(siteID, drmPolicyID, updateMetadata)
	assert.Equal(t, drmPolicyID, drmPolicy.ID)
	assert.Equal(t, nil, err)
}

func TestListDRMPolicy(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	drmPolicyID := "aamkdcaz"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/drm_policies", siteID)
	mockDRMPolicysResponse := map[string]interface{}{
		"page_length":  pageLength,
		"page":         page,
		"drm_policies": []map[string]string{{"id": drmPolicyID}},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		MatchParam("page", strconv.Itoa(page)).
		MatchParam("page_length", strconv.Itoa(pageLength)).
		Reply(200).
		JSON(mockDRMPolicysResponse)

	testClient := New(mockAuthToken)
	params := &QueryParams{PageLength: pageLength, Page: page}
	drmPolicysResponse, err := testClient.DRMPolicies.List(siteID, params)
	assert.Equal(t, page, drmPolicysResponse.Page)
	assert.Equal(t, pageLength, drmPolicysResponse.PageLength)
	assert.Equal(t, drmPolicyID, drmPolicysResponse.DRMPolicies[0].ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshaDrmPolicy(t *testing.T) {
	drmPolicyData := map[string]interface{}{
		"id":            "OiUUoa90",
		"type":          "drm_policy",
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"metadata": map[string]interface{}{
			"name":                       "my drm",
			"max_width":                  1680,
			"widevine_security":          "sw_secure_crypto",
			"playready_security":         3000,
			"allow_offline_persisteence": false,
			"digital_output_protection":  "not_required",
			"license_duration":           12600,
			"playback_duration":          24600,
		},
	}

	bytes, err := json.Marshal(&drmPolicyData)
	assert.NoError(t, err)

	var drmPolicyResource DRMPolicyResource
	err = json.Unmarshal(bytes, &drmPolicyResource)
	assert.NoError(t, err)

	assert.Equal(t, "OiUUoa90", drmPolicyResource.ID)
	assert.Equal(t, "drm_policy", drmPolicyResource.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", drmPolicyResource.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", drmPolicyResource.LastModified)

	assert.Equal(t, "my drm", drmPolicyResource.Metadata.Name)
	assert.Equal(t, 1680, drmPolicyResource.Metadata.MaxWidth)
	assert.Equal(t, "sw_secure_crypto", drmPolicyResource.Metadata.WidevineSecurity)
	assert.Equal(t, 3000, drmPolicyResource.Metadata.PlayreadySecurity)
	assert.Equal(t, false, drmPolicyResource.Metadata.AllowOfflinePersistence)
	assert.Equal(t, "not_required", drmPolicyResource.Metadata.DigitalOutputProtection)
	assert.Equal(t, 12600, drmPolicyResource.Metadata.LicenseDuration)
	assert.Equal(t, 24600, drmPolicyResource.Metadata.PlaybackDuration)
}

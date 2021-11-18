package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetAdSchedule(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	scheduleID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/advertising/schedules/%s", siteID, scheduleID)
	mockSchduleResponse := map[string]string{"id": scheduleID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockSchduleResponse)

	testClient := New(mockAuthToken)
	schedule, err := testClient.AdSchedules.Get(siteID, scheduleID)
	assert.Equal(t, scheduleID, schedule.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteAdSchedule(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	scheduleID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/advertising/schedules/%s", siteID, scheduleID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.AdSchedules.Delete(siteID, scheduleID)
	assert.Equal(t, nil, err)
}

func TestCreateAdSchedule(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	scheduleID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/advertising/schedules", siteID)
	mockSchduleResponse := map[string]string{"id": scheduleID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockSchduleResponse)

	testClient := New(mockAuthToken)
	metadata := &AdScheduleMetadata{}
	schedule, err := testClient.AdSchedules.Create(siteID, metadata)
	assert.Equal(t, scheduleID, schedule.ID)
	assert.Equal(t, nil, err)
}

func TestUpdateAdSchedule(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	scheduleID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/advertising/schedules/%s", siteID, scheduleID)
	mockSchduleResponse := map[string]string{"id": scheduleID}

	gock.New("https://api.jwplayer.com").
		Patch(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockSchduleResponse)

	testClient := New(mockAuthToken)
	updateMetadata := &AdScheduleMetadata{}
	schedule, err := testClient.AdSchedules.Update(siteID, scheduleID, updateMetadata)
	assert.Equal(t, scheduleID, schedule.ID)
	assert.Equal(t, nil, err)
}

func TestListAdSchedules(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	scheduleID := "aamkdcaz"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/advertising/schedules", siteID)
	mockSchdulesResponse := map[string]interface{}{
		"page_length": pageLength,
		"page":        page,
		"schedules":   []map[string]string{{"id": scheduleID}},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		MatchParam("page", strconv.Itoa(page)).
		MatchParam("page_length", strconv.Itoa(pageLength)).
		Reply(200).
		JSON(mockSchdulesResponse)

	testClient := New(mockAuthToken)
	params := &QueryParams{PageLength: pageLength, Page: page}
	schedules, err := testClient.AdSchedules.List(siteID, params)
	assert.Equal(t, page, schedules.Page)
	assert.Equal(t, pageLength, schedules.PageLength)
	assert.Equal(t, scheduleID, schedules.Schedules[0].ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalAdSchdule(t *testing.T) {
	scheduleData := map[string]interface{}{
		"created":       "2021-11-17T16:54:39.487381+00:00",
		"id":            "bC2jMY5Z",
		"last_modified": "2021-11-17T16:54:39.487381+00:00",
		"metadata": map[string]interface{}{
			"breaks": []interface{}{
				map[string]interface{}{
					"offset":     "pre",
					"skipoffset": 3,
					"tags": []string{
						"http://online.com",
					},
					"type": "linear",
				},
				map[string]interface{}{
					"offset": "post",
					"tags": []string{
						"https://online.com",
					},
					"type": "linear",
				},
			},
			"client":  "vast",
			"is_vmap": false,
			"name":    "Ad Schedule",
			"rules": map[string]interface{}{
				"frequency":      1,
				"startOn":        1,
				"startOnSeek":    "pre",
				"timeBetweenAds": 0,
			},
			"version": "1.4",
		},
		"schema": "https://schema.jwplayer.com/types/adschedule.json",
		"type":   "adschedule",
	}

	bytes, err := json.Marshal(&scheduleData)
	assert.NoError(t, err)

	var schedule AdScheduleResource
	err = json.Unmarshal(bytes, &schedule)
	assert.NoError(t, err)

	assert.Equal(t, "bC2jMY5Z", schedule.ID)
	assert.Equal(t, "adschedule", schedule.Type)
	assert.Equal(t, "2021-11-17T16:54:39.487381+00:00", schedule.Created)
	assert.Equal(t, "2021-11-17T16:54:39.487381+00:00", schedule.LastModified)

	assert.Equal(t, false, schedule.Metadata.IsVMAP)
	assert.Equal(t, "Ad Schedule", schedule.Metadata.Name)
	assert.Equal(t, "vast", schedule.Metadata.Client)
	assert.Equal(t, "1.4", schedule.Metadata.Version)

	assert.Equal(t, "pre", schedule.Metadata.Breaks[0].Offset)
	assert.Equal(t, 3, schedule.Metadata.Breaks[0].SkipOffset)
	assert.Equal(t, "http://online.com", schedule.Metadata.Breaks[0].Tags[0])
	assert.Equal(t, "linear", schedule.Metadata.Breaks[0].Type)

	assert.Equal(t, 1, schedule.Metadata.Rules.Frequency)
	assert.Equal(t, 1, schedule.Metadata.Rules.StartOn)
	assert.Equal(t, "pre", schedule.Metadata.Rules.StartOnSeek)
	assert.Equal(t, 0, schedule.Metadata.Rules.TimeBetweenAds)
}

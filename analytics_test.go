package jwplatform

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestAnalyticsQuery(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/analytics/queries", siteID)
	mockResponse := map[string]bool{"include_metadata": true}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		MatchParam("source", "default").
		MatchParam("format", "json").
		Reply(200).
		JSON(mockResponse)

	testClient := New(mockAuthToken)
	analyticsParams := &AnalyticsQueryParameters{Source: "default", Format: "json"}
	analyticsResp, _ := testClient.Analytics.Query(siteID, analyticsParams)
	assert.Equal(t, true, analyticsResp.IncludeMetadata)
}

func TestUnmarshalAnalyticsResponse(t *testing.T) {
	analyticsData := map[string]interface{}{
		"dimensions":       []string{"dimension_a", "dimension_b"},
		"start_date":       "2019-09-25T15:29:11.042095+00:00",
		"end_date":         "2019-09-25T15:29:11.042095+00:00",
		"filter":           "a_filter",
		"include_metadata": true,
		"metrics": []interface{}{
			map[string]string{
				"field":     "field_a",
				"operation": "=",
			},
		},
		"sort": []interface{}{
			map[string]string{
				"field":     "field_b",
				"operation": "max",
				"order":     "ascending",
			},
		},
		"page":               1,
		"page_length":        10,
		"relative_timeframe": "7 Days",
	}

	bytes, err := json.Marshal(&analyticsData)
	assert.NoError(t, err)

	var analytics AnalyticsResponse
	err = json.Unmarshal(bytes, &analytics)
	assert.NoError(t, err)

	assert.Equal(t, []string{"dimension_a", "dimension_b"}, analytics.Dimensions)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", analytics.StartDate)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", analytics.EndDate)
	assert.Equal(t, "a_filter", analytics.Filter)
	assert.Equal(t, 1, analytics.Page)
	assert.Equal(t, 10, analytics.PageLength)
	assert.Equal(t, "7 Days", analytics.RelativeTimeframe)
	assert.Equal(t, true, analytics.IncludeMetadata)

	assert.Equal(t, "field_a", analytics.Metrics[0].Field)
	assert.Equal(t, "=", analytics.Metrics[0].Operation)

	assert.Equal(t, "field_b", analytics.Sort[0].Field)
	assert.Equal(t, "max", analytics.Sort[0].Operation)
	assert.Equal(t, "ascending", analytics.Sort[0].Order)

}

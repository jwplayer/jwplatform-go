package jwplatform

import (
	"fmt"
	"net/http"
	"strings"
)

// AnalyticsQueryParameters define the allowed parameters on the Query action.
type AnalyticsQueryParameters struct {
	Source string
	Format string
}

// AnalyticsResponse is the structure returned via the Query action.
type AnalyticsResponse struct {
	Dimensions        []string       `json:"dimensions"`
	StartDate         string         `json:"start_date"`
	EndDate           string         `json:"end_date"`
	Filter            string         `json:"filter"`
	IncludeMetadata   bool           `json:"include_metadata"`
	Metrics           []reportMetric `json:"metrics"`
	Sort              []reportSort   `json:"sort"`
	Page              int            `json:"page"`
	PageLength        int            `json:"page_length"`
	RelativeTimeframe string         `json:"relative_timeframe"`
}

type reportMetric struct {
	Field     string `json:"field"`
	Operation string `json:"operation"`
}

type reportSort struct {
	Field     string `json:"field"`
	Operation string `json:"operation"`
	Order     string `json:"order"`
}

// AnalyticsClient for interacting with V2 Analytics API.
type AnalyticsClient struct {
	v2Client *V2Client
}

// Query the Analytics API
func (c *AnalyticsClient) Query(siteID string, queryParams *AnalyticsQueryParameters) (*AnalyticsResponse, error) {
	analyticResponse := &AnalyticsResponse{}
	path := fmt.Sprintf("/v2/sites/%s/analytics/queries", siteID)
	if queryParams != nil {
		path += "?"
		var pathParams []string
		if queryParams.Source != "" {
			pathParams = append(pathParams, fmt.Sprintf("source=%s", queryParams.Source))
		}
		if queryParams.Format != "" {
			pathParams = append(pathParams, fmt.Sprintf("format=%s", queryParams.Format))
		}
		path += strings.Join(pathParams, "&")
	}
	err := c.v2Client.Request(http.MethodPost, path, analyticResponse, nil, nil)
	return analyticResponse, err
}

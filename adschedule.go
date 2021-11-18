package jwplatform

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type AdScheduleResource struct {
	V2ResourceResponse
	Metadata AdScheduleMetadata `json:"metadata"`
}

type AdScheduleResources struct {
	V2ResourcesResponse
	Schedules []AdScheduleResource `json:"schedules"`
}

type AdScheduleMetadata struct {
	Breaks  []AdBreaks   `json:"breaks"`
	Client  string       `json:"client"`
	IsVMAP  bool         `json:"is_vmap"`
	Name    string       `json:"name"`
	Rules   AdRules      `json:"rules"`
	Version string       `json:"version"`
	Bids    BidsMetadata `json:"bids"`
}

type AdBreaks struct {
	Offset     string   `json:"offset"`
	SkipOffset int      `json:"skipoffset,omitempty"`
	Tags       []string `json:"tags"`
	Type       string   `json:"type"`
}

type AdRules struct {
	Frequency      int    `json:"frequency"`
	StartOn        int    `json:"startOn"`
	StartOnSeek    string `json:"startOnSeek"`
	TimeBetweenAds int    `json:"timeBetweenAds"`
}

type AdScheduleWriteRequest struct {
	Metadata AdScheduleMetadata `json:"metadata"`
}

type AdScheduleClient struct {
	v2Client *V2Client
}

func (c *AdScheduleClient) Get(siteID, scheduleID string) (*AdScheduleResource, error) {
	schedule := &AdScheduleResource{}
	path := fmt.Sprintf("/v2/sites/%s/advertising/schedules/%s", siteID, scheduleID)
	err := c.v2Client.Request(http.MethodGet, path, schedule, nil, nil)
	return schedule, err
}

func (c *AdScheduleClient) Create(siteID string, metadata *AdScheduleMetadata) (*AdScheduleResource, error) {
	createRequestData := &AdScheduleWriteRequest{Metadata: *metadata}
	schedule := &AdScheduleResource{}
	path := fmt.Sprintf("/v2/sites/%s/advertising/schedules", siteID)
	err := c.v2Client.Request(http.MethodPost, path, schedule, createRequestData, nil)
	return schedule, err
}

func (c *AdScheduleClient) List(siteID string, queryParams *QueryParams) (*AdScheduleResources, error) {
	schedules := &AdScheduleResources{}
	path := fmt.Sprintf("/v2/sites/%s/advertising/schedules", siteID)
	urlValues, _ := query.Values(queryParams)
	err := c.v2Client.Request(http.MethodGet, path, schedules, nil, urlValues)
	return schedules, err
}

func (c *AdScheduleClient) Update(siteID, scheduleID string, metadata *AdScheduleMetadata) (*AdScheduleResource, error) {
	updateRequestData := &AdScheduleWriteRequest{Metadata: *metadata}
	schedule := &AdScheduleResource{}
	path := fmt.Sprintf("/v2/sites/%s/advertising/schedules/%s", siteID, scheduleID)
	err := c.v2Client.Request(http.MethodPatch, path, schedule, updateRequestData, nil)
	return schedule, err
}

func (c *AdScheduleClient) Delete(siteID, scheduleID string) error {
	path := fmt.Sprintf("/v2/sites/%s/advertising/schedules/%s", siteID, scheduleID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

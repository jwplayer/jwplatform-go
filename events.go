package jwplatform

import (
	"fmt"
	"net/http"
)

// EventResource is the resource that is returned for all Event resource requests,
// with the exception of the Create action, which extends this struct with upload-related data.
type EventResource struct {
	V2ResourceResponse
	MasterAccess masterAccess `json:"master_access"`
	MediaID      string       `json:"media_id"`
	Status       string       `json:"status"`
}

type masterAccess struct {
	Status     string `json:"status"`
	Expiration string `json:"expiration"`
}

// EventResourcesResponse is the response structure for Event list calls.
type EventResourcesResponse struct {
	V2ResourcesResponse
	Events []EventResource `json:"events"`
}

// EventsClient for interacting with V2 Events API.
type EventsClient struct {
	v2Client *V2Client
}

// Get a single Event resource by Channel and Event ID.
func (c *EventsClient) Get(siteID, channelID, eventID string) (*EventResource, error) {
	channel := &EventResource{}
	path := fmt.Sprintf("/v2/sites/%s/channels/%s/events/%s", siteID, channelID, eventID)
	err := c.v2Client.Request(http.MethodGet, path, channel, nil, nil)
	return channel, err
}

// List all Event resources associated with a given Site and Channel ID.
func (c *EventsClient) List(siteID, channelID string, queryParams *QueryParams) (*EventResourcesResponse, error) {
	channels := &EventResourcesResponse{}
	path := fmt.Sprintf("/v2/sites/%s/channels/%s/events", siteID, channelID)
	err := c.v2Client.Request(http.MethodGet, path, channels, nil, queryParams)
	return channels, err
}

// RequestMaster reqyests the master asset resources associated with a given Site ID.
func (c *EventsClient) RequestMaster(siteID, channelID, eventID string) error {
	path := fmt.Sprintf("/v2/sites/%s/channels/%s/events/%s/request_master", siteID, channelID, eventID)
	err := c.v2Client.Request(http.MethodPut, path, nil, nil, nil)
	return err
}

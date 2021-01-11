package jwplatform

import (
	"fmt"
	"net/http"
)

// ChannelResource is the resource that is returned for all Channel resource requests
type ChannelResource struct {
	V2ResourceResponse

	Metadata        ChannelMetadata `json:"metadata"`
	Latency         string          `json:"latency"`
	RecentEvents    []recentEvent   `json:"recent_events"`
	ReconnectWindow int             `json:"reconnect_window"`
	Status          string          `json:"status"`
	StreamKey       string          `json:"stream_key"`
}

type recentEvent struct {
	MediaID string `json:"media_id"`
	Status  string `json:"status"`
}

// ChannelCreateRequest is the request structure required for Channel create calls.
type ChannelCreateRequest struct {
	Metadata ChannelCreateMetadata `json:"metadata"`
}

// ChannelUpdateRequest is the request structure required for Channel update calls.
type ChannelUpdateRequest struct {
	Metadata ChannelMetadata `json:"metadata"`
}

// ChannelMetadata describes a Channel resource
type ChannelMetadata struct {
	CustomParams     map[string]string `json:"custom_params"`
	Dvr              string            `json:"dvr"`
	SimulcastTargets []simulcastTarget `json:"simulcast_targets"`
	Tags             []string          `json:"tags"`
	Title            string            `json:"title"`
}

type simulcastTarget struct {
	StreamKey string `json:"stream_key"`
	StreamURL string `json:"stream_url"`
	Title     string `json:"title"`
}

// ChannelCreateMetadata describes the request structure used to create a Channel resource
type ChannelCreateMetadata struct {
	CustomParams     map[string]string `json:"custom_params"`
	Dvr              string            `json:"dvr"`
	SimulcastTargets []simulcastTarget `json:"simulcast_targets"`
	Tags             []string          `json:"tags"`
	Title            string            `json:"title"`
	Latency          string            `json:"latency"`
	ReconnectWindow  int               `json:"reconnect_window"`
}

// ChannelResourcesResponse is the response structure for Channel list calls.
type ChannelResourcesResponse struct {
	V2ResourcesResponse

	Channels []ChannelResource `json:"channels"`
}

// ChannelsClient for interacting with V2 Channels and Channel Events API.
type ChannelsClient struct {
	v2Client *V2Client
	Events   *EventsClient
}

// NewChannelsClient returns a new Channels Client
func NewChannelsClient(v2Client *V2Client) *ChannelsClient {
	return &ChannelsClient{
		v2Client: v2Client,
		Events: &EventsClient{
			v2Client: v2Client,
		},
	}
}

// Get a single Channel resource by ID.
func (c *ChannelsClient) Get(siteID, channelID string) (*ChannelResource, error) {
	channel := &ChannelResource{}
	path := fmt.Sprintf("/v2/sites/%s/channels/%s", siteID, channelID)
	err := c.v2Client.Request(http.MethodGet, path, channel, nil, nil)
	return channel, err
}

// Create a Channel resource.
func (c *ChannelsClient) Create(siteID string, ChannelCreateMetadata *ChannelCreateMetadata) (*ChannelResource, error) {
	createRequestData := &ChannelCreateRequest{Metadata: *ChannelCreateMetadata}
	channel := &ChannelResource{}
	path := fmt.Sprintf("/v2/sites/%s/channels", siteID)
	err := c.v2Client.Request(http.MethodPost, path, channel, createRequestData, nil)
	return channel, err
}

// List all Channel resources associated with a given Site ID.
func (c *ChannelsClient) List(siteID string, queryParams *QueryParams) (*ChannelResourcesResponse, error) {
	channels := &ChannelResourcesResponse{}
	path := fmt.Sprintf("/v2/sites/%s/channels", siteID)
	err := c.v2Client.Request(http.MethodGet, path, channels, nil, queryParams)
	return channels, err
}

// Update a Channel resource by ID.
func (c *ChannelsClient) Update(siteID, channelID string, channelMetadata *ChannelMetadata) (*ChannelResource, error) {
	updateRequestData := &ChannelUpdateRequest{Metadata: *channelMetadata}
	channel := &ChannelResource{}
	path := fmt.Sprintf("/v2/sites/%s/channels/%s", siteID, channelID)
	err := c.v2Client.Request(http.MethodPatch, path, channel, updateRequestData, nil)
	return channel, err
}

// Delete a Channel resource by ID.
func (c *ChannelsClient) Delete(siteID, channelID string) error {
	path := fmt.Sprintf("/v2/sites/%s/channels/%s", siteID, channelID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

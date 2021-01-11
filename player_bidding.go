package jwplatform

import (
	"fmt"
	"net/http"
)

// PlayerBiddingConfigurationResource is the resource that is returned for all Player Bidding Configuration resource requests,
// with the exception of the Create action, which extends this struct with upload-related data.
type PlayerBiddingConfigurationResource struct {
	V2ResourceResponse
	Metadata PlayerBiddingConfigurationMetadata `json:"metadata"`
}

// PlayerBiddingConfigurationMetadata describes the metadata for an Player Bidding Configuration resource
type PlayerBiddingConfigurationMetadata struct {
	Bids BidsMetadata `json:"bids"`
}

// BidsMetadata represents the player bidding configuration as used by the JW Player
type BidsMetadata struct {
	Settings BidSettingsMetadata `json:"settings"`
	Bidders  []BiddersMetadata   `json:"bidders"`
}

// BidSettingsMetadata represents the configuration for the player bidding plugin
type BidSettingsMetadata struct {
	BidTimeout             int      `json:"bidTimeout"`
	FloorPriceCents        int      `json:"floorPriceCents"`
	MediationLayerAdServer string   `json:"mediationLayerAdServer"`
	Buckets                []Bucket `json:"buckets"`
}

// Bucket represents a minimum, maximum value to which to apply an increment
type Bucket struct {
	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
	Increment float64 `json:"increment"`
}

// BiddersMetadata describes a configured Player Bidding bidder
type BiddersMetadata struct {
	Name         string            `json:"name"`
	ID           string            `json:"id"`
	PubID        string            `json:"pubid"`
	CustomParams map[string]string `json:"custom_params"`
}

// PlayerBiddingWriteRequest is the request structure required for Player Bidding Configuration create and update calls.
type PlayerBiddingWriteRequest struct {
	Metadata PlayerBiddingConfigurationMetadata `json:"metadata"`
}

// PlayerBiddingConfigurationResourcesResponse is the response structure for Player Bidding Configuration list calls.
type PlayerBiddingConfigurationResourcesResponse struct {
	V2ResourcesResponse
	PlayerBiddingConfigs []PlayerBiddingConfigurationResource `json:"vpb_configs"`
}

// PlayerBiddingClient for interacting with V2 Player Bidding Configurations API.
type PlayerBiddingClient struct {
	v2Client *V2Client
}

// Get a single Player Bidding Configuration resource by ID.
func (c *PlayerBiddingClient) Get(siteID, importID string) (*PlayerBiddingConfigurationResource, error) {
	playerBiddingConfig := &PlayerBiddingConfigurationResource{}
	path := fmt.Sprintf("/v2/sites/%s/vpb_configs/%s", siteID, importID)
	err := c.v2Client.Request(http.MethodGet, path, playerBiddingConfig, nil, nil)
	return playerBiddingConfig, err
}

// Create a Player Bidding Configuration resource.
func (c *PlayerBiddingClient) Create(siteID string, PlayerBiddingConfigurationMetadata *PlayerBiddingConfigurationMetadata) (*PlayerBiddingConfigurationResource, error) {
	createRequestData := &PlayerBiddingWriteRequest{Metadata: *PlayerBiddingConfigurationMetadata}
	playerBiddingConfig := &PlayerBiddingConfigurationResource{}
	path := fmt.Sprintf("/v2/sites/%s/vpb_configs", siteID)
	err := c.v2Client.Request(http.MethodPost, path, playerBiddingConfig, createRequestData, nil)
	return playerBiddingConfig, err
}

// List all Player Bidding Configuration resources associated with a given Site ID.
func (c *PlayerBiddingClient) List(siteID string, queryParams *QueryParams) (*PlayerBiddingConfigurationResourcesResponse, error) {
	playerBiddingConfigs := &PlayerBiddingConfigurationResourcesResponse{}
	path := fmt.Sprintf("/v2/sites/%s/vpb_configs", siteID)
	err := c.v2Client.Request(http.MethodGet, path, playerBiddingConfigs, nil, queryParams)
	return playerBiddingConfigs, err
}

// Update a Player Bidding Configuration resource by ID.
func (c *PlayerBiddingClient) Update(siteID, importID string, PlayerBiddingConfigurationMetadata *PlayerBiddingConfigurationMetadata) (*PlayerBiddingConfigurationResource, error) {
	updateRequestData := &PlayerBiddingWriteRequest{Metadata: *PlayerBiddingConfigurationMetadata}
	playerBiddingConfig := &PlayerBiddingConfigurationResource{}
	path := fmt.Sprintf("/v2/sites/%s/vpb_configs/%s", siteID, importID)
	err := c.v2Client.Request(http.MethodPatch, path, playerBiddingConfig, updateRequestData, nil)
	return playerBiddingConfig, err
}

// Delete a Player Bidding Configuration resource by ID.
func (c *PlayerBiddingClient) Delete(siteID, importID string) error {
	path := fmt.Sprintf("/v2/sites/%s/vpb_configs/%s", siteID, importID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

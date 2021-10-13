package jwplatform

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// PlayerResource is the resource that is returned for all DRM Policy resource requests
type PlayerResource struct {
	V2ResourceResponse
	Metadata PlayerMetadata `json:"metadata"`
}

type PlayerRelationships struct {
	RecommendationsPlaylistID string
	AdScheduleID              string
}

// PlayerWriteRequest is the request structure required for Player create and update calls.
type PlayerWriteRequest struct {
	Metadata      PlayerMetadata         `json:"metadata"`
	Relationships map[string]interface{} `json:"relationships"`
}

// PlayerMetadata describes a Player resource
type PlayerMetadata struct {
	Name           string                 `json:"name"`
	ReleaseChannel string                 `json:"release_channel"`
	SetupConfig    map[string]interface{} `json:"setup_config"`
	CustomParams   map[string]string      `json:"custom_params"`
}

// PlayerResourcesResponse is the response structure for Player list calls.
type PlayerResourcesResponse struct {
	V2ResourcesResponse

	Players []PlayerResource `json:"players"`
}

// PlayersClient for interacting with V2 DRM Policies API.
type PlayersClient struct {
	v2Client *V2Client
}

// Get a single Player resource by ID.
func (c *PlayersClient) Get(siteID, playerID string) (*PlayerResource, error) {
	player := &PlayerResource{}
	path := fmt.Sprintf("/v2/sites/%s/players/%s", siteID, playerID)
	err := c.v2Client.Request(http.MethodGet, path, player, nil, nil)
	return player, err
}

// Create a Player resource.
func (c *PlayersClient) Create(siteID string, playerMetadata *PlayerMetadata, playerRelationships *PlayerRelationships) (*PlayerResource, error) {
	relationshipsMap := make(map[string]interface{})
	if playerRelationships != nil {
		if playerRelationships.RecommendationsPlaylistID != "" {
			relationshipsMap["recommendations_playlist"] = map[string]string{
				"id": playerRelationships.RecommendationsPlaylistID,
			}
		}

		if playerRelationships.AdScheduleID != "" {
			relationshipsMap["adschedule"] = map[string]string{
				"id": playerRelationships.AdScheduleID,
			}
		}
	}

	createRequestData := &PlayerWriteRequest{Metadata: *playerMetadata, Relationships: relationshipsMap}
	player := &PlayerResource{}
	path := fmt.Sprintf("/v2/sites/%s/players", siteID)
	err := c.v2Client.Request(http.MethodPost, path, player, createRequestData, nil)
	return player, err
}

// List all Player resources.
func (c *PlayersClient) List(siteID string, queryParams *QueryParams) (*PlayerResourcesResponse, error) {
	drmPolicies := &PlayerResourcesResponse{}
	path := fmt.Sprintf("/v2/sites/%s/players", siteID)
	urlValues, _ := query.Values(queryParams)
	err := c.v2Client.Request(http.MethodGet, path, drmPolicies, nil, urlValues)
	return drmPolicies, err
}

// Update a Player resource by ID.
func (c *PlayersClient) Update(siteID, playerID string, playerMetadata *PlayerMetadata) (*PlayerResource, error) {
	updateRequestData := &PlayerWriteRequest{Metadata: *playerMetadata}
	player := &PlayerResource{}
	path := fmt.Sprintf("/v2/sites/%s/players/%s", siteID, playerID)
	err := c.v2Client.Request(http.MethodPatch, path, player, updateRequestData, nil)
	return player, err
}

// Delete a Player resource by ID.
func (c *PlayersClient) Delete(siteID, playerID string) error {
	path := fmt.Sprintf("/v2/sites/%s/players/%s", siteID, playerID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

package jwplatform

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// MediaRendition is the resource that is returned for all Media Rendition resource requests,
type MediaRenditionResource struct {
	V2ResourceResponse
	Metadata      MediaRenditionMetadata `json:"metadata"`
	DeliveryURL   string                 `json:"delivery_url"`
	ErrorMessage  string                 `json:"error_message"`
	FileSize      int                    `json:"filesize"`
	Height        int                    `json:"height"`
	Width         int                    `json:"width"`
	MediaType     string                 `json:"media_type"`
	Status        string                 `json:"status"`
	Relationships map[string]interface{} `json:"relationships"`
}

// MediaRenditionWriteRequest is the request structure required for Media Rendition create and update calls.
type MediaRenditionWriteRequest struct {
	Metadata      MediaRenditionMetadata `json:"metadata"`
	Relationships map[string]interface{} `json:"relationships"`
}

// MediaRenditionMetadata describes a Media Rendition resource
type MediaRenditionMetadata struct{}

// MediaRenditionResourcesResponse is the response structure for Media Rendition list calls.
type MediaRenditionResourcesResponse struct {
	V2ResourcesResponse
	MediaRenditions []MediaRenditionResource `json:"media_renditions"`
}

// MediaRenditionsClient for interacting with V2 Media Rendition API.
type MediaRenditionsClient struct {
	v2Client *V2Client
}

// Get a single Media Rendition resource by ID.
func (c *MediaRenditionsClient) Get(siteID, mediaID, mediaRenditionID string) (*MediaRenditionResource, error) {
	mediaRendition := &MediaRenditionResource{}
	path := fmt.Sprintf("/v2/sites/%s/media/%s/media_renditions/%s", siteID, mediaID, mediaRenditionID)
	err := c.v2Client.Request(http.MethodGet, path, mediaRendition, nil, nil)
	return mediaRendition, err
}

// Create a Media Rendition resource.
func (c *MediaRenditionsClient) Create(siteID, mediaID, templateID string) (*MediaRenditionResource, error) {
	metadata := &MediaRenditionMetadata{}
	templateRelationship := map[string]interface{}{
		"media_rendition_template": map[string]string{
			"id": templateID,
		},
	}
	createRequestData := &MediaRenditionWriteRequest{Metadata: *metadata, Relationships: templateRelationship}
	mediaRendition := &MediaRenditionResource{}
	path := fmt.Sprintf("/v2/sites/%s/media/%s/media_renditions", siteID, mediaID)
	err := c.v2Client.Request(http.MethodPost, path, mediaRendition, createRequestData, nil)
	return mediaRendition, err
}

// List all Media Rendition resources.
func (c *MediaRenditionsClient) List(siteID, mediaID string, queryParams *QueryParams) (*MediaRenditionResourcesResponse, error) {
	mediaRenditions := &MediaRenditionResourcesResponse{}
	urlValues, _ := query.Values(queryParams)
	path := fmt.Sprintf("/v2/sites/%s/media/%s/media_renditions", siteID, mediaID)
	err := c.v2Client.Request(http.MethodGet, path, mediaRenditions, nil, urlValues)
	return mediaRenditions, err
}

// Delete a Media Rendition resource by ID.
func (c *MediaRenditionsClient) Delete(siteID, mediaID, mediaRenditionID string) error {
	path := fmt.Sprintf("/v2/sites/%s/media/%s/media_renditions/%s", siteID, mediaID, mediaRenditionID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

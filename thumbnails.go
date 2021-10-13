package jwplatform

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// ThumbnailResource is the resource that is returned for all Thumbnail resource requests,
// with the exception of the Create action, which extends this struct with upload-related data.
type ThumbnailResource struct {
	V2ResourceResponse
	Relationships ThumbnailRelationships `json:"relationships"`
	ThumbnailType string                 `json:"thumbnail_type"`
	DeliveryURL   string                 `json:"delivery_url"`
	SourceType    string                 `json:"source_type"`
	Metadata      ThumbnailMetadata      `json:"metadata"`
}

// CreateThumbnailResponse is the response structure for Thumbnail create calls.
// If "direct" or "multipart" were selected as the upload method, the response includes additional data required to complete your upload.
//
// For direct uploads, the UploadLink will return a pre-signed S3 Upload Link.
//
// For multipart uploads, the UploadToken and UploadID will be returned, to be used
// in subsequent requests to the V2 Upload API.
type CreateThumbnailResponse struct {
	V2ResourceResponse
	Metadata      ThumbnailMetadata      `json:"metadata"`
	Relationships ThumbnailRelationships `json:"relationships"`
	UploadLink    string                 `json:"upload_link,omitempty"`
	UploadToken   string                 `json:"upload_token,omitempty"`
	UploadID      string                 `json:"upload_id,omitempty"`
}

// CreateThumbnailRequest is the request structure required for Thumbnail create calls.
// By default, the 'direct' upload method is used.
type CreateThumbnailRequest struct {
	Metadata      ThumbnailMetadata      `json:"metadata"`
	Upload        ThumbnailUpload        `json:"upload"`
	Relationships ThumbnailRelationships `json:"relationships"`
}

// UpdateThumbnailRequest is the request structure required for Thumbnail update calls.
type UpdateThumbnailRequest struct {
	Relationships ThumbnailRelationships `json:"relationships"`
	Metadata      ThumbnailMetadata      `json:"metadata"`
}

// ThumbnailRelationships describes the media relationship for a thumbnail.
type ThumbnailRelationships struct {
	Media []ThumbnailMedia `json:"media"`
}

// ThumbnailMedia describes the media and its relationship to the thumbnail.
type ThumbnailMedia struct {
	MediaID  string `json:"id"`
	IsPoster bool   `json:"is_poster"`
}

// Upload contains the data used to describe the upload.
// Available upload method's include "direct" (default), "multipart", "external", and "fetch".
//
// Direct uploads can be used for assets up to 5GB.
//
// MimeType and SourceURL are required only for "direct" and "fetch", respectively.
type ThumbnailUpload struct {
	Method          string `json:"method,omitempty"`
	ThumbnailType   string `json:"thumbnail_type,omitempty"`
	DownloadURL     string `json:"download_url,omitempty"`
	SourceMediaID   string `json:"source_media_id,omitempty"`
	VideoPosition   int    `json:"video_position,omitempty"`
	ThumbstripIndex int    `json:"thumbstrip_index,omitempty"`
}

// ThumbnailMetadata describes a Thumbnail resource
type ThumbnailMetadata struct{}

// ThumbnailResourcesResponse is the response structure for Thumbnail list calls.
type ThumbnailResourcesResponse struct {
	V2ResourcesResponse
	Thumbnails []ThumbnailResource `json:"thumbnail"`
}

// ThumbnailsClient for interacting with V2 Thumbnail API.
type ThumbnailsClient struct {
	v2Client *V2Client
}

// Get a single Thumbnail resource by ID.
func (c *ThumbnailsClient) Get(siteID, thumbnailID string) (*ThumbnailResource, error) {
	media := &ThumbnailResource{}
	path := fmt.Sprintf("/v2/sites/%s/thumbnails/%s", siteID, thumbnailID)
	err := c.v2Client.Request(http.MethodGet, path, media, nil, nil)
	return media, err
}

// Create a Thumbnail resource.
func (c *ThumbnailsClient) Create(siteID string, thumbnailCreateRequest *CreateThumbnailRequest) (*CreateThumbnailResponse, error) {
	media := &CreateThumbnailResponse{}
	path := fmt.Sprintf("/v2/sites/%s/thumbnails", siteID)
	err := c.v2Client.Request(http.MethodPost, path, media, thumbnailCreateRequest, nil)
	return media, err
}

// List all Thumbnail resources associated with a given Site ID.
func (c *ThumbnailsClient) List(siteID string, queryParams *QueryParams) (*ThumbnailResourcesResponse, error) {
	media := &ThumbnailResourcesResponse{}
	path := fmt.Sprintf("/v2/sites/%s/thumbnails", siteID)
	urlValues, _ := query.Values(queryParams)
	err := c.v2Client.Request(http.MethodGet, path, media, nil, urlValues)
	return media, err
}

// Update a Thumbnail resource by ID.
func (c *ThumbnailsClient) Update(siteID, thumbnailID string, updateThumbnailRequest *UpdateThumbnailRequest) (*ThumbnailResource, error) {
	media := &ThumbnailResource{}
	path := fmt.Sprintf("/v2/sites/%s/thumbnails/%s", siteID, thumbnailID)
	err := c.v2Client.Request(http.MethodPatch, path, media, updateThumbnailRequest, nil)
	return media, err
}

// Delete a Thumbnail resource by ID.
func (c *ThumbnailsClient) Delete(siteID, thumbnailID string) error {
	path := fmt.Sprintf("/v2/sites/%s/thumbnails/%s", siteID, thumbnailID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

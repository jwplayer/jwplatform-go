package jwplatform

import (
	"fmt"
	"net/http"
)

// MediaResource is the resource that is returned for all Media resource requests,
// with the exception of the Create action, which extends this struct with upload-related data.
type MediaResource struct {
	V2ResourceResponse

	Duration     float64 `json:"duration"`
	ExternalID   string  `json:"external_id"`
	TrimInPoint  string  `json:"trim_in_point"`
	TrimOutPoint string  `json:"trim_out_point"`
	Status       string  `json:"status"`
	ErrorMessage string  `json:"error_message"`
	MimeType     string  `json:"mime_type"`
	MediaType    string  `json:"media_type"`
	HostingType  string  `json:"hosting_type"`
	SourceURL    string  `json:"source_url"`

	Metadata MediaMetadata `json:"metadata"`
}

// CreateMediaResponse is the response structure for Media create calls.
// If "direct" or "multipart" were selected as the upload method, the response includes additional data required to complete your upload.
//
// For direct uploads, the UploadLink will return a pre-signed S3 Upload Link.
//
// For multipart uploads, the UploadToken and UploadID will be returned, to be used
// in subsequent requests to the V2 Upload API.
type CreateMediaResponse struct {
	V2ResourceResponse
	MediaResource
	UploadLink  string `json:"upload_link,omitempty"`
	UploadToken string `json:"upload_token,omitempty"`
	UploadID    string `json:"upload_id,omitempty"`
}

// CreateMediaRequest is the request structure required for Media create calls.
// By default, the 'direct' upload method is used.
type CreateMediaRequest struct {
	Metadata MediaMetadata `json:"metadata,omitempty"`
	Upload   Upload        `json:"upload,omitempty"`
}

// ReuploadRequest is the request structure required for Media reupload calls.
type ReuploadRequest struct {
	Upload Upload `json:"upload"`
}

// UpdateMediaRequest is the request structure required for Media update calls.
type UpdateMediaRequest struct {
	Metadata MediaMetadata `json:"metadata"`
}

// Upload contains the data used to describe the upload.
// Available upload method's include "direct" (default), "multipart", "external", and "fetch".
//
// Direct uploads can be used for assets up to 5GB.
//
// MimeType and SourceURL are required only for "direct" and "fetch", respectively.
//
// TrimInPoint and TrimOutPoint cannot be specified for "external".
type Upload struct {
	Method       string `json:"method,omitempty"`
	MimeType     string `json:"mime_type,omitempty"`
	SourceURL    string `json:"source_url,omitempty"`
	TrimInPoint  string `json:"trim_in_point,omitempty"`
	TrimOutPoint string `json:"trim_out_point,omitempty"`
}

// MediaMetadata describes a Media resource
type MediaMetadata struct {
	Title            string            `json:"title,omitempty"`
	Description      string            `json:"description,omitempty"`
	Author           string            `json:"author,omitempty"`
	Permalink        string            `json:"permalink,omitempty"`
	Category         string            `json:"category,omitempty"`
	PublishStartDate string            `json:"publish_start_date,omitempty"`
	PublishEndDate   string            `json:"publish_end_date,omitempty"`
	Tags             []string          `json:"tags,omitempty"`
	CustomParams     map[string]string `json:"custom_params,omitempty"`
	ExternalID       string            `json:"external_id,omitempty"`
}

// MediaResourcesResponse is the response structure for Media list calls.
type MediaResourcesResponse struct {
	V2ResourcesResponse
	Media []MediaResource `json:"media"`
}

// MediaClient for interacting with V2 Media API.
type MediaClient struct {
	v2Client *V2Client
}

// Get a single Media resource by ID.
func (c *MediaClient) Get(siteID, mediaID string) (*MediaResource, error) {
	media := &MediaResource{}
	path := fmt.Sprintf("/v2/sites/%s/media/%s", siteID, mediaID)
	err := c.v2Client.Request(http.MethodGet, path, media, nil, nil)
	return media, err
}

// Create a Media resource.
func (c *MediaClient) Create(siteID string, mediaMetadata *MediaMetadata) (*CreateMediaResponse, error) {
	createRequestData := &CreateMediaRequest{Metadata: *mediaMetadata}
	media := &CreateMediaResponse{}
	path := fmt.Sprintf("/v2/sites/%s/media", siteID)
	err := c.v2Client.Request(http.MethodPost, path, media, createRequestData, nil)
	return media, err
}

// List all Media resources associated with a given Site ID.
func (c *MediaClient) List(siteID string, queryParams *QueryParams) (*MediaResourcesResponse, error) {
	media := &MediaResourcesResponse{}
	path := fmt.Sprintf("/v2/sites/%s/media", siteID)
	err := c.v2Client.Request(http.MethodGet, path, media, nil, queryParams)
	return media, err
}

// Update a Media resource by ID.
func (c *MediaClient) Update(siteID, mediaID string, mediaMetadata *MediaMetadata) (*MediaResource, error) {
	updateRequestData := &UpdateMediaRequest{Metadata: *mediaMetadata}
	media := &MediaResource{}
	path := fmt.Sprintf("/v2/sites/%s/media/%s", siteID, mediaID)
	err := c.v2Client.Request(http.MethodPatch, path, media, updateRequestData, nil)
	return media, err
}

// Delete a Media resource by ID.
func (c *MediaClient) Delete(siteID, mediaID string) error {
	path := fmt.Sprintf("/v2/sites/%s/media/%s", siteID, mediaID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

// Reupload a Media resource by ID.
func (c *MediaClient) Reupload(siteID, mediaID string, upload *Upload) (*CreateMediaResponse, error) {
	reuploadRequest := &ReuploadRequest{Upload: *upload}
	media := &CreateMediaResponse{}
	path := fmt.Sprintf("/v2/sites/%s/media/%s/reupload", siteID, mediaID)
	err := c.v2Client.Request(http.MethodPost, path, media, reuploadRequest, nil)
	return media, err
}

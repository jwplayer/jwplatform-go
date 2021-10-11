package jwplatform

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Original is the resource that is returned for all Original resource requests.
// Create responses may include an upload link if the "direct" upload method is used,
// or an `upload_id` and `upload_token` if "fetch" is used.
type OriginalResource struct {
	V2ResourceResponse
	Metadata        OriginalMetadata       `json:"metadata"`
	ContainerFormat string                 `json:"container_format"`
	Includes        map[string]interface{} `json:"includes"`
	MD5             string                 `json:"md5"`
	Size            int                    `json:"size"`
	ErrorMessage    string                 `json:"error_message"`
	Status          string                 `json:"status"`
}

// OriginalCreateResponse is the response generated when the "direct" upload method is used on
// create requests. Unlike a standard Original resource response, it includes an upload link to upload
// the track asset to.
type OriginalCreateResponse struct {
	OriginalResource
	UploadLink  string `json:"upload_link,omitempty"`
	UploadID    string `json:"upload_id,omitempty"`
	UploadToken string `json:"upload_token,omitempty"`
}

// OriginalWriteRequest is the request structure required for Original create calls.
type OriginalCreateRequest struct {
	Metadata OriginalMetadata      `json:"metadata"`
	Upload   OriginalUploadRequest `json:"upload"`
}

// OriginalWriteRequest is the request structure required for Original update calls.
type OriginalUpdateRequest struct {
	Metadata OriginalMetadata `json:"metadata"`
}

// OriginalUploadRequest defines the upload request structure when creating an upload
//
// Available upload method's include "direct" (default), "multipart", and "fetch".
//
// DownloadUrl is required for "fetch" uploads.
type OriginalUploadRequest struct {
	Method      string `json:"method"`
	MimeType    string `json:"mime_type,omitempty"`
	DownloadUrl string `json:"download_url,omitempty"`
}

// OriginalMetadata describes a Original resource
type OriginalMetadata struct {
	Language     string `json:"language"`
	LanguageCode string `json:"language_code"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

// OriginalResourcesResponse is the response structure for Original list calls.
type OriginalResourcesResponse struct {
	V2ResourcesResponse
	Originals []OriginalResource `json:"originals"`
}

// OriginalsClient for interacting with V2 Original API.
type OriginalsClient struct {
	v2Client *V2Client
}

// Get a single Original resource by ID.
func (c *OriginalsClient) Get(siteID, mediaID, originalID string) (*OriginalResource, error) {
	original := &OriginalResource{}
	path := fmt.Sprintf("/v2/sites/%s/media/%s/originals/%s", siteID, mediaID, originalID)
	err := c.v2Client.Request(http.MethodGet, path, original, nil, nil)
	return original, err
}

// Create a Original resource.
func (c *OriginalsClient) Create(siteID, mediaID string, originalCreateRequest *OriginalCreateRequest) (*OriginalCreateResponse, error) {
	original := &OriginalCreateResponse{}
	path := fmt.Sprintf("/v2/sites/%s/media/%s/originals", siteID, mediaID)
	err := c.v2Client.Request(http.MethodPost, path, original, originalCreateRequest, nil)
	return original, err
}

// Update a Original resource by ID.
func (c *OriginalsClient) Update(siteID, mediaID, originalID string, originalMetadata *OriginalMetadata) (*OriginalResource, error) {
	updateRequestData := &OriginalUpdateRequest{Metadata: *originalMetadata}
	original := &OriginalResource{}
	path := fmt.Sprintf("/v2/sites/%s/media/%s/originals/%s", siteID, mediaID, originalID)
	err := c.v2Client.Request(http.MethodPatch, path, original, updateRequestData, nil)
	return original, err
}

// List all Original resources.
func (c *OriginalsClient) List(siteID, mediaID string, queryParams *QueryParams) (*OriginalResourcesResponse, error) {
	originals := &OriginalResourcesResponse{}
	urlValues, _ := query.Values(queryParams)
	path := fmt.Sprintf("/v2/sites/%s/media/%s/originals", siteID, mediaID)
	err := c.v2Client.Request(http.MethodGet, path, originals, nil, urlValues)
	return originals, err
}

// Delete a Original resource by ID.
func (c *OriginalsClient) Delete(siteID, mediaID, originalID string) error {
	path := fmt.Sprintf("/v2/sites/%s/media/%s/originals/%s", siteID, mediaID, originalID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

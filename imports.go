package jwplatform

import (
	"fmt"
	"net/http"
)

// ImportResource is the resource that is returned for all Import resource requests,
// with the exception of the Create action, which extends this struct with upload-related data.
type ImportResource struct {
	V2ResourceResponse

	Metadata           ImportReadMetadata `json:"metadata"`
	TotalItemsIngested int                `json:"total_items_ingested"`
	LastImport         string             `json:"last_import"`
}

// ImportReadMetadata describes the read structure of an Import resource metadata.
// This extends the base metadata, including an additional field, Password, which cannot be updated
// on a write call (update/create)
type ImportReadMetadata struct {
	ImportMetadata
	Password string `json:"password"`
}

// ImportMetadata describes the metadata for an Import resource
type ImportMetadata struct {
	URL            string         `json:"url"`
	HostOnImport   bool           `json:"host_on_import"`
	Title          string         `json:"title,omitempty"`
	State          string         `json:"state"`
	Type           string         `json:"type"`
	Username       string         `json:"username,omitempty"`
	Tags           []string       `json:"tags"`
	IngestMetadata IngestMetadata `json:"ingest_metadata"`
	IngestTags     []string       `json:"ingest_tags"`
}

// ImportWriteRequest is the request structure required for Import create calls.
type ImportWriteRequest struct {
	Metadata ImportMetadata `json:"metadata"`
}

// IngestMetadata describes which data will be captured in the import from the
// MRSS feed.
type IngestMetadata struct {
	Captions    bool `json:"captions"`
	Categories  bool `json:"categories"`
	Credits     bool `json:"credits"`
	Description bool `json:"description"`
	Keywords    bool `json:"keywords"`
	PublishDate bool `json:"publish_date"`
	Tags        bool `json:"tags"`
	Thumbnails  bool `json:"thumbnails"`
}

// ImportResourcesResponse is the response structure for Import list calls.
type ImportResourcesResponse struct {
	V2ResourcesResponse
	Imports []ImportResource `json:"imports"`
}

// ImportsClient for interacting with V2 Imports API.
type ImportsClient struct {
	v2Client *V2Client
}

// Get a single Import resource by ID.
func (c *ImportsClient) Get(siteID, importID string) (*ImportResource, error) {
	importResource := &ImportResource{}
	path := fmt.Sprintf("/v2/sites/%s/imports/%s", siteID, importID)
	err := c.v2Client.Request(http.MethodGet, path, importResource, nil, nil)
	return importResource, err
}

// Create a Import resource.
func (c *ImportsClient) Create(siteID string, importMetadata *ImportMetadata) (*ImportResource, error) {
	createRequestData := &ImportWriteRequest{Metadata: *importMetadata}
	importResource := &ImportResource{}
	path := fmt.Sprintf("/v2/sites/%s/imports", siteID)
	err := c.v2Client.Request(http.MethodPost, path, importResource, createRequestData, nil)
	return importResource, err
}

// List all Import resources associated with a given Site ID.
func (c *ImportsClient) List(siteID string, queryParams *QueryParams) (*ImportResourcesResponse, error) {
	importResources := &ImportResourcesResponse{}
	path := fmt.Sprintf("/v2/sites/%s/imports", siteID)
	err := c.v2Client.Request(http.MethodGet, path, importResources, nil, queryParams)
	return importResources, err
}

// Update a Import resource by ID.
func (c *ImportsClient) Update(siteID, importID string, importMetadata *ImportMetadata) (*ImportResource, error) {
	updateRequestData := &ImportWriteRequest{Metadata: *importMetadata}
	importResource := &ImportResource{}
	path := fmt.Sprintf("/v2/sites/%s/imports/%s", siteID, importID)
	err := c.v2Client.Request(http.MethodPatch, path, importResource, updateRequestData, nil)
	return importResource, err
}

// Delete a Import resource by ID.
func (c *ImportsClient) Delete(siteID, importID string) error {
	path := fmt.Sprintf("/v2/sites/%s/imports/%s", siteID, importID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

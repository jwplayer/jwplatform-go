package jwplatform

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// TextTrack is the resource that is returned for all Text Track resource requests.
// Create responses may include an upload link if the "direct" upload method is used.
type TextTrackResource struct {
	V2ResourceResponse
	Metadata     TextTrackMetadata `json:"metadata"`
	TrackKind    string            `json:"track_kind"`
	ErrorMessage string            `json:"error_message"`
	DeliveryURL  string            `json:"delivery_url"`
	Status       string            `json:"status"`
}

// TextTrackCreateResponse is the response generated when the "direct" upload method is used on
// create requests. Unlike a standard Text Track resource response, it includes an upload link to upload
// the track asset to.
type TextTrackCreateResponse struct {
	TextTrackResource
	UploadLink string `json:"upload_link,omitempty"`
}

// TextTrackWriteRequest is the request structure required for Text Track create calls.
type TextTrackCreateRequest struct {
	Metadata TextTrackCreateMetadata `json:"metadata"`
	Upload   TextTrackUploadRequest  `json:"upload"`
}

// TextTrackWriteRequest is the request structure required for Text Track update calls.
type TextTrackUpdateRequest struct {
	Metadata TextTrackMetadata `json:"metadata"`
}

// TextTrackUploadRequest defines the upload request structure when creating an upload
//
// Available upload method's include "direct" (default) and "fetch".
//
// DownloadUrl is required for "fetch" uploads.
//
// FileFormat is required for "direct" uploads, and required for "fetch" uploads when the
// DownloadUrl does not contain an explicit (valid) file extension.
type TextTrackUploadRequest struct {
	FileFormat  string `json:"file_format,omitempty"`
	AutoPublish bool   `json:"auto_publish,omitempty"`
	Method      string `json:"method"`
	MimeType    string `json:"mime_type,omitempty"`
	DownloadUrl string `json:"download_url,omitempty"`
}

// TextTrackCreateMetadata is used to define a Text Track on create
type TextTrackCreateMetadata struct {
	Label     string `json:"label"`
	Srclang   string `json:"srclang"`
	Position  int    `json:"position"`
	TrackKind string `json:"track_kind"`
}

// TextTrackMetadata describes a Text Track resource
type TextTrackMetadata struct {
	Label    string `json:"label"`
	Srclang  string `json:"srclang"`
	Position int    `json:"position"`
}

// TextTrackResourcesResponse is the response structure for Text Track list calls.
type TextTrackResourcesResponse struct {
	V2ResourcesResponse
	TextTracks []TextTrackResource `json:"text_tracks"`
}

// TextTracksClient for interacting with V2 Text Track API.
type TextTracksClient struct {
	v2Client *V2Client
}

// Get a single Text Track resource by ID.
func (c *TextTracksClient) Get(siteID, mediaID, trackID string) (*TextTrackResource, error) {
	textTrack := &TextTrackResource{}
	path := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks/%s", siteID, mediaID, trackID)
	err := c.v2Client.Request(http.MethodGet, path, textTrack, nil, nil)
	return textTrack, err
}

// Create a Text Track resource.
func (c *TextTracksClient) Create(siteID, mediaID string, textTrackCreateRequest *TextTrackCreateRequest) (*TextTrackCreateResponse, error) {
	textTrack := &TextTrackCreateResponse{}
	path := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks", siteID, mediaID)
	err := c.v2Client.Request(http.MethodPost, path, textTrack, textTrackCreateRequest, nil)
	return textTrack, err
}

// Update a Text Track resource by ID.
func (c *TextTracksClient) Update(siteID, mediaID, trackID string, textTrackMetadata *TextTrackMetadata) (*TextTrackResource, error) {
	updateRequestData := &TextTrackUpdateRequest{Metadata: *textTrackMetadata}
	textTrack := &TextTrackResource{}
	path := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks/%s", siteID, mediaID, trackID)
	err := c.v2Client.Request(http.MethodPatch, path, textTrack, updateRequestData, nil)
	return textTrack, err
}

// List all Text Track resources.
func (c *TextTracksClient) List(siteID, mediaID string, queryParams *QueryParams) (*TextTrackResourcesResponse, error) {
	textTracks := &TextTrackResourcesResponse{}
	urlValues, _ := query.Values(queryParams)
	path := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks", siteID, mediaID)
	err := c.v2Client.Request(http.MethodGet, path, textTracks, nil, urlValues)
	return textTracks, err
}

// Delete a Text Track resource by ID.
func (c *TextTracksClient) Delete(siteID, mediaID, trackID string) error {
	path := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks/%s", siteID, mediaID, trackID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

// Publish a Text Track resource, enabling it to be returned in JW Player delivery responses.
func (c *TextTracksClient) Publish(siteID, mediaID, trackID string) error {
	path := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks/%s/publish", siteID, mediaID, trackID)
	err := c.v2Client.Request(http.MethodPut, path, nil, nil, nil)
	return err
}

// Unpublish a Text Track resource, preventing it from being returned in JW Player delivery responses.
func (c *TextTracksClient) Unpublish(siteID, mediaID, trackID string) error {
	path := fmt.Sprintf("/v2/sites/%s/media/%s/text_tracks/%s/unpublish", siteID, mediaID, trackID)
	err := c.v2Client.Request(http.MethodPut, path, nil, nil, nil)
	return err
}

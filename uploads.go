package jwplatform

import (
	"fmt"
	"net/http"
)

// UploadsClient for performing bulk actions on tags across JW Player resources.
type UploadsClient struct {
	v2Client *V2Client
}

// ListUploadParts list all upload parts, both completed and uncompleted, each of which represents a range of bytes for a multipart upload
func (c *UploadsClient) ListUploadParts(uploadID string) error {
	path := fmt.Sprintf("/v2/uploads/%s/parts", uploadID)
	err := c.v2Client.Request(http.MethodGet, path, nil, nil, nil)
	return err
}

// CompleteUpload marks as an upload as complete. All parts must be uploaded to complete an upload.
func (c *UploadsClient) CompleteUpload(uploadID string) error {
	path := fmt.Sprintf("/v2/uploads/%s/complete", uploadID)
	err := c.v2Client.Request(http.MethodPut, path, nil, nil, nil)
	return err
}

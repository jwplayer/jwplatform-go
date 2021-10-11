package jwplatform

import (
	"fmt"
	"net/http"
)

// BulkRemoveTagRequest is the request structure required to remove tags across all playlists and media.
type BulkRemoveTagRequest struct {
	Tag string `json:"tag"`
}

// BulkRenameTagRequest is the request structure required to rename tags across all playlists and media.
type BulkRenameTagRequest struct {
	OldTag string `json:"old_tag"`
	NewTag string `json:"new_tag"`
}

// TagsClient for performing bulk actions on tags across JW Player resources.
type TagsClient struct {
	v2Client *V2Client
}

// BulkRemoveTags removes a given tag across all playlist and media in your site. This is an asynchronous request; the API will return a quick-202 and then process the changes. Depending on the
// number of resources the tag is associated with, the time to complete the operation will vary.
func (c *TagsClient) BulkRemove(siteID, tagName string) error {
	removeTagRequest := &BulkRemoveTagRequest{Tag: tagName}
	path := fmt.Sprintf("/v2/sites/%s/remove_tag", siteID)
	err := c.v2Client.Request(http.MethodPut, path, nil, removeTagRequest, nil)
	return err
}

// BulkRemoveTags renames a given tag across all playlist and media in your site.
// This is an asynchronous request; the API will return a quick-202 and then process the changes. Depending on the
// number of resources the tag is associated with, the time to complete the operation will vary.
func (c *TagsClient) BulkRename(siteID, oldTagName, newTagName string) error {
	renameTagRequest := &BulkRenameTagRequest{OldTag: oldTagName, NewTag: newTagName}
	path := fmt.Sprintf("/v2/sites/%s/rename_tag", siteID)
	err := c.v2Client.Request(http.MethodPut, path, nil, renameTagRequest, nil)
	return err
}

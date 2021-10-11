package jwplatform

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestBulkRenameTags(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/rename_tag", siteID)

	gock.New("https://api.jwplayer.com").
		Put(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(202)

	testClient := New(mockAuthToken)
	err := testClient.Tags.BulkRename(siteID, "old tag", "new tag")
	assert.Equal(t, nil, err)
}

func TestBulkRemoveTags(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/remove_tag", siteID)

	gock.New("https://api.jwplayer.com").
		Put(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(202)

	testClient := New(mockAuthToken)
	err := testClient.Tags.BulkRemove(siteID, "tag to remove")
	assert.Equal(t, nil, err)
}

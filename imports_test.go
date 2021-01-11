package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetImport(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	importID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/imports/%s", siteID, importID)
	mockImportResponse := map[string]string{"id": importID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockImportResponse)

	testClient := New(mockAuthToken)
	channel, err := testClient.Imports.Get(siteID, importID)
	assert.Equal(t, importID, channel.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteImport(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	importID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/imports/%s", siteID, importID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.Imports.Delete(siteID, importID)
	assert.Equal(t, nil, err)
}

func TestCreateImport(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	importID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/imports", siteID)
	mockImportResponse := map[string]string{"id": importID}

	gock.New("https://api.jwplayer.com").
		Post(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockImportResponse)

	testClient := New(mockAuthToken)
	newImport := &ImportMetadata{Title: "Another test video"}
	channel, err := testClient.Imports.Create(siteID, newImport)
	assert.Equal(t, importID, channel.ID)
	assert.Equal(t, nil, err)
}

func TestUpdateImport(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	importID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/sites/%s/imports/%s", siteID, importID)
	mockImportResponse := map[string]string{"id": importID}

	gock.New("https://api.jwplayer.com").
		Patch(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockImportResponse)

	testClient := New(mockAuthToken)
	updateMetadata := &ImportMetadata{Title: "Updating another test media"}
	channel, err := testClient.Imports.Update(siteID, importID, updateMetadata)
	assert.Equal(t, importID, channel.ID)
	assert.Equal(t, nil, err)
}

func TestListImport(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	importID := "mnbvcxkj"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4

	requestPath := fmt.Sprintf("/v2/sites/%s/imports", siteID)
	mockImportsResponse := map[string]interface{}{
		"page_length": pageLength,
		"page":        page,
		"imports":     []map[string]string{{"id": importID}},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		MatchParam("page", strconv.Itoa(page)).
		MatchParam("page_length", strconv.Itoa(pageLength)).
		Reply(200).
		JSON(mockImportsResponse)

	testClient := New(mockAuthToken)
	params := &QueryParams{PageLength: pageLength, Page: page}
	importsResponse, err := testClient.Imports.List(siteID, params)
	assert.Equal(t, page, importsResponse.Page)
	assert.Equal(t, pageLength, importsResponse.PageLength)
	assert.Equal(t, importID, importsResponse.Imports[0].ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalImport(t *testing.T) {
	importData := map[string]interface{}{
		"id":                   "abZqokMz",
		"type":                 "import",
		"created":              "2019-09-25T15:29:11.042095+00:00",
		"last_modified":        "2019-09-25T15:29:11.042095+00:00",
		"total_items_ingested": 50,
		"last_import":          "2019-09-25T15:29:11.042095+00:00",
		"metadata": map[string]interface{}{
			"username":       "username",
			"password":       "password",
			"url":            "http://url.com",
			"host_on_import": true,
			"title":          "test title",
			"state":          "ready",
			"type":           "active",
			"tags":           []string{"tag_a", "tag_b"},
			"ingest_tags":    []string{"ingest_tag_a", "ingest_tag_b"},
			"ingest_metadata": map[string]bool{
				"captions":     true,
				"categories":   false,
				"credits":      false,
				"description":  true,
				"keywords":     false,
				"publish_date": true,
				"tags":         true,
				"thumbnails":   false,
			},
		},
	}

	bytes, err := json.Marshal(&importData)
	assert.NoError(t, err)

	var importResource ImportResource
	err = json.Unmarshal(bytes, &importResource)
	assert.NoError(t, err)

	assert.Equal(t, "abZqokMz", importResource.ID)
	assert.Equal(t, "import", importResource.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", importResource.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", importResource.LastModified)

	assert.Equal(t, "username", importResource.Metadata.Username)
	assert.Equal(t, "password", importResource.Metadata.Password)
	assert.Equal(t, "http://url.com", importResource.Metadata.URL)
	assert.Equal(t, true, importResource.Metadata.HostOnImport)
	assert.Equal(t, "ready", importResource.Metadata.State)
	assert.Equal(t, "active", importResource.Metadata.Type)
	assert.Equal(t, []string{"tag_a", "tag_b"}, importResource.Metadata.Tags)
	assert.Equal(t, []string{"ingest_tag_a", "ingest_tag_b"}, importResource.Metadata.IngestTags)

	assert.Equal(t, true, importResource.Metadata.IngestMetadata.Captions)
	assert.Equal(t, false, importResource.Metadata.IngestMetadata.Categories)
	assert.Equal(t, false, importResource.Metadata.IngestMetadata.Credits)
	assert.Equal(t, true, importResource.Metadata.IngestMetadata.Description)
	assert.Equal(t, false, importResource.Metadata.IngestMetadata.Keywords)
	assert.Equal(t, true, importResource.Metadata.IngestMetadata.PublishDate)
	assert.Equal(t, true, importResource.Metadata.IngestMetadata.Tags)
	assert.Equal(t, false, importResource.Metadata.IngestMetadata.Thumbnails)
}

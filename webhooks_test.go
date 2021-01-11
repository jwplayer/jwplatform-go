package jwplatform

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetWebhook(t *testing.T) {
	defer gock.Off()

	webhookID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/webhooks/%s", webhookID)
	mockWebhookResponse := map[string]string{"id": webhookID}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockWebhookResponse)

	testClient := New(mockAuthToken)
	webhook, err := testClient.Webhooks.Get(webhookID)
	assert.Equal(t, webhookID, webhook.ID)
	assert.Equal(t, nil, err)
}

func TestDeleteWebhook(t *testing.T) {
	defer gock.Off()

	webhookID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/webhooks/%s", webhookID)

	gock.New("https://api.jwplayer.com").
		Delete(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(204)

	testClient := New(mockAuthToken)
	err := testClient.Webhooks.Delete(webhookID)
	assert.Equal(t, nil, err)
}

func TestCreateWebhook(t *testing.T) {
	defer gock.Off()

	webhookID := "mnbvcxkj"
	mockAuthToken := "shhh"

	mockWebhookResponse := map[string]string{"id": webhookID}

	gock.New("https://api.jwplayer.com").
		Post("/v2/webhooks").
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(201).
		JSON(mockWebhookResponse)

	testClient := New(mockAuthToken)
	newWebhook := &WebhookMetadata{Name: "My first webhook", Sites: []string{"abcdefgh"}}
	webhook, err := testClient.Webhooks.Create(newWebhook)
	assert.Equal(t, webhookID, webhook.ID)
	assert.Equal(t, nil, err)
}

func TestUpdateWebhook(t *testing.T) {
	defer gock.Off()

	webhookID := "mnbvcxkj"
	mockAuthToken := "shhh"

	requestPath := fmt.Sprintf("/v2/webhooks/%s", webhookID)
	mockWebhookResponse := map[string]string{"id": webhookID}

	gock.New("https://api.jwplayer.com").
		Patch(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(200).
		JSON(mockWebhookResponse)

	testClient := New(mockAuthToken)
	updateMetadata := &WebhookMetadata{Name: "My first webhook", Sites: []string{"abcdefgh"}}
	webhook, err := testClient.Webhooks.Update(webhookID, updateMetadata)
	assert.Equal(t, webhookID, webhook.ID)
	assert.Equal(t, nil, err)
}

func TestListWebhooks(t *testing.T) {
	defer gock.Off()

	webhookID := "mnbvcxkj"
	mockAuthToken := "shhh"
	page := 2
	pageLength := 4
	mockWebhooksResponse := map[string]interface{}{
		"page_length": pageLength,
		"page":        page,
		"webhooks":    []map[string]string{{"id": webhookID}},
	}

	gock.New("https://api.jwplayer.com").
		Get("/v2/webhooks").
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		MatchParam("page", strconv.Itoa(page)).
		MatchParam("page_length", strconv.Itoa(pageLength)).
		Reply(200).
		JSON(mockWebhooksResponse)

	testClient := New(mockAuthToken)
	params := &QueryParams{PageLength: pageLength, Page: page}
	webhooksResponse, err := testClient.Webhooks.List(params)
	assert.Equal(t, page, webhooksResponse.Page)
	assert.Equal(t, pageLength, webhooksResponse.PageLength)
	assert.Equal(t, webhookID, webhooksResponse.Webhooks[0].ID)
	assert.Equal(t, nil, err)
}

func TestUnmarshalWebhook(t *testing.T) {
	webhookData := map[string]interface{}{
		"id":            "abZqokMz",
		"type":          "webhook",
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"metadata": map[string]interface{}{
			"name":        "Webhook",
			"description": "Describes a webhook",
			"webhook_url": "https://webhook.com",
			"site_ids":    []string{"a", "b"},
			"events":      []string{"event_a", "event_b"},
		},
	}

	bytes, err := json.Marshal(&webhookData)
	assert.NoError(t, err)

	var webhook WebhookResource
	err = json.Unmarshal(bytes, &webhook)
	assert.NoError(t, err)

	assert.Equal(t, "abZqokMz", webhook.ID)
	assert.Equal(t, "webhook", webhook.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", webhook.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", webhook.LastModified)
	assert.Equal(t, "Webhook", webhook.Metadata.Name)
	assert.Equal(t, "Describes a webhook", webhook.Metadata.Description)
	assert.Equal(t, "https://webhook.com", webhook.Metadata.WebhookURL)
	assert.Equal(t, []string{"a", "b"}, webhook.Metadata.Sites)
	assert.Equal(t, []string{"event_a", "event_b"}, webhook.Metadata.Events)
}

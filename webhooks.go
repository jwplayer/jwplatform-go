package jwplatform

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// WebhookResource is the resource that is returned for all Webhook resource requests,
// with the exception of the Create action, which extends this struct with upload-related data.
type WebhookResource struct {
	V2ResourceResponse
	Metadata WebhookMetadata `json:"metadata"`
}

// CreateWebhookResponse is the response structure for Webhook create calls.
//
// The Secret is returned only on Create calls and can be used to authenticate incoming webhooks
// Please see the documentation for additional details:
// https://developer.jwplayer.com/jwplayer/docs/learn-about-webhooks#section-verify-the-authenticity-of-a-webhook
type CreateWebhookResponse struct {
	V2ResourceResponse
	Metadata WebhookMetadata `json:"metadata"`
	Secret   string          `json:"secret"`
}

// WebhookWriteRequest is the request structure required for Webhook create and update calls.
type WebhookWriteRequest struct {
	Metadata WebhookMetadata `json:"metadata"`
}

// WebhookMetadata describes a Webhook resource
type WebhookMetadata struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Events      []string `json:"events"`
	Sites       []string `json:"site_ids"`
	WebhookURL  string   `json:"webhook_url"`
}

// WebhookResourcesResponse is the response structure for Webhook list calls.
type WebhookResourcesResponse struct {
	V2ResourcesResponse
	Webhooks []WebhookResource `json:"webhooks"`
}

// WebhooksClient for interacting with V2 Webhooks API.
type WebhooksClient struct {
	v2Client *V2Client
}

// Get a single Webhook resource by ID.
func (c *WebhooksClient) Get(webhookID string) (*WebhookResource, error) {
	webhook := &WebhookResource{}
	path := fmt.Sprintf("/v2/webhooks/%s", webhookID)
	err := c.v2Client.Request(http.MethodGet, path, webhook, nil, nil)
	return webhook, err
}

// Create a Webhook resource.
func (c *WebhooksClient) Create(webhookMetadata *WebhookMetadata) (*CreateWebhookResponse, error) {
	createRequestData := &WebhookWriteRequest{Metadata: *webhookMetadata}
	webhook := &CreateWebhookResponse{}
	err := c.v2Client.Request(http.MethodPost, "/v2/webhooks", webhook, createRequestData, nil)
	return webhook, err
}

// List all Webhook resources.
func (c *WebhooksClient) List(queryParams *QueryParams) (*WebhookResourcesResponse, error) {
	webhooks := &WebhookResourcesResponse{}
	urlValues, _ := query.Values(queryParams)
	err := c.v2Client.Request(http.MethodGet, "/v2/webhooks", webhooks, nil, urlValues)
	return webhooks, err
}

// Update a Webhook resource by ID.
func (c *WebhooksClient) Update(webhookID string, webhookMetadata *WebhookMetadata) (*WebhookResource, error) {
	updateRequestData := &WebhookWriteRequest{Metadata: *webhookMetadata}
	webhook := &WebhookResource{}
	path := fmt.Sprintf("/v2/webhooks/%s", webhookID)
	err := c.v2Client.Request(http.MethodPatch, path, webhook, updateRequestData, nil)
	return webhook, err
}

// Delete a Webhook resource by ID.
func (c *WebhooksClient) Delete(webhookID string) error {
	path := fmt.Sprintf("/v2/webhooks/%s", webhookID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

package jwplatform

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// MediaProtectionRule is the resource that is returned for all Media Protection Rule resource requests,
type MediaProtectionRuleResource struct {
	V2ResourceResponse
	Metadata MediaProtectionRuleMetadata `json:"metadata"`
}

// MediaProtectionRuleWriteRequest is the request structure required for Media Protection Rule create and update calls.
type MediaProtectionRuleWriteRequest struct {
	Metadata MediaProtectionRuleMetadata `json:"metadata"`
}

// MediaProtectionRuleMetadata describes a Media Protection Rule resource
type MediaProtectionRuleMetadata struct {
	Name      string   `json:"name"`
	RuleType  string   `json:"rule_type"`
	Countries []string `json:"countries"`
}

// MediaProtectionRuleResourcesResponse is the response structure for Media Protection Rule list calls.
type MediaProtectionRuleResourcesResponse struct {
	V2ResourcesResponse
	MediaProtectionRules []MediaProtectionRuleResource `json:"media_protection_rules"`
}

// MediaProtectionRulesClient for interacting with V2 Media Protection Rule API.
type MediaProtectionRulesClient struct {
	v2Client *V2Client
}

// Get a single Media Protection Rule resource by ID.
func (c *MediaProtectionRulesClient) Get(siteID, protectionRuleID string) (*MediaProtectionRuleResource, error) {
	mediaProtectionRule := &MediaProtectionRuleResource{}
	path := fmt.Sprintf("/v2/sites/%s/media_protection_rules/%s", siteID, protectionRuleID)
	err := c.v2Client.Request(http.MethodGet, path, mediaProtectionRule, nil, nil)
	return mediaProtectionRule, err
}

// Create a Media Protection Rule resource.
func (c *MediaProtectionRulesClient) Create(siteID string, mediaProtectionRuleMetadata *MediaProtectionRuleMetadata) (*MediaProtectionRuleResource, error) {
	createRequestData := &MediaProtectionRuleWriteRequest{Metadata: *mediaProtectionRuleMetadata}
	mediaProtectionRule := &MediaProtectionRuleResource{}
	path := fmt.Sprintf("/v2/sites/%s/media_protection_rules", siteID)
	err := c.v2Client.Request(http.MethodPost, path, mediaProtectionRule, createRequestData, nil)
	return mediaProtectionRule, err
}

// List all Media Protection Rule resources.
func (c *MediaProtectionRulesClient) List(siteID string, queryParams *QueryParams) (*MediaProtectionRuleResourcesResponse, error) {
	mediaProtectionRules := &MediaProtectionRuleResourcesResponse{}
	urlValues, _ := query.Values(queryParams)
	path := fmt.Sprintf("/v2/sites/%s/media_protection_rules", siteID)
	err := c.v2Client.Request(http.MethodGet, path, mediaProtectionRules, nil, urlValues)
	return mediaProtectionRules, err
}

// Update a Media Protection Rule resource by ID.
func (c *MediaProtectionRulesClient) Update(siteID, protectionRuleID string, mediaProtectionRuleMetadata *MediaProtectionRuleMetadata) (*MediaProtectionRuleResource, error) {
	updateRequestData := &MediaProtectionRuleWriteRequest{Metadata: *mediaProtectionRuleMetadata}
	mediaProtectionRule := &MediaProtectionRuleResource{}
	path := fmt.Sprintf("/v2/sites/%s/media_protection_rules/%s", siteID, protectionRuleID)
	err := c.v2Client.Request(http.MethodPatch, path, mediaProtectionRule, updateRequestData, nil)
	return mediaProtectionRule, err
}

// Delete a Media Protection Rule resource by ID.
func (c *MediaProtectionRulesClient) Delete(siteID, protectionRuleID string) error {
	path := fmt.Sprintf("/v2/sites/%s/media_protection_rules/%s", siteID, protectionRuleID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

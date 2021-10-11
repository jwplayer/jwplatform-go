package jwplatform

import (
	"fmt"
	"net/http"
)

// SiteProtectionRule is the resource that is returned for all Site Protection Rule resource requests,
type SiteProtectionRuleResource struct {
	V2ResourceResponse
	Metadata SiteProtectionRuleMetadata `json:"metadata"`
}

// SiteProtectionRuleWriteRequest is the request structure required for Site Protection Rule create and update calls.
type SiteProtectionRuleWriteRequest struct {
	Metadata SiteProtectionRuleMetadata `json:"metadata"`
}

// SiteProtectionRuleMetadata describes a Site Protection Rule resource
type SiteProtectionRuleMetadata struct {
	RuleType  string   `json:"rule_type"`
	Countries []string `json:"countries"`
}

// SiteProtectionRuleClient for interacting with V2 Site Protection Rule API.
type SiteProtectionRuleClient struct {
	v2Client *V2Client
}

// Get a single Site Protection Rule resource by ID.
func (c *SiteProtectionRuleClient) Get(siteID string) (*SiteProtectionRuleResource, error) {
	siteProtectionRule := &SiteProtectionRuleResource{}
	path := fmt.Sprintf("/v2/sites/%s/site_protection_rule", siteID)
	err := c.v2Client.Request(http.MethodGet, path, siteProtectionRule, nil, nil)
	return siteProtectionRule, err
}

// Update a Site Protection Rule resource by ID.
func (c *SiteProtectionRuleClient) Update(siteID string, siteProtectionRuleMetadata *SiteProtectionRuleMetadata) (*SiteProtectionRuleResource, error) {
	updateRequestData := &SiteProtectionRuleWriteRequest{Metadata: *siteProtectionRuleMetadata}
	siteProtectionRule := &SiteProtectionRuleResource{}
	path := fmt.Sprintf("/v2/sites/%s/site_protection_rule", siteID)
	err := c.v2Client.Request(http.MethodPatch, path, siteProtectionRule, updateRequestData, nil)
	return siteProtectionRule, err
}

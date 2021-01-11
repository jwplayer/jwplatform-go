package jwplatform

import (
	"fmt"
	"net/http"
)

// DRMPolicyResource is the resource that is returned for all DRM Policy resource requests
type DRMPolicyResource struct {
	V2ResourceResponse
	Metadata DRMPolicyMetadata `json:"metadata"`
}

// DRMPolicyWriteRequest is the request structure required for DRMPolicy create and update calls.
type DRMPolicyWriteRequest struct {
	Metadata DRMPolicyMetadata `json:"metadata"`
}

// DRMPolicyMetadata describes a DRMPolicy resource
type DRMPolicyMetadata struct {
	Name                    string `json:"name"`
	MaxWidth                int    `json:"max_width"`
	WidevineSecurity        string `json:"widevine_security"`
	PlayreadySecurity       int    `json:"playready_security"`
	AllowOfflinePersistence bool   `json:"allow_offline_persistence"`
	DigitalOutputProtection string `json:"digital_output_protection"`
	LicenseDuration         int    `json:"license_duration"`
	PlaybackDuration        int    `json:"playback_duration"`
}

// DRMPolicyResourcesResponse is the response structure for DRMPolicy list calls.
type DRMPolicyResourcesResponse struct {
	V2ResourcesResponse

	DRMPolicies []DRMPolicyResource `json:"drm_policies"`
}

// DRMPoliciesClient for interacting with V2 DRM Policies API.
type DRMPoliciesClient struct {
	v2Client *V2Client
}

// Get a single DRMPolicy resource by ID.
func (c *DRMPoliciesClient) Get(siteID, drmPolicyID string) (*DRMPolicyResource, error) {
	drmPolicy := &DRMPolicyResource{}
	path := fmt.Sprintf("/v2/sites/%s/drm_policies/%s", siteID, drmPolicyID)
	err := c.v2Client.Request(http.MethodGet, path, drmPolicy, nil, nil)
	return drmPolicy, err
}

// Create a DRMPolicy resource.
func (c *DRMPoliciesClient) Create(siteID string, DRMPolicyMetadata *DRMPolicyMetadata) (*DRMPolicyResource, error) {
	createRequestData := &DRMPolicyWriteRequest{Metadata: *DRMPolicyMetadata}
	drmPolicy := &DRMPolicyResource{}
	path := fmt.Sprintf("/v2/sites/%s/drm_policies", siteID)
	err := c.v2Client.Request(http.MethodPost, path, drmPolicy, createRequestData, nil)
	return drmPolicy, err
}

// List all DRMPolicy resources.
func (c *DRMPoliciesClient) List(siteID string, queryParams *QueryParams) (*DRMPolicyResourcesResponse, error) {
	drmPolicies := &DRMPolicyResourcesResponse{}
	path := fmt.Sprintf("/v2/sites/%s/drm_policies", siteID)
	err := c.v2Client.Request(http.MethodGet, path, drmPolicies, nil, queryParams)
	return drmPolicies, err
}

// Update a DRMPolicy resource by ID.
func (c *DRMPoliciesClient) Update(siteID, drmPolicyID string, DRMPolicyMetadata *DRMPolicyMetadata) (*DRMPolicyResource, error) {
	updateRequestData := &DRMPolicyWriteRequest{Metadata: *DRMPolicyMetadata}
	drmPolicy := &DRMPolicyResource{}
	path := fmt.Sprintf("/v2/sites/%s/drm_policies/%s", siteID, drmPolicyID)
	err := c.v2Client.Request(http.MethodPatch, path, drmPolicy, updateRequestData, nil)
	return drmPolicy, err
}

// Delete a DRMPolicy resource by ID.
func (c *DRMPoliciesClient) Delete(siteID, drmPolicyID string) error {
	path := fmt.Sprintf("/v2/sites/%s/drm_policies/%s", siteID, drmPolicyID)
	err := c.v2Client.Request(http.MethodDelete, path, nil, nil, nil)
	return err
}

/*
An API client for the V2 JW Platform. For the API documentation see:
https://developer.jwplayer.com/jwplayer/reference#introduction-to-api-v2

Sample Usage:

		import (
			"github.com/jwplayer/jwplatform-go"
			"github.com/jwplayer/jwplatform-go/media"
		)

		jwplatform := jwplatform.New("API_SECRET")
		siteID := "9kzNUpe4"
		mediaID := "LaJFzc9d"

		// Get a Resource
		media, err := jwplatform.Media.Get(siteID, mediaID)

		// Create a Resource
		mediaToCreate := &jwplatform.MediaMetadata(Title: "My new video")
		media, err := jwplatform.Media.Create(siteID, mediaToCreate)

		// List a Resource
		mediaResources, err := jwplatform.Media.List(siteID, nil)
		// Optionally include query parameters, including page, page length, sort, and filters.
		params := jwplatform.QueryParams{Page: 2, PageLength: 5}
		mediaResources, err := jwplatform.Media.List(siteID, params)

		// Update a Resource
		updateMetadata := &jwplatform.MediaMetadata{Title: "Updated video title"}
		updatedMedia, err := jwplatform.Media.Update(siteID, mediaID, updateMetadata)

		// Delete a Resource
		err := jwplatform.Media.Delete(siteID, mediaID)
		if err != nil {
			fmt.Println("Success")
		}
*/

package jwplatform

// JWPlatform client for interacting with JW Player V2 Platform APIs.
type JWPlatform struct {
	Version              string
	Analytics            *AnalyticsClient
	Channels             *ChannelsClient
	DRMPolicies          *DRMPoliciesClient
	Imports              *ImportsClient
	Media                *MediaClient
	PlayerBidding        *PlayerBiddingClient
	Webhooks             *WebhooksClient
	MediaProtectionRules *MediaProtectionRulesClient
}

// New generates an authenticated client for interacting with JW Player V2 Platform APIs.
func New(apiSecret string) *JWPlatform {
	v2Client := NewV2Client(apiSecret)
	channelsClient := NewChannelsClient(v2Client)
	return &JWPlatform{
		Version:              version,
		Analytics:            &AnalyticsClient{v2Client: v2Client},
		Channels:             channelsClient,
		DRMPolicies:          &DRMPoliciesClient{v2Client: v2Client},
		Imports:              &ImportsClient{v2Client: v2Client},
		Media:                &MediaClient{v2Client: v2Client},
		PlayerBidding:        &PlayerBiddingClient{v2Client: v2Client},
		Webhooks:             &WebhooksClient{v2Client: v2Client},
		MediaProtectionRules: &MediaProtectionRulesClient{v2Client: v2Client},
	}
}

// Private consts
const (
	version    = "1.0.0"
	apiHost    = "api.jwplayer.com"
	apiVersion = "v2"
)

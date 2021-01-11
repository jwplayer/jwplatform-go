package jwplatform

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestConstructQuery(t *testing.T) {
	newQueryParams := &QueryParams{
		Page:       7,
		PageLength: 11,
		Sort:       "created:dsc",
	}
	expectedQsString := "page=7&page_length=11&sort=created%3Adsc"

	urlValues := &url.Values{}
	result := constructQuery(urlValues, newQueryParams)
	assert.Equal(t, expectedQsString, result)
}

func TestConstructPartialQuery(t *testing.T) {
	newQueryParams := &QueryParams{
		PageLength: 23,
	}
	expectedQsString := "page_length=23"

	urlValues := &url.Values{}
	result := constructQuery(urlValues, newQueryParams)
	assert.Equal(t, expectedQsString, result)
}

func TestUrlFromPath(t *testing.T) {
	client := NewV2Client("authToken")
	path := "/v2/sites/abc/media/123"
	expected := "https://api.jwplayer.com/v2/sites/abc/media/123"
	resultUrl, _ := client.urlFromPath(path)
	assert.Equal(t, resultUrl.String(), expected)
}

func TestInvalidBody(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "mnbvcxkj"
	mockAuthToken := "shhh"
	errorCode := "invalid_body"
	errorDescription := "'name' is too long"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s", siteID, mediaID)
	jwErrors := []JWError{
		JWError{
			Code:        errorCode,
			Description: errorDescription,
		}}
	jwResponseError := &JWErrorResponse{
		Errors: jwErrors,
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(400).
		JSON(jwResponseError)

	testClient := New(mockAuthToken)
	_, err := testClient.Media.Get(siteID, mediaID)
	assert.Error(t, err)
	jwErr := err.(*JWErrorResponse)
	assert.Equal(t, 400, jwErr.StatusCode)
	jwError := jwErr.Errors[0]
	assert.Equal(t, errorCode, jwError.Code)
	assert.Equal(t, errorDescription, jwError.Description)
}

func TestNotFound(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "mnbvcxkj"
	mockAuthToken := "shhh"
	errorCode := "not_found"
	errorDescription := "The requested resource could not be found."

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s", siteID, mediaID)
	response := map[string]interface{}{
		"errors": []interface{}{
			map[string]string{
				"code": errorCode,
				"description": errorDescription,
			},
		},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(404).
		JSON(response)

	testClient := New(mockAuthToken)
	_, err := testClient.Media.Get(siteID, mediaID)
	assert.Error(t, err)
	jwErr := err.(*JWErrorResponse)
	assert.Equal(t, 404, jwErr.StatusCode)
	jwError := jwErr.Errors[0]
	assert.Equal(t, errorCode, jwError.Code)
	assert.Equal(t, errorDescription, jwError.Description)
}

func TestUnauthorized(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "mnbvcxkj"
	mockAuthToken := "shhh"
	errorCode := "unauthorized"
	errorDescription := "Unauthorized"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s", siteID, mediaID)
	response := map[string]interface{}{
		"errors": []interface{}{
			map[string]string{
				"code": errorCode,
				"description": errorDescription,
			},
		},
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(403).
		JSON(response)

	testClient := New(mockAuthToken)
	_, err := testClient.Media.Get(siteID, mediaID)
	assert.Error(t, err)
	jwErr := err.(*JWErrorResponse)
	assert.Equal(t, 403, jwErr.StatusCode)
	jwError := jwErr.Errors[0]
	assert.Equal(t, errorCode, jwError.Code)
	assert.Equal(t, errorDescription, jwError.Description)
}

func TestJWErrorHandling(t *testing.T) {
	defer gock.Off()

	siteID := "abcdefgh"
	mediaID := "mnbvcxkj"
	mockAuthToken := "shhh"
	errorCode := "invalid_body"
	errorDescription := "'name' is too long"

	requestPath := fmt.Sprintf("/v2/sites/%s/media/%s", siteID, mediaID)
	jwErrors := []JWError{
		JWError{
			Code:        errorCode,
			Description: errorDescription,
		}}
	jwResponseError := &JWErrorResponse{
		Errors: jwErrors,
	}

	gock.New("https://api.jwplayer.com").
		Get(requestPath).
		MatchHeader("Authorization", "^Bearer .+").
		MatchHeader("User-Agent", "^jwplatform-go/+").
		Reply(400).
		JSON(jwResponseError)

	testClient := New(mockAuthToken)
	_, err := testClient.Media.Get(siteID, mediaID)
	assert.Error(t, err)
	jwErr := err.(*JWErrorResponse)
	assert.Equal(t, 400, jwErr.StatusCode)
	jwError := jwErr.Errors[0]
	assert.Equal(t, errorCode, jwError.Code)
	assert.Equal(t, errorDescription, jwError.Description)
}

func TestUnmarshalResourceResponse(t *testing.T) {
	resourceResponseData := map[string]interface{}{
		"id":            "9jTnCiPO",
		"type":          "resource_type",
		"created":       "2019-09-25T15:29:11.042095+00:00",
		"last_modified": "2019-09-25T15:29:11.042095+00:00",
		"relationships": map[string]interface{}{
			"protectionrule": map[string]string{
				"id": "protectionrule_id",
			},
		},
	}

	bytes, err := json.Marshal(&resourceResponseData)
	assert.NoError(t, err)

	var resourceResponse V2ResourceResponse
	err = json.Unmarshal(bytes, &resourceResponse)
	assert.NoError(t, err)

	assert.Equal(t, "9jTnCiPO", resourceResponse.ID)
	assert.Equal(t, "resource_type", resourceResponse.Type)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", resourceResponse.Created)
	assert.Equal(t, "2019-09-25T15:29:11.042095+00:00", resourceResponse.LastModified)
	assert.Equal(t, map[string]interface{}{"id": "protectionrule_id"}, resourceResponse.Relationships["protectionrule"])
}

func TestUnmarshalResourcesResponse(t *testing.T) {
	resourcesResponseData := map[string]int{
		"page":        1,
		"page_length": 10,
		"total":       541,
	}

	bytes, err := json.Marshal(&resourcesResponseData)
	assert.NoError(t, err)

	var resourcesResponse V2ResourcesResponse
	err = json.Unmarshal(bytes, &resourcesResponse)
	assert.NoError(t, err)

	assert.Equal(t, 541, resourcesResponse.Total)
	assert.Equal(t, 1, resourcesResponse.Page)
	assert.Equal(t, 10, resourcesResponse.PageLength)
}

func TestUnmarshalJWErrorResponse(t *testing.T) {
	errorResponseData := map[string]interface{}{
		"errors": []interface{}{
			map[string]string{
				"code":        "invalid_body",
				"description": "name was too long",
			},
		},
	}

	bytes, err := json.Marshal(&errorResponseData)
	assert.NoError(t, err)

	var errorResponse JWErrorResponse
	err = json.Unmarshal(bytes, &errorResponse)
	assert.NoError(t, err)

	assert.Equal(t, "invalid_body", errorResponse.Errors[0].Code)
	assert.Equal(t, "name was too long", errorResponse.Errors[0].Description)
}

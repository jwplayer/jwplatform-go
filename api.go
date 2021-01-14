package jwplatform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// V2Client is a light wrapper around the http.defaultClient for interacting with JW Player V2 Platform APIs
type V2Client struct {
	Version   string
	authToken string
	baseURL   *url.URL
	client    *http.Client
}

// V2ResourcesResponse describes the response structure for list calls
type V2ResourcesResponse struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	PageLength int `json:"page_length"`
}

// V2ResourceResponse describes the response structure for resource calls
type V2ResourceResponse struct {
	ID            string                 `json:"id"`
	Created       string                 `json:"created"`
	LastModified  string                 `json:"last_modified"`
	Type          string                 `json:"type"`
	Relationships map[string]interface{} `json:"relationships"`
}

// QueryParams that can be specified on all resource list calls.
type QueryParams struct {
	PageLength int		`url:"page_length"`
	Page       int		`url:"page"`
	Query      string	`url:"q"`
	Sort       string	`url:"sort"`
}

// JWErrorResponse represents a V2 Platform error response.
type JWErrorResponse struct {
	Errors     []JWError `json:"errors"`
	StatusCode int
}

// JWError represents a single error from the V2 Platform API.
type JWError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type rawError struct {
	Error *JWErrorResponse
}

// Error serializes the error object to JSON and returns it as a string.
func (e *JWErrorResponse) Error() string {
	ret, err := json.Marshal(e)
	if err != nil {
		errorMsg := fmt.Sprintf("Unknown error when parsing JSON response: %s", err)
		return errorMsg
	}
	return string(ret)
}

// NewV2Client creates an authenticated V2 Client.
func NewV2Client(authToken string) *V2Client {
	return &V2Client{
		Version:   version,
		authToken: authToken,
		baseURL: &url.URL{
			Scheme: "https",
			Host:   apiHost,
		},
		client: http.DefaultClient,
	}
}

// Request performs an authenticated HTTP request to the V2 Platform API.
func (c *V2Client) Request(method, path string, response interface{}, data interface{}, queryParams url.Values) error {
	var err error
	requestURL, err := c.urlFromPath(path)

	if queryParams != nil {
		requestURL.RawQuery = queryParams.Encode()
	}

	payload := []byte{}
	if data != nil {
		payload, _ = json.Marshal(data)
	}

	request, err := http.NewRequest(method, requestURL.String(), bytes.NewBuffer(payload))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken))
	request.Header.Set("User-Agent", fmt.Sprintf("jwplatform-go/%s", c.Version))

	err = c.Do(request, &response)
	return err
}

// Do  executes the request and parses V2 Platform API errors.
func (c *V2Client) Do(req *http.Request, v interface{}) error {
	var resp *http.Response
	var err error

	resp, err = c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var raw rawError
		var e JWErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return err
		}
		raw.Error = &e
		raw.Error.StatusCode = resp.StatusCode
		return raw.Error
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	switch {
	case err == io.EOF:
		return nil
	case err != nil:
		return err
	}
	return err
}

func (c *V2Client) urlFromPath(path string) (*url.URL, error) {
	url, e := url.Parse(path)
	absoluteURL := c.baseURL.ResolveReference(url)
	return absoluteURL, e
}

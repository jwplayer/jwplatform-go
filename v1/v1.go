/*
Deprecated package providing a client to talk to the V1 JW Platform API.

		import (
		  "github.com/jwplayer/jwplatform-go/v1"
		)

		client := v1.NewClient("API_KEY", "API_SECRET")
*/
package v1

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

const version = "1.0.0"

// Client represents the JWPlatform v1 client object.
type Client struct {
	APIVersion string
	Version    string
	BaseURL    *url.URL
	apiKey     string
	apiSecret  string
	httpClient *http.Client
}

// NewClient creates a V1 new client object.
func NewClient(apiKey string, apiSecret string) *Client {
	return &Client{
		APIVersion: "v1",
		Version:    version,
		BaseURL: &url.URL{
			Scheme: "https",
			Host:   "api.jwplatform.com",
		},
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		httpClient: http.DefaultClient,
	}
}

// generateNonce generates a random 8 digit as a string.
func generateNonce() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return fmt.Sprintf("%08d", rand.Intn(100000000))
}

// makeTimestamp gets the unix timestamp in seconds.
func makeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Second) / int64(time.Nanosecond))
}

// buildParams generates all parameters for api request.
func (c *Client) buildParams(params url.Values) url.Values {
	if params == nil {
		params = url.Values{}
	}
	params.Set("api_nonce", generateNonce())
	params.Set("api_key", c.apiKey)
	params.Set("api_format", "json")
	params.Set("api_timestamp", strconv.FormatInt(makeTimestamp(), 10))

	// create sorted keys array
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// construct signature base string
	var sbs strings.Builder
	for i, k := range keys {
		if i != 0 {
			sbs.WriteString("&")
		}
		// iterate over values of type []string
		for _, val := range params[k] {
			sbs.WriteString(k)
			sbs.WriteString("=")
			sbs.WriteString(val)
		}
	}
	sbs.WriteString(c.apiSecret)

	// hash signature base string
	h := sha1.New()
	h.Write([]byte(sbs.String()))
	sha := hex.EncodeToString(h.Sum(nil))

	params.Set("api_signature", sha)

	return params
}

// newRequestWithContext creates a new request with signed params.
func (c *Client) newRequestWithContext(ctx context.Context, method, pathPart string, params url.Values) (*http.Request, error) {
	rel := &url.URL{Path: path.Join(c.APIVersion, pathPart)}
	absoluteURL := c.BaseURL.ResolveReference(rel)
	absoluteURL.RawQuery = c.buildParams(params).Encode()

	req, err := http.NewRequestWithContext(ctx, method, absoluteURL.String(), nil)
	if err != nil {
		return nil, err
	}

	userAgent := fmt.Sprintf("jwplatform-go/%s", c.Version)

	req.Header.Set("Accept", "application/json")
	req.Header.Add("User-Agent", userAgent)

	return req, nil
}

// do executes request and decodes response body.
func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}

// MakeRequest requests with api signature and decodes json result.
func (c *Client) MakeRequest(ctx context.Context, method, pathPart string, params url.Values, v interface{}) error {
	req, err := c.newRequestWithContext(ctx, method, pathPart, params)
	if err != nil {
		return err
	}

	return c.do(req, &v)
}

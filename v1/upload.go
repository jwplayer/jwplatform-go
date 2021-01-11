package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/mitchellh/go-homedir"
)

// postFile creates a form file and posts it.
func postFile(filepath string, targetURL string) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", filepath)
	if err != nil {
		return nil, err
	}

	fh, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	return http.Post(targetURL, contentType, bodyBuf)
}

// Upload posts a file using the direct upload method.
func (c *Client) Upload(ctx context.Context, filepath string, params url.Values, v interface{}) error {
	// declare an empty interface
	var result map[string]interface{}

	err := c.MakeRequest(ctx, http.MethodPost, "/videos/create/", params, &result)

	if err != nil {
		return err
	}

	if result["status"] != "ok" {
		return fmt.Errorf("Error creating video: %s", result["message"])
	}

	link := result["link"].(map[string]interface{})

	// create upload URL
	uploadURL, err := url.Parse("https://" + fmt.Sprintf("%v%v", link["address"], link["path"]))
	if err != nil {
		return err
	}

	values := url.Values{}
	query := link["query"].(map[string]interface{})
	// create query paramaters from map
	for k, v := range query {
		values.Set(k, fmt.Sprint(v))
	}

	// add query string
	uploadURL.RawQuery = values.Encode() + "&api_format=json"

	abspath, err := homedir.Expand(filepath)
	if err != nil {
		return err
	}

	// upload file
	resp, err := postFile(abspath, uploadURL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}

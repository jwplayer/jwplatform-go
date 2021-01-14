# Go JW Platform

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/jwplayer/jwplatform-go)
[![Build Status](https://travis-ci.org/jwplayer/jwplatform-go.svg?branch=master)](https://travis-ci.org/jwplayer/jwplatform-go)

The official Go client library for accessing the [JW Platform](https://www.jwplayer.com/video-delivery/) API.

## Requirements

Go 1.15+

## Usage

```go
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
_ := jwplatform.Media.Delete(siteID, mediaID)
```

## Supported operations

All API methods documentated on the API are available in this client. Please refer to our [api documentation](https://developer.jwplayer.com/jwplayer/reference#introduction-to-api-v2).

## Test

Before running the tests, make sure to grab all of the package's dependencies:

    go get -t -v

Run all tests:

    make test

For any requests, bug or comments, please [open an issue][issues] or [submit a
pull request][pulls].

## V1 Client

The V1 Client remains available for use, but is deprecated. We strongly recommend using the V2 Client. For documentation on the V1 Client, please refer to the `v1` submodule.

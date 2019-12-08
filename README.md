[![Build Status](https://travis-ci.org/jwplayer/jwplatform-go.svg?branch=master)](https://travis-ci.org/jwplayer/jwplatform-go)

# Go JW Platform

Official Go client library for accessing JW Platform API.

## Requirements

Go 1.13+

## Installation

Install jwplatform-go with:

```sh
go get -u github.com/jwplayer/jwplatform-go
```

### Using Go modules

``` go
module github.com/my/package

require (
    github.com/jwplayer/jwplatform-go v0.2.0
)
```

## Usage

```go
import (
  "github.com/jwplayer/jwplatform-go"
)

client := jwplatform.NewClient("API_KEY", "API_SECRET")
```

### Example: Get video metadata

```go
package main

import (
  "context"
  "fmt"
  "log"
  "net/http"
  "net/url"
  "os"

  "github.com/jwplayer/jwplatform-go"
)

func main() {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  apiKey := os.Getenv("JWPLATFORM_API_KEY")
  apiSecret := os.Getenv("JWPLATFORM_API_SECRET")

  client := jwplatform.NewClient(apiKey, apiSecret)

  // set URL params
  params := url.Values{}
  params.Set("video_key", "VIDEO_KEY")  // some video key, e.g. gIRtMhYM

  // declare an empty interface
  var result map[string]interface{}

  err := client.MakeReque"gopkg.in/h2non/gock.v1"st(ctx, http.MethodGet, "/videos/show/", params, &result)

  if err != nil {
  	log.Fatal(err)
  }

  fmt.Println(result["status"])  // ok
}
```

### Example: Upload video

```go
package main

import (
  "context"
  "fmt"
  "log"
  "net/url"
  "os"

  "github.com/jwplayer/jwplatform-go"
)

func main() {
  filepath := "path/to/your/video.mp4"

  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  // set URL params
  params := url.Values{}
  params.Set("title", "Your video title")
  params.Set("description", "Your video description")

  apiKey := os.Getenv("JWPLATFORM_API_KEY")
  apiSecret := os.Getenv("JWPLATFORM_API_SECRET")

  client := jwplatform.NewClient(apiKey, apiSecret)

  // declare an empty interface
  var result map[string]interface{}

  // upload video usind direct upload method
  err := client.Upload(ctx, filepath, params, &result)

  if err != nil {
  	log.Fatal(err)
  }

  fmt.Println(result["status"])  // ok
}
```

## Supported operations

All API methods documentated on the API are available in this client. Please refer to our [api documentation](https://developer.jwplayer.com/jwplayer/reference).

## Test

Before running the tests, make sure to grab all of the package's dependencies:

    go get -t -v

Run all tests:

    make test

For any requests, bug or comments, please [open an issue][issues] or [submit a
pull request][pulls].

## License

JW Platform API Go library is distributed under the
[Apache v2.0 license](LICENSE).

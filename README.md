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
    github.com/jwplayer/jwplatform-go v0.1.0
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
  params.Set("video_key", "VIDEO_KEY")

  // declare an empty interface
  var result map[string]interface{}

  _, err := client.MakeRequest(ctx, http.MethodGet, "/videos/show/", params, &result)

  if err != nil {
  	log.Fatal(err)
  }

  fmt.Println(result["status"])  // ok
}
```

## Test

The test suite needs testify's `assert` package to run:

    github.com/stretchr/testify/assert

Before running the tests, make sure to grab all of the package's dependencies:

    go get -t -v

Run all tests:

    make test

For any requests, bug or comments, please [open an issue][issues] or [submit a
pull request][pulls].

## License

JW Platform API Go library is distributed under the
[Apache v2.0 license](LICENSE).

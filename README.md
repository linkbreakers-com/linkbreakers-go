# Linkbreakers Go SDK

Official Go SDK for the Linkbreakers API, auto-generated from OpenAPI specification.

## Installation

```bash
go get github.com/linkbreakers-com/linkbreakers-go
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    linkbreakers "github.com/linkbreakers-com/linkbreakers-go"
)

func main() {
    // Create configuration
    cfg := linkbreakers.NewConfiguration()

    // Create API client with Bearer token
    client := linkbreakers.NewAPIClient(cfg)
    ctx := context.WithValue(context.Background(), linkbreakers.ContextAccessToken, "your-api-token")

    // Example: List links
    links, resp, err := client.LinksApi.LinksList(ctx).PageSize(10).Execute()
    if err != nil {
        log.Fatalf("Error listing links: %v", err)
    }

    fmt.Printf("Found %d links\n", len(links.GetLinks()))
}
```

## Authentication

The SDK supports Bearer token authentication:

```go
cfg := linkbreakers.NewConfiguration()
client := linkbreakers.NewAPIClient(cfg)

// Add authentication token to context
ctx := context.WithValue(
    context.Background(),
    linkbreakers.ContextAccessToken,
    "your-api-token",
)

// Use authenticated context in API calls
result, resp, err := client.LinksApi.LinksList(ctx).Execute()
```

## API Documentation

For detailed API documentation, visit: https://docs.linkbreakers.com

## Auto-Generated

This SDK is automatically generated from the Linkbreakers OpenAPI specification and published when the API is updated.

Current API version: See [OPENAPI_VERSION](./OPENAPI_VERSION)

## Issues

Report issues at: https://github.com/linkbreakers-com/linkbreakers-go/issues

## License

MIT License

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
    client := linkbreakers.NewAPIClient(linkbreakers.NewConfiguration())

    // Every request reads the Bearer token from its context
    ctx := context.WithValue(context.Background(), linkbreakers.ContextAccessToken, "your-api-token")

    // Create a shortened link
    request := linkbreakers.NewCreateLinkRequest("https://example.com")
    request.SetName("My Link")

    created, _, err := client.LinksAPI.LinksServiceCreate(ctx).CreateLinkRequest(*request).Execute()
    if err != nil {
        log.Fatalf("Error creating link: %v", err)
    }
    link := created.GetLink()
    fmt.Println("Short link:", link.GetShortlink())

    // List links
    links, _, err := client.LinksAPI.LinksServiceList(ctx).PageSize(10).Execute()
    if err != nil {
        log.Fatalf("Error listing links: %v", err)
    }
    fmt.Printf("Found %d links\n", len(links.GetLinks()))
}
```

Each API is a field on the client (`client.LinksAPI`, `client.VisitorsAPI`, ...) and each operation is named after the API's operation ID (`LinksServiceCreate`, `LinksServiceList`, `VisitorsServiceIdentify`, ...). Path parameters are arguments; query parameters and the request body are builder methods, and `Execute()` sends the request.

## Authentication

The SDK uses Bearer token authentication. Put your workspace API token in the context passed to each call with `linkbreakers.ContextAccessToken`, as in the example above. Get a token from the [Linkbreakers dashboard](https://app.linkbreakers.com/settings/api).

## API Documentation

For detailed API documentation, visit: https://docs.linkbreakers.com

## Auto-Generated

This SDK is automatically generated from the Linkbreakers OpenAPI specification and published when the API is updated. This README is written by hand: `scripts/check-docs.sh` compiles every Go block in it against the generated code, on pull requests and before every release.

Current API version: See [OPENAPI_VERSION](./OPENAPI_VERSION)

## Issues

Report issues at: https://github.com/linkbreakers-com/linkbreakers-go/issues

## License

MIT License

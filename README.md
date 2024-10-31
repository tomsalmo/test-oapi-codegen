# `oapi-codegen` demo

## Starting point

The starting point is `internal/api/oapi.yaml`. That is the OpenAPI v3.0 spec that we would create for our API.

From there we would generate code using [oapi-codegen](github.com/oapi-codegen/oapi-codegen), which is configured in `internal/api/oapi-codegen.yaml` and outputs `internal/api/oapi-codegen.gen.go`.

We use the struct server method that does the url parameter parsing and json unmarshaling for us, and our "handler" signatures are simple. This generates the models, router, and handler definitions.

## Handlers

The handlers can be found in `internal/api/handlers`. These are different methods on the same `Handler` struct, but can be broken into multiple handler structs and added to the main struct with struct embedding.

The handler just needs to conform to the `StrictServerInterface`. This interface helpfully ensures that the return types match what the API Spec is intending.

```go
// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /ping)
	GetPing(ctx context.Context, request GetPingRequestObject) (GetPingResponseObject, error)

	// (POST /tit/{userId})
	PostTitUserId(ctx context.Context, request PostTitUserIdRequestObject) (PostTitUserIdResponseObject, error)
}
```

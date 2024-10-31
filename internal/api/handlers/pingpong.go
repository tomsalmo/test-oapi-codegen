package handlers

import (
	"context"

	"github.com/tomsalmo/test-oapi-codegen/internal/api"
)

func (h Handler) GetPing(ctx context.Context, request api.GetPingRequestObject) (api.GetPingResponseObject, error) {
	return api.GetPing200JSONResponse{}, nil
}

package handlers

import (
	"context"

	"github.com/tomsalmo/test-oapi-codegen/internal/api"
)

func (h Handler) PostTitUserId(ctx context.Context, request api.PostTitUserIdRequestObject) (api.PostTitUserIdResponseObject, error) {
	return api.PostTitUserId201JSONResponse{}, nil
}

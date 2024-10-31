package titfortat

import (
	"context"

	"github.com/tomsalmo/test-oapi-codegen/cmd/api/api"
)

type TitForTatHandler struct{}

func (h TitForTatHandler) PostTitUserId(ctx context.Context, request api.PostTitUserIdRequestObject) (api.PostTitUserIdResponseObject, error) {
	return api.PostTitUserId201JSONResponse{}, nil
}

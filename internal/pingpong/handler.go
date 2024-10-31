package pingpong

import (
	"context"

	"github.com/tomsalmo/test-oapi-codegen/cmd/api/api"
)

type PingPongHandler struct{}

func (h PingPongHandler) GetPing(ctx context.Context, request api.GetPingRequestObject) (api.GetPingResponseObject, error) {
	return api.GetPing200JSONResponse{}, nil
}

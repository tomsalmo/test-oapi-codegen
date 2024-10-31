package handlers

import "context"

type Handler struct {
	PingPongService interface {
		Ping(ctx context.Context) error
	}
	TitForTatService interface{}
}

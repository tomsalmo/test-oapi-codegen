package pingpong

import "context"

type Service struct{}

func (s Service) Ping(ctx context.Context) error {
	return nil
}

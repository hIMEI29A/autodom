package advisor

import (
	"context"
	"errors"
)

var (
	ErrOrderNotFound   = errors.New("order not found")
	ErrCmdRepository   = errors.New("unable to command repository")
	ErrQueryRepository = errors.New("unable to query repository")
)

type Service interface {
	GetByTitle(ctx context.Context, title string, number int) ([]Solution, error)
}

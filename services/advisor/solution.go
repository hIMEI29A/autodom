package advisor

import (
	"context"
)

type Solution struct {
	Category    string `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Repository interface {
	GetSolutionsByTitle(ctx context.Context, title string, number int) ([]Solution, error)
}

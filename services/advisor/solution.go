package advisor

import (
	"context"
)

// Solution is a basic model
type Solution struct {
	Category    string `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Repository describes the persistence on solution model
type Repository interface {
	GetSolutionsByTitle(ctx context.Context, title string, number int) ([]Solution, error)
}

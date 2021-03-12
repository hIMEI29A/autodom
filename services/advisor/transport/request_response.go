package transport

import (
	"autodom/services/advisor"
)

// GetByTitleRequest stores data of request
type GetByTitleRequest struct {
	Title  string
	Number int
}

// GetByTitleResponse stores data for response
type GetByTitleResponse struct {
	Solutions []advisor.Solution `json:"solutions"`
	Err       error              `json:"error,omitempty"`
}

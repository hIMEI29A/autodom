package transport

import (
	"autodom/services/advisor"
)

type GetByTitleRequest struct {
	Title  string
	Number int
}

type GetByTitleResponse struct {
	Solutions []advisor.Solution `json:"solutions"`
	Err       error              `json:"error,omitempty"`
}

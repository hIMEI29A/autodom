package middleware

import (
	"autodom/services/advisor"
)

// Middleware describes a service middleware.
type Middleware func(service advisor.Service) advisor.Service

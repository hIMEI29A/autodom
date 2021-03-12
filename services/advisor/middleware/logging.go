package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"autodom/services/advisor"
)

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next advisor.Service) advisor.Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   advisor.Service
	logger log.Logger
}

func (mw loggingMiddleware) GetByTitle(ctx context.Context, title string, num int) (solutions []advisor.Solution, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetByTitle", "Title", title, "took", time.Since(begin), "err", err)
	}(time.Now())

	return mw.next.GetByTitle(ctx, title, num)
}

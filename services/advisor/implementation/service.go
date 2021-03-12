package implementation

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"autodom/services/advisor"
)

type service struct {
	repository advisor.Repository
	logger     log.Logger
}

func NewService(rep advisor.Repository, logger log.Logger) advisor.Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) GetByTitle(ctx context.Context, title string, number int) ([]advisor.Solution, error) {
	logger := log.With(s.logger, "method", "GetByTitle")
	solutions, err := s.repository.GetSolutionsByTitle(ctx, title, number)

	if err != nil {
		level.Error(logger).Log("err", err)

		if err == sql.ErrNoRows {
			return solutions, advisor.ErrSolutionNotFound
		}

		return solutions, advisor.ErrQueryRepository
	}

	return solutions, nil
}

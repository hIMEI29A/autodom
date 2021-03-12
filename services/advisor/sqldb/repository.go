package sqldb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/go-sql-driver/mysql"

	"autodom/services/advisor"
)

var (
	ErrRepository = errors.New("unable to handle request")
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

// New returns a concrete repository backed by MySQL
func New(db *sql.DB, logger log.Logger) (advisor.Repository, error) {
	return &repository{
		db:     db,
		logger: log.With(logger, "rep", "sqldb"),
	}, nil
}

// GetSolutionsByTitle creates full test search request to BD. First argument is a
// problem's title to search, second is a wanted number of results.
// Returns slice of advisor.Solution and error/
func (repo *repository) GetSolutionsByTitle(ctx context.Context, title string, num int) ([]advisor.Solution, error) {
	rows, err := repo.db.Query("SELECT * FROM cases WHERE MATCH (title) AGAINST (? IN NATURAL LANGUAGE MODE) LIMIT ?", title, num)

	if err != nil {
		level.Error(repo.logger).Log("err", err.Error())
		return nil, err
	}

	defer rows.Close()

	entities := []advisor.Solution{}

	for rows.Next() {
		e := advisor.Solution{}
		err := rows.Scan(&e.Category, &e.Title, &e.Description)

		if err != nil {
			level.Error(repo.logger).Log("err", err.Error())
			continue
		}

		entities = append(entities, e)
	}

	return entities, err
}

// Close closes db connection.
func (repo *repository) Close() error {
	return repo.db.Close()
}

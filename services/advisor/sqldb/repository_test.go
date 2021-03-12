package sqldb

import (
	"autodom/services/advisor"
	"context"
	"database/sql"
	"reflect"
	"testing"

	//"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
)

func TestNew(t *testing.T) {
	var (
		sqlDB          *sql.DB
		l              log.Logger
		repositoryType = reflect.TypeOf(new(advisor.Repository)).Elem()
	)
	type args struct {
		db     *sql.DB
		logger log.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"test1", args{sqlDB, l}, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.db, tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if reflect.TypeOf(got).Implements(repositoryType) != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetSolutionsByTitle(t *testing.T) {

	type fields struct {
		db     *sql.DB
		logger log.Logger
	}
	type args struct {
		ctx   context.Context
		title string
		num   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []advisor.Solution
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := repo.GetSolutionsByTitle(tt.args.ctx, tt.args.title, tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetSolutionsByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetSolutionsByTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Close(t *testing.T) {
	type fields struct {
		db     *sql.DB
		logger log.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			if err := repo.Close(); (err != nil) != tt.wantErr {
				t.Errorf("repository.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

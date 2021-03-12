package implementation

import (
	"autodom/services/advisor"
	"context"
	"reflect"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestNewService(t *testing.T) {
	type args struct {
		rep    advisor.Repository
		logger log.Logger
	}
	tests := []struct {
		name string
		args args
		want advisor.Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.rep, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetByTitle(t *testing.T) {
	type fields struct {
		repository advisor.Repository
		logger     log.Logger
	}
	type args struct {
		ctx    context.Context
		title  string
		number int
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
			s := &service{
				repository: tt.fields.repository,
				logger:     tt.fields.logger,
			}
			got, err := s.GetByTitle(tt.args.ctx, tt.args.title, tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetByTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

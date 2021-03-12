package middleware

import (
	"autodom/services/advisor"
	"context"
	"reflect"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestLoggingMiddleware(t *testing.T) {
	type args struct {
		logger log.Logger
	}
	tests := []struct {
		name string
		args args
		want Middleware
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoggingMiddleware(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggingMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loggingMiddleware_GetByTitle(t *testing.T) {
	type fields struct {
		next   advisor.Service
		logger log.Logger
	}
	type args struct {
		ctx   context.Context
		title string
		num   int
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantSolutions []advisor.Solution
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := loggingMiddleware{
				next:   tt.fields.next,
				logger: tt.fields.logger,
			}
			gotSolutions, err := mw.GetByTitle(tt.args.ctx, tt.args.title, tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("loggingMiddleware.GetByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSolutions, tt.wantSolutions) {
				t.Errorf("loggingMiddleware.GetByTitle() = %v, want %v", gotSolutions, tt.wantSolutions)
			}
		})
	}
}

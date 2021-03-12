package http

import (
	"autodom/services/advisor/transport"
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

func TestNewService(t *testing.T) {
	type args struct {
		svcEndpoints transport.Endpoints
		options      []kithttp.ServerOption
		logger       log.Logger
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.svcEndpoints, tt.args.options, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeGetByTitleRequest(t *testing.T) {
	type args struct {
		in0 context.Context
		r   *http.Request
	}
	tests := []struct {
		name        string
		args        args
		wantRequest interface{}
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRequest, err := decodeGetByTitleRequest(tt.args.in0, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeGetByTitleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRequest, tt.wantRequest) {
				t.Errorf("decodeGetByTitleRequest() = %v, want %v", gotRequest, tt.wantRequest)
			}
		})
	}
}

func Test_encodeResponse(t *testing.T) {
	type args struct {
		ctx      context.Context
		w        http.ResponseWriter
		response interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := encodeResponse(tt.args.ctx, tt.args.w, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("encodeResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeErrorResponse(t *testing.T) {
	type args struct {
		in0 context.Context
		err error
		w   http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodeErrorResponse(tt.args.in0, tt.args.err, tt.args.w)
		})
	}
}

func Test_codeFrom(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := codeFrom(tt.args.err); got != tt.want {
				t.Errorf("codeFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeIncomingRequest(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := decodeIncomingRequest(tt.args.r)
			if got != tt.want {
				t.Errorf("decodeIncomingRequest() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("decodeIncomingRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

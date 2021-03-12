package transport

import (
	"autodom/services/advisor"
	"reflect"
	"testing"

	"github.com/go-kit/kit/endpoint"
)

func TestMakeEndpoints(t *testing.T) {
	type args struct {
		s advisor.Service
	}
	tests := []struct {
		name string
		args args
		want Endpoints
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeEndpoints(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeEndpoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeGetByTitleEndpoint(t *testing.T) {
	type args struct {
		s advisor.Service
	}
	tests := []struct {
		name string
		args args
		want endpoint.Endpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGetByTitleEndpoint(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeGetByTitleEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

package transport

import (
	"context"

	"autodom/services/advisor"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetByTitle endpoint.Endpoint
}

func MakeEndpoints(s advisor.Service) Endpoints {
	return Endpoints{
		GetByTitle: makeGetByTitleEndpoint(s),
	}
}

func makeGetByTitleEndpoint(s advisor.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByTitleRequest)
		solutionRes, err := s.GetByTitle(ctx, req.Title, req.Number)
		return GetByTitleResponse{Solutions: solutionRes, Err: err}, nil
	}
}

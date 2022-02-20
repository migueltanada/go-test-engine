package questionaire

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateSet endpoint.Endpoint
	GetSet    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateSet: makeCreateSetEndpoint(s),
		GetSet:    makeGetSetEndpoint(s),
	}
}

func makeCreateSetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateSetRequest)
		ok, err := s.CreateSet(ctx, req.Name, req.Category)
		return CreateSetResponse{Ok: ok}, err
	}
}

func makeGetSetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSetRequest)
		name, err := s.GetSet(ctx, req.Id)

		return GetSetResponse{
			Name: name,
		}, err
	}
}

package questionaire

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateSet endpoint.Endpoint
	GetSet    endpoint.Endpoint
	CreateQna endpoint.Endpoint
	GetQna    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateSet: makeCreateSetEndpoint(s),
		GetSet:    makeGetSetEndpoint(s),
		CreateQna: makeCreateQnaEndpoint(s),
		GetQna:    makeGetQnaEndpoint(s),
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
		name, err := s.GetSet(ctx, req.SetID)

		return GetSetResponse{
			Name: name,
		}, err
	}
}

func makeCreateQnaEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateQnaRequest)
		ok, err := s.CreateQna(ctx, req.SetID, req.Question, req.Choice)
		return CreateSetResponse{Ok: ok}, err
	}
}

func makeGetQnaEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetQnaRequest)
		question, err := s.GetQna(ctx, req.QnaID)

		return GetQnaResponse{
			Question: question,
		}, err
	}
}

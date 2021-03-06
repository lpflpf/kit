package endpoints

import (
	"context"

	"github.com/lpflpf/kit/cmd/kitgen/testdata/anonfields/default/service"
	"github.com/lpflpf/kit/endpoint"
)

type FooRequest struct {
	I int
	S string
}
type FooResponse struct {
	I   int
	Err error
}

func MakeFooEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		i, err := s.Foo(ctx, req.I, req.S)
		return FooResponse{I: i, Err: err}, nil
	}
}

type Endpoints struct {
	Foo endpoint.Endpoint
}

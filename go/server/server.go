package server

import (
	"context"
	"github.com/ngyewch/twirp-playground/rpc"
	"github.com/ngyewch/twirp-playground/rpc2"
	"strings"
)

type Server struct {
}

func (s *Server) Add(ctx context.Context, request *rpc.AddRequest) (*rpc.AddResponse, error) {
	return &rpc.AddResponse{
		Value: request.A + request.B,
	}, nil
}

func (s *Server) DoSomething(ctx context.Context, request *rpc2.DoSomethingRequest) (*rpc2.DoSomethingResponse, error) {
	return &rpc2.DoSomethingResponse{
		Text: strings.ToUpper(request.Text),
	}, nil
}

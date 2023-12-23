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

func (s *Server) ToUpper(ctx context.Context, request *rpc2.ToUpperRequest) (*rpc2.ToUpperResponse, error) {
	return &rpc2.ToUpperResponse{
		Text: strings.ToUpper(request.Text),
	}, nil
}

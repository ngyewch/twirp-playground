package server

import (
	"context"
	"github.com/ngyewch/twirp-playground/rpc"
)

type Server struct {
}

func (s *Server) Add(ctx context.Context, request *rpc.AddRequest) (*rpc.AddResponse, error) {
	return &rpc.AddResponse{
		Value: request.A + request.B,
	}, nil
}

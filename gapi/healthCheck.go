package gapi

import (
	"context"
	"economic-school/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) HealthCheck(ctx context.Context, req *emptypb.Empty) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: "SERVING"}, nil
}

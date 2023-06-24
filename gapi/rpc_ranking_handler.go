package gapi

import (
	"context"
	"economic-school/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) GetGlobalRanking(ctx context.Context, req *emptypb.Empty) (*pb.GetGlobalRankingResponse, error) {

	globalRanking, err := server.store.GetGlobalRanking(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get global ranking: %s", err)
	}
	responseRanking := make([]*pb.GlobalRanking, len(globalRanking))
	for i, ranking := range globalRanking {
		responseRanking[i] = convertGlobalRanking(ranking)
	}

	return &pb.GetGlobalRankingResponse{Rankings: responseRanking}, nil
}

func (server *Server) GetWeeklyGlobalRanking(ctx context.Context, req *emptypb.Empty) (*pb.GetWeeklyGlobalRankingResponse, error) {

	weeklyGlobalRanking, err := server.store.GetWeeklyGlobalRanking(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get weekly global ranking: %s", err)
	}
	responseRanking := make([]*pb.WeeklyGlobalRanking, len(weeklyGlobalRanking))
	for i, ranking := range weeklyGlobalRanking {
		responseRanking[i] = convertWeeklyGlobalRanking(ranking)
	}

	return &pb.GetWeeklyGlobalRankingResponse{Rankings: responseRanking}, nil
}

func (server *Server) GetUniversityRanking(ctx context.Context, req *emptypb.Empty) (*pb.GetUniversityRankingResponse, error) {

	universityRanking, err := server.store.GetUniversityRanking(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get university ranking: %s", err)
	}
	responseRanking := make([]*pb.UniversityRanking, len(universityRanking))
	for i, ranking := range universityRanking {
		responseRanking[i] = convertUniversityRanking(ranking)
	}

	return &pb.GetUniversityRankingResponse{Rankings: responseRanking}, nil
}

func (server *Server) GetWeeklyUniversityRanking(ctx context.Context, req *emptypb.Empty) (*pb.GetWeeklyUniversityRankingResponse, error) {

	weeklyUniversityRanking, err := server.store.GetWeeklyUniversityRanking(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get weekly university ranking: %s", err)
	}
	responseRanking := make([]*pb.WeeklyUniversityRanking, len(weeklyUniversityRanking))
	for i, ranking := range weeklyUniversityRanking {
		responseRanking[i] = convertWeeklyUniversityRanking(ranking)
	}

	return &pb.GetWeeklyUniversityRankingResponse{Rankings: responseRanking}, nil
}

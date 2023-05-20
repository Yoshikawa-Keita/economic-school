package gapi

import (
	db "economic-school/db/sqlc"
	"economic-school/pb"
	"economic-school/service"
	"economic-school/token"
	"economic-school/util"
	"economic-school/worker"
	"fmt"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedEconomicSchoolServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
	s3Uploader      *service.S3Uploader
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	s3Uploader, err := service.NewS3Uploader(config)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
		s3Uploader:      s3Uploader,
	}

	return server, nil
}

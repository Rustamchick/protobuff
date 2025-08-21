package auth

import (
	"context"

	grpcpetv1 "github.com/Rustamchick/protobuff/gen/go/pet"
	"google.golang.org/grpc"
)

type ServerApi struct {
	grpcpetv1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	grpcpetv1.RegisterAuthServer(gRPC, &ServerApi{})
}

func (s *ServerApi) Login(ctx context.Context, req *grpcpetv1.LoginRequest) (*grpcpetv1.LoginResponse, error) {
	return &grpcpetv1.LoginResponse{
		Token: req.GetEmail(),
	}, nil
}

func (s *ServerApi) Register(ctx context.Context, req *grpcpetv1.RegisterRequest) (*grpcpetv1.RegisterResponse, error) {
	return &grpcpetv1.RegisterResponse{}, nil
}

func (s *ServerApi) IsAdmin(ctx context.Context, req *grpcpetv1.IsAdminRequest) (*grpcpetv1.IsAdminResponse, error) {
	return &grpcpetv1.IsAdminResponse{}, nil
}

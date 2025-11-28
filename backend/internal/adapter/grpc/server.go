package grpc_adapter

import (
	"Calculator/backend/internal/core/service"
	"Calculator/pkg/pb"
	"context"
)

type GrpcServer struct {
	pb.UnimplementedCalculatorServer
	Service service.CalculatorService
}

func NewGrpcServer(svc service.CalculatorService) *GrpcServer {
	return &GrpcServer{Service: svc}
}

func (s *GrpcServer) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	result := s.Service.Add(req.Num1, req.Num2)
	return &pb.Response{Num3: result}, nil
}

func (s *GrpcServer) Sub(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	result := s.Service.Sub(req.Num1, req.Num2)
	return &pb.Response{Num3: result}, nil
}

func (s *GrpcServer) Mul(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	result := s.Service.Mul(req.Num1, req.Num2)
	return &pb.Response{Num3: result}, nil
}
func (s *GrpcServer) Div(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	result, err := s.Service.Div(req.Num1, req.Num2)
	return &pb.Response{Num3: result}, err
}
func (s *GrpcServer) Mod(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	result, err := s.Service.Mod(req.Num1, req.Num2)
	return &pb.Response{Num3: result}, err
}

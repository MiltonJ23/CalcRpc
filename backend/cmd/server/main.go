package main

import (
	grpc_adapter "Calculator/backend/internal/adapter/grpc"
	http_adapter "Calculator/backend/internal/adapter/http"
	"Calculator/backend/internal/core/service"
	"Calculator/pkg/pb"
	"log"
	"net"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// 1. Initialize Core Service
	coreService := service.NewCalculatorService()

	// 2. Start gRPC Server (in a goroutine)
	go func() {
		lis, err := net.Listen("tcp", ":50051") // [cite: 119]
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterCalculatorServer(grpcServer, grpc_adapter.NewGrpcServer(coreService))
		log.Println("gRPC Server listening on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// 3. Start Gin Server (Main thread)
	r := gin.Default()
	// Add CORS for Next.js
	r.Use(cors.Default())

	httpHandler := http_adapter.NewHttpHandler(coreService)
	http_adapter.RegisterRoutes(r, httpHandler)

	log.Println("HTTP Gin Server listening on port 7080")
	r.Run(":7080")
}

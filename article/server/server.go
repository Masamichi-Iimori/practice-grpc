package main

import (
	"log"
	"net"

	"github.com/Masamichi-Iimori/practice-grpc/article/pb"
	"github.com/Masamichi-Iimori/practice-grpc/article/repository"
	"github.com/Masamichi-Iimori/practice-grpc/article/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()

	repository, err := repository.NewsqliteRepo()
	if err != nil {
		log.Fatalf("Failed to create sqlite repository: %v\n", err)
	}

	service := service.NewService(repository)

	if err != nil {
		log.Fatalf("Failed to create sqlite repository: %v\n", err)
	}

	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, service)

	log.Println("Listening on port 50051...")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

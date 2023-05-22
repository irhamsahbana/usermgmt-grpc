package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/irhamsahbana/usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

func main() {
	port := ":50051"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateUser(ctx context.Context, in *pb.NewUSer) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())

	var userID int32 = int32(rand.Intn(100))

	return &pb.User{Id: userID, Name: in.GetName(), Age: in.GetAge()}, nil
}

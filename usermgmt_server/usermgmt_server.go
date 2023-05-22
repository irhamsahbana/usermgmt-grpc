package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/irhamsahbana/usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	PORT = ":50051"
)

func main() {
	var user_management_server = NewUserManagementServer()
	if err := user_management_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		user_list: &pb.UserList{},
	}
}

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	user_list *pb.UserList
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("Starting gRPC listener on port " + PORT)
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *UserManagementServer) CreateUser(ctx context.Context, in *pb.NewUSer) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())

	var userID int32 = int32(rand.Intn(100))

	created_user := &pb.User{Id: userID, Name: in.GetName(), Age: in.GetAge()}
	s.user_list.Users = append(s.user_list.Users, created_user)

	return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	return s.user_list, nil
}

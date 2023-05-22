package main

import (
	"context"
	"log"
	"time"

	pb "github.com/irhamsahbana/usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newUsers = make(map[string]int32)
	newUsers["irham"] = 22
	newUsers["sahbana"] = 23

	for name, age := range newUsers {
		result, err := c.CreateUser(ctx, &pb.NewUSer{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not greet : %v", err)
		}

		log.Printf(`User Details:
		Name : %s
		Age : %d
		Id : %d`, result.GetName(), result.GetAge(), result.GetId())
	}
}

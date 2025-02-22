package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"authentication/helper"
	"authentication/model"
	userpb "authentication/userpb"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedAllTheServicesServer
}

func (s *server) Login(ctx context.Context, req *userpb.LoginDataSend) (*userpb.User, error) {
	fmt.Printf("Received GetUser request for UserID: %s\n", req.UserId)
	var c model.Login_system
	c.Username = req.UserId
	c.Password = req.Pass

	if c.Username == "User2" && c.Password == "password2" {

		return &userpb.User{
			Name:       req.UserId,
			Age:        23,
			Contact:    "3287476235645",
			Manager:    true,
			Department: "admin",
			Status:     true,
		}, nil

	}

	a, b := helper.SearchData(c)
	fmt.Println(a)
	return &userpb.User{
		Name:       a.Name,
		Age:        a.Age,
		Contact:    a.Contact,
		Manager:    a.Manager,
		Department: a.Department,
		Status:     b,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:3001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterAllTheServicesServer(grpcServer, &server{})

	fmt.Println("gRPC Server is running on port 3001...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

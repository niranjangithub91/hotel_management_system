package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"user_management/helper"
	"user_management/model"
	userpb "user_management/userpb"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedUsermanagementServiceServer
}

func (s *server) Add_Users(ctx context.Context, req *userpb.Send_User_Data) (*userpb.Status, error) {
	var t model.User
	t.Name = req.Name
	t.Password = req.Password
	t.Age = req.Age
	t.Email = req.Gmail
	t.Contact = req.Contact
	t.Manager = req.Manager
	t.Department = req.Department
	helper.Add_Users(t)
	return &userpb.Status{
		Status: true,
	}, nil
}
func (s *server) DeleteUser(ctx context.Context, req *userpb.SendDeleteUserData) (*userpb.Status, error) {
	var t model.Person
	t.Name = req.Name
	helper.Delete_users(t)
	return &userpb.Status{
		Status: true,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":3002")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUsermanagementServiceServer(grpcServer, &server{})

	fmt.Println("gRPC Server is running on port 3002...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

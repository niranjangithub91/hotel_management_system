package main

import (
	"context"
	"customer/helper"
	"customer/model"
	"customer/userpb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedCustomerManagementServer
}

func (s *server) Add_Customer(ctx context.Context, req *userpb.Send_CustomerDetails) (*userpb.Status, error) {
	var f model.Customer
	f.Name = req.Name
	f.Email = req.Email
	f.Contact = req.Contact
	helper.Add_Customers_helper(f)
	return &userpb.Status{
		Status: true,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:3004")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterCustomerManagementServer(grpcServer, &server{})

	fmt.Println("gRPC Server is running on port 3004...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

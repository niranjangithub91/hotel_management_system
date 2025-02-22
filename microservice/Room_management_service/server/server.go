package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"room_management/helper"
	"room_management/model"
	"room_management/userpb"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedRoommanagementServiceServer
}

func (s *server) AddRooms(ctx context.Context, req *userpb.SendRoomDetails) (*userpb.Status, error) {

	var a model.Room

	a.Room_number = req.Roomnumber
	a.Price = req.Price
	a.OccupencyStatus = req.Occupencystatus
	a.Features = req.Features
	helper.AddRooms_helper(a)
	return &userpb.Status{
		Status: true,
	}, nil
}

func (s *server) DeleteRoom(ctx context.Context, req *userpb.SendDeleteRoomDetail) (*userpb.Status, error) {
	helper.DeleteRoom_helper(*req)
	return &userpb.Status{
		Status: true,
	}, nil
}

func (s *server) UpdateFeatures(ctx context.Context, req *userpb.SendFeatureUpdate) (*userpb.Status, error) {
	var a model.Room_feature_update
	a.Room_numer = req.Roomnumber
	a.Features = req.Features
	helper.UpdateFeatures_helper(a)
	return &userpb.Status{
		Status: true,
	}, nil
}

func (s *server) UpdatePrice(ctx context.Context, req *userpb.SendPriceUpdate) (*userpb.Status, error) {
	var a model.Room_price_update
	a.Room_number = req.Roomnumber
	a.Price = req.Price
	helper.UpdatePrice_helper(a)
	return &userpb.Status{
		Status: true,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:3003")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterRoommanagementServiceServer(grpcServer, &server{})

	fmt.Println("gRPC Server is running on port 3003...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

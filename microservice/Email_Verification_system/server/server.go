package main

import (
	"context"
	"email/email"
	"email/userpb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedEmailVerificationServer
}

func (s *server) Send_OTP(ctx context.Context, req *userpb.Send_CustomerEmail) (*userpb.Send_OTPBack, error) {
	x := email.Give_the_mail(req.Email)
	fmt.Println(x)
	return &userpb.Send_OTPBack{
		Otp: x,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:3005")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterEmailVerificationServer(grpcServer, &server{})

	fmt.Println("gRPC Server is running on port 3005...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

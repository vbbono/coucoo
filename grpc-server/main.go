package main

import (
	"context"
	"github.com/vbbono/coucoo/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedCouponServer
}

func (s *server) RequestCoupon(ctx context.Context, in *pb.User) (*pb.Success, error) {
	log.Printf("Received: %v", in.GetID())
	return &pb.Success{OK: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCouponServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

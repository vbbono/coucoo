package main

import (
	"context"
	"github.com/vbbono/coucoo/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCouponClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RequestCoupon(ctx, &pb.User{ID: 123})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// push userid to kafka queue
	// type of client and server should be changed to stream.
	log.Printf("Check: %v", r.GetOK())
}

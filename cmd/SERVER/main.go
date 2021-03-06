package main

import (
	"fmt"
	"log"
	"net"

	"github.com/shashijangra22/bootcamp-project/pkg/ServerCore"

	"github.com/shashijangra22/bootcamp-project/pkg/customer"
	"github.com/shashijangra22/bootcamp-project/pkg/order"
	"github.com/shashijangra22/bootcamp-project/pkg/restaurant"
	"google.golang.org/grpc"
)

func main() {
	// fire the gRPC Server
	fmt.Println("Hello from grpc server.")
	lis, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		log.Fatalf("Sorry failed to load server %v:", err)
	}
	s := grpc.NewServer()

	customer.RegisterCustomerServiceServer(s, &ServerCore.Server{})
	order.RegisterOrderServiceServer(s, &ServerCore.Server{})
	restaurant.RegisterRestaurantServiceServer(s, &ServerCore.Server{})

	if s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-bi-directional-streaming/hello/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to correct: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	doHelloEveryone(c)
}

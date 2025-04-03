package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFileServiceClient(conn)
	callListFiles(c)
}

func callListFiles(c pb.FileServiceClient) {
	res, err := c.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalf("error while calling ListFiles: %v", err)
	}
	fmt.Printf("Response from ListFiles: %v", res.GetFilenames())
}

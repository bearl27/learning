package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"io"
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
	// callListFiles(c)
	callDownload(c)
}

// type FileServiceClient interface {
// 	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error)
// 	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadResponse], error)
// }

func callListFiles(c pb.FileServiceClient) {
	res, err := c.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalf("error while calling ListFiles: %v", err)
	}
	fmt.Printf("Response from ListFiles: %v", res.GetFilenames())
}

func callDownload(c pb.FileServiceClient) {
	req := &pb.DownloadRequest{Filename: "name.txt"}
	stream, err := c.Download(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Download: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from Download: %v", res.GetData())
		log.Printf("Response from Download: %v", string(res.GetData()))
	}

}

package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"io"
	"log"
	"os"
	"time"

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
	//callDownload(c)
	// callUpload(c)
	callUploadAndNotifyProgress(c)
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
	req := &pb.DownloadRequest{Filename: "sports.txt"}
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

// Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadRequest, UploadResponse], error)

func callUpload(c pb.FileServiceClient) {
	filename := "sports.txt"
	path := "/Users/Nagahari.Kai/Github/learning/grpc-lesson/storage/" + filename

	data, err := os.Open(path)
	if err != nil {
		log.Fatalf("error while opening file: %v", err)
	}
	defer data.Close()

	stream, err := c.Upload(context.Background())
	if err != nil {
		log.Fatalf("error while calling Upload: %v", err)
	}
	buf := make([]byte, 5)
	for {
		n, err := data.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading file: %v", err)
		}
		req := &pb.UploadRequest{Data: buf[:n]}
		sendErr := stream.Send(req)
		if sendErr != nil {
			log.Fatalf("error while sending request: %v", sendErr)
		}
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v", err)
	}

	log.Printf("Response from Upload: %v", res.GetSize())
}

// UploadAndNotifyProgress(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadAndNotifyProgressRequest, UploadAndNotifyProgressResponse], error)

func callUploadAndNotifyProgress(c pb.FileServiceClient) {
	filename := "sports.txt"
	path := "/Users/Nagahari.Kai/Github/learning/grpc-lesson/storage/" + filename

	data, err := os.Open(path)
	if err != nil {
		log.Fatalf("error while opening file: %v", err)
	}
	defer data.Close()

	stream, err := c.UploadAndNotifyProgress(context.Background())
	if err != nil {
		log.Fatalf("error while calling UploadAndNotifyProgress: %v", err)
	}

	// request
	buf := make([]byte, 5)
	go func() {
		for {
			n, err := data.Read(buf)
			if n == 0 || err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while reading file: %v", err)
			}
			req := &pb.UploadAndNotifyProgressRequest{Data: buf[:n]}
			sendErr := stream.Send(req)
			if sendErr != nil {
				log.Fatalf("error while sending request: %v", sendErr)
			}
			time.Sleep(1 * time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("error while closing stream: %v", err)
		}
	}()

	// response
	ch := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving response: %v", err)
			}

			log.Printf("Response from UploadAndNotifyProgress: %v", res.GetMsg())
		}
		close(ch)
	}()
	<-ch
}

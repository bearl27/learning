package main

import (
	"bytes"
	"context"
	"fmt"
	"grpc/pb"
	"io"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

// type FileServiceServer interface {
// 	ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
// 	Download(*DownloadRequest, grpc.ServerStreamingServer[DownloadResponse]) error
// 	mustEmbedUnimplementedFileServiceServer()
// }

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println("ListFiles was invoked")

	dir := "/Users/Nagahari.Kai/Github/learning/grpc-lesson/storage"

	paths, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, path := range paths {
		if !path.IsDir() {
			filenames = append(filenames, path.Name())
		}
	}

	res := &pb.ListFilesResponse{
		Filenames: filenames,
	}

	return res, nil
}

func (*server) Download(req *pb.DownloadRequest, stream pb.FileService_DownloadServer) error {
	fmt.Println("Download was invoked")

	filename := req.GetFilename()
	path := "/Users/Nagahari.Kai/Github/learning/grpc-lesson/storage/" + filename

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 5)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		res := &pb.DownloadResponse{Data: buf[:n]}
		sendErr := stream.Send(res)
		if sendErr != nil {
			return sendErr
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

// Upload(grpc.ClientStreamingServer[UploadRequest, UploadResponse]) error

func (*server) Upload(stream pb.FileService_UploadServer) error {
	fmt.Println("Upload was invoked")
	var buf bytes.Buffer
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := &pb.UploadResponse{Size: int32(buf.Len())}
			return stream.SendAndClose(res)
		}
		if err != nil {
			return err
		}

		data := req.GetData()
		log.Printf("received data(bytes): %v", data)
		log.Printf("received data(string): %v", string(data))
		buf.Write(data)
	}
}

// UploadAndNotifyProgress(grpc.ClientStreamingServer[UploadAndNotifyProgressRequest, UploadAndNotifyProgressResponse]) error
func (*server) UploadAndNotifyProgress(stream pb.FileService_UploadAndNotifyProgressServer) error {
	fmt.Println("UploadAndNotifyProgress was invoked")

	size := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		data := req.GetData()
		log.Printf("received data: %v", data)
		log.Printf("received data(string): %v", string(data))
		size += len(data)

		res := &pb.UploadAndNotifyProgressResponse{
			Msg: fmt.Sprintf("received %v bytes", size),
		}
		err = stream.Send(res)
		if err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

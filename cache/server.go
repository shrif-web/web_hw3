package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/status"
	"log"
	"net"

	pb "cache/grpc"
	"cache/lru"
	"google.golang.org/grpc"
)

var cache = lru.New(10)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCacheHandlerServer
}

func (s *server) GetKey(_ context.Context, in *pb.GetKeyRequest) (*pb.GetKeyReply, error) {
	log.Printf("Get Key: %s", in.Key)
	if value, exists := cache.Get(in.Key); exists {
		return &pb.GetKeyReply{Value: value}, nil
	} else {
		return &pb.GetKeyReply{Value: value}, status.Errorf(404, "Key not found.")
	}
}

func (s *server) SetKey(_ context.Context, in *pb.SetKeyRequest) (*pb.SetKeyReply, error) {
	log.Printf("Set Key: %s -> %s", in.Key, in.Value)
	cache.Put(in.Key, in.Value)
	return &pb.SetKeyReply{}, nil
}

func (s *server) Clear(_ context.Context, _ *pb.ClearRequest) (*pb.ClearReply, error) {
	log.Printf("Clear Cache")
	cache.Clear()
	return &pb.ClearReply{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCacheHandlerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
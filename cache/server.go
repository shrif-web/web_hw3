package main

import (
	pb "cache/grpc"
	"cache/lru"
	"cache/utils"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strconv"
)

var capacity, _ = strconv.Atoi(utils.GetEnv("cache_capacity", "100"))
var cache = lru.New(capacity)

var (
	port = flag.String(
		"port",
		utils.GetEnv("grpc_port", "50051"),
		"The server port")

	crt        = "certs/cert.pem"
	key        = "certs/key.pem"
	sslEnabled = utils.GetEnv("SSL_ENABLE", "true")
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

	if len(in.Key) > 64 {
		return &pb.SetKeyReply{}, status.Error(400, "key should be less than 64 character.")
	}
	if len(in.Value) > 2048 {
		return &pb.SetKeyReply{}, status.Error(400, "Value should be less than 1024 character.")
	}

	cache.Put(in.Key, in.Value)
	return &pb.SetKeyReply{}, nil
}

func (s *server) Clear(_ context.Context, _ *pb.ClearRequest) (*pb.ClearReply, error) {
	log.Printf("Clear Cache")
	cache.Clear()
	return &pb.ClearReply{}, nil
}

func (s *server) Remove(_ context.Context, in *pb.RemoveKeyRequest) (*pb.RemoveKeyReply, error) {
	log.Printf("Remove Key: %s", in.Key)
	cache.Remove(in.Key)
	return &pb.RemoveKeyReply{}, nil
}

func initServer(transportCredentials credentials.TransportCredentials) *grpc.Server {
	log.Printf("ssl: %s", sslEnabled)
	if sslEnabled == "true" {
		return grpc.NewServer(grpc.Creds(transportCredentials))
	} else {
		return grpc.NewServer()
	}
}

func main() {
	transportCredentials, _ := credentials.NewServerTLSFromFile(crt, key)
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := initServer(transportCredentials)
	pb.RegisterCacheHandlerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

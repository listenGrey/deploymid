package main

import (
	"context"
	proto "github.com/listenGrey/lucianagRpcPKG/ask"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedRequestServer
}

func (s server) Request(ctx context.Context, re *proto.Prompt) (*proto.Response, error) {
	log.Printf("Received: %s\n", re.GetPrompt())
	return &proto.Response{}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	proto.RegisterRequestServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

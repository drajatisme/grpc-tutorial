package main

import (
	"context"
	"net"

	"example.com/grpc-tutorial/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedMathServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterMathServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) Add(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	a, b := req.GetA(), req.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	a, b := req.GetA(), req.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}

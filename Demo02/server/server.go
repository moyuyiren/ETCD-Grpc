package main

import (
	"Etcd_Service_register/Demo02/server/proto"
	"context"
)

type server struct {
}

func (s *server) Hello(ctx context.Context, req *rpc.Empty) (*rpc.HelloResponse, error) {
	return &rpc.HelloResponse{Hello: "Hello"}, nil
}
func (s *server) Register(ctx context.Context, req *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	return &rpc.RegisterResponse{Uid: "764952492"}, nil
}

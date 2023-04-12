package main

import (
	"Etcd_Service_register/Demo03_Etcd_Balance/balance"
	"Etcd_Service_register/Demo03_Etcd_Balance/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	Address = "loaclhost:8000"
	Network = "tcp"
	SerName = "yiren_service"
)

var EtcdPoints = []string{"loaclhost:2379"}

func main() {
	//启动监听
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	//新建Grpc实例并注册
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleServer(grpcServer, new(service))
	//服务注册到ETCD
	ser, err := balance.NewServiceRegister(EtcdPoints, SerName+"/"+Address, "1", 5)
	if err != nil {
		log.Fatalf("register service err: %v", err)
	}
	defer ser.Close()
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

// 实现服务
type service struct{}

func (s *service) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {

	return nil, nil
}

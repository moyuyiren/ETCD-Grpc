package main

import (
	rpc "Etcd_Service_register/Demo02/server/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const netAddr = "127.0.0.1:8088"

func main() {
	err := serverRegister(netAddr)
	if err != nil {
		panic(err)
	}

	conn, err := net.Listen("tcp", netAddr)
	if err != nil {
		log.Fatal("端口监听启动失败")
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor()))
	rpc.RegisterServerServer(grpcServer, new(server))
	grpcServer.Serve(conn)

	//优雅关闭
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		serverUnRegister(netAddr)
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()

}

func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Printf("call %s\n", info.FullMethod)
		resp, err = handler(ctx, req)
		return resp, err
	}
}

package main

import (
	rpc "Etcd_Service_register/Demo02_Etcd_RegServer/client/rpc"
	"context"
	"fmt"
	etcd "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const etcdUrl = "http://127.0.0.1:2379"
const serverName = "/yiren/grpcserver"

func main() {
	etcdClient, err := etcd.NewFromURL(etcdUrl)
	defer etcdClient.Close()
	if err != nil {
		return
	}
	etcdResolver, err := resolver.NewBuilder(etcdClient)
	conn, err := grpc.Dial(fmt.Sprintf("etcd:///%s", serverName),
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
	)
	serverClient := rpc.NewServerClient(conn)
	for {
		helloRespone, err := serverClient.Hello(context.Background(), &rpc.Empty{})
		if err != nil {
			fmt.Printf("err: %v", err)
			return
		}

		log.Println(helloRespone.Hello, err)
		time.Sleep(time.Second)
	}
	return

}

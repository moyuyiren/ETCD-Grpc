package main

import (
	"context"
	"fmt"
	etcd "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
)

const etcdUrl = "http://127.0.0.1:2379"
const serverName = "/yiren/grpcserver"
const ttl = 10

var etcdClient *etcd.Client
var etcdClientUn *etcd.Client

func serverRegister(addr string) error {
	log.Println("etcdRegister" + addr)
	etcdClient, err := etcd.NewFromURL(etcdUrl)
	defer etcdClient.Close()
	if err != nil {
		log.Println("etcdClient初始化失败")
		return err
	}
	em, err := endpoints.NewManager(etcdClient, serverName)
	if err != nil {
		log.Println("Manager 初始化失败")
		return err
	}
	//创建租约
	lease, err := etcdClient.Grant(context.TODO(), ttl)
	if err != nil {
		log.Println("租约初始化失败")
		return err
	}
	err = em.AddEndpoint(context.TODO(), fmt.Sprintf("%s/%s", serverName, addr), endpoints.Endpoint{Addr: addr}, etcd.WithLease(lease.ID))
	if err != nil {
		return err
	}
	//保活
	alive, err := etcdClient.KeepAlive(context.TODO(), lease.ID)
	if err != nil {
		return err
	}

	go func() {
		for {
			<-alive
			fmt.Println("etcd server keep alive")
		}
	}()

	return nil
}

func serverUnRegister(addr string) error {
	log.Printf("etcdUnRegister %s\b", addr)
	etcdClientUn, err := etcd.NewFromURL(etcdUrl)

	defer etcdClient.Close()
	if err != nil {
		log.Println("etcdClient初始化失败")
		return err
	}
	if etcdClientUn != nil {
		em, err := endpoints.NewManager(etcdClientUn, serverName)
		if err != nil {
			return err
		}
		err = em.DeleteEndpoint(context.TODO(), fmt.Sprintf("%s/%s", serverName, addr))
		if err != nil {
			return err
		}
		return err
	}

	return nil

}

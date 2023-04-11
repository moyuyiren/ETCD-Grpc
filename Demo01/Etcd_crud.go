package main

import (
	"context"
	"fmt"
	etcd "go.etcd.io/etcd/client/v3"
	"log"
)

const etcdctl = "http://localhost:2379"
const serviceName = "/yiren/server"
const ttl = 10

var etcdClient *etcd.Client

func main() {
	//TestCrud()
	TestetcdLeave(serviceName)

}

func TestCrud() {
	if err := etcdCreate("127.0.0.1:8080"); err != nil {
		fmt.Println(err)
	}

	if addr, err := etcdGet(serviceName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(addr)
	}

	if addr, err := etcdDel(serviceName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(addr)
	}
}

func etcdCreate(addr string) error {
	etcdClient, err := etcd.NewFromURL(etcdctl)
	if err != nil {
		return err
	}
	resp, err := etcdClient.Put(context.Background(), serviceName, addr)
	if err != nil {
		log.Println("put failed")
		return err
	}
	fmt.Println(resp)
	return nil
}

func etcdGet(serviceName string) (addr string, err error) {
	etcdClient, err := etcd.NewFromURL(etcdctl)
	if err != nil {
		return "never", err
	}
	resp, err := etcdClient.Get(context.Background(), serviceName)
	if err != nil {
		return "never", err
	}
	for _, kv := range resp.Kvs {
		fmt.Println(kv.Key, kv.Value)
		addr = string(kv.Value)
	}
	return
}

func etcdDel(serviceName string) (addr string, err error) {
	etcdClient, err := etcd.NewFromURL(etcdctl)
	if err != nil {
		return "never", err
	}
	resp, err := etcdClient.Delete(context.Background(), serviceName)
	if err != nil {
		return "never", err
	}
	addr = string(resp.Deleted)
	return
}

func TestetcdLeave(serviceName string) {
	etcdClient, err := etcd.NewFromURL(etcdctl)
	if err != nil {
		log.Println("连接失败")
	}
	fmt.Println("connect to etcd success.")
	defer etcdClient.Close()

	// 创建一个5秒的租约
	resp, err := etcdClient.Grant(context.TODO(), ttl)
	if err != nil {
		log.Fatal(err)
	}

	// 5秒钟之后, /nazha/ 这个key就会被移除
	_, err = etcdClient.Put(context.TODO(), serviceName, "127.0.0.1：8080", etcd.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

}

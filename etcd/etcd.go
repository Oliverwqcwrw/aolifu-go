package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"time"
)

var client *clientv3.Client

func init() {
	client, _ = clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	fmt.Println("connect succ")
}

func main() {
	defer client.Close()
	TestKeepalive(client)

}

func TestDistributedLock(client *clientv3.Client) {
	s1, err := concurrency.NewSession(client)
	if err != nil {
		log.Fatal(err)
	}
}

func TestKeepalive(client *clientv3.Client) {
	resp, err := client.Grant(context.TODO(), 6)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Put(context.TODO(), "/lmh/", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
	ch, err := client.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		log.Fatal(err)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}

func TestGrant(client *clientv3.Client) {
	resp, err := client.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Put(context.TODO(), "/lmh/", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
}

func TestWatch(client *clientv3.Client) {
	rch := client.Watch(context.Background(), "key1")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Println("type:", ev.Type, "kv", string(ev.Kv.Key), string(ev.Kv.Value))
		}

	}
}

func TestSet(client *clientv3.Client) {
	conte, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := client.Put(conte, "key1", "value1")
	cancel()
	if err != nil {
		fmt.Println("put failed", err)
		return
	}
	conte, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := client.Get(conte, "key1")
	cancel()
	if err != nil {
		fmt.Println("get failed", err)
	}
	for _, env := range res.Kvs {
		fmt.Println("get succ", string(env.Key), string(env.Value))
	}
}

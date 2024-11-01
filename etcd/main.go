package main

import (
	"log"

	"github.com/rpcxio/libkv"
	estore "github.com/rpcxio/rpcx-etcd/store"
	"github.com/rpcxio/rpcx-etcd/store/etcdv3"
)

func main() {
	libkv.AddStore(estore.ETCDV3, etcdv3.New)

	kv, err := libkv.NewStore(estore.ETCDV3, []string{"127.0.0.1:2379"}, nil)
	if err != nil {
		log.Fatalf("cannot create etcd registry: %v", err)
		return
	}
	log.Printf("Successfully connected to etcd")

	key := "test"
	value := []byte("value")

	// 写入键值对
	err = kv.Put(key, value, nil)
	if err != nil {
		log.Fatalf("error putting key: %v", err)
	}
	log.Printf("Set key %s to value %s successfully", key, value)

	// 读取键值对
	kvPair, err := kv.Get(key)
	if err != nil {
		log.Fatalf("error getting value: %v", err)
	}

	retrievedValue := string(kvPair.Value)
	if retrievedValue != string(value) {
		log.Fatalf("value mismatch: expected %s, got %s", string(value), retrievedValue)
	} else {
		log.Printf("Key %s set successfully with value %s", key, retrievedValue)
	}

	log.Printf("[ok]")
}

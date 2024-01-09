package utils

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func EtcdPut(cli *clientv3.Client, key string, msg string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := cli.Put(ctx, key, msg)
	if err != nil {
		return err
	}
	return nil
}

func EtcdGet(cli *clientv3.Client, key string) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := cli.Get(ctx, key)
	if err != nil {
	}
	return resp, nil
}

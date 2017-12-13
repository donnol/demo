// 介绍 https://yuerblog.cc/2017/12/10/principle-about-etcd-v3/
// http://www.infoq.com/cn/articles/etcd-interpretation-application-scenario-implement-principle
package main

import (
	"context"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	log.Printf("cli is: %+v\n", cli)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Put(ctx, "/test/key", "value")
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	log.Printf("resp is: %+v\n", resp)
}

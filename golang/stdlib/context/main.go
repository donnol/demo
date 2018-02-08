package main

import (
	"context"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	contextTest()
	goroutineContext()
	otherContext()
}

type contextKey string

func contextTest() {
	ctx := context.Background()
	t, _ := ctx.Deadline()
	log.Println(t)
	// ch := ctx.Done()
	// <-ch
	err := ctx.Err()
	if err != nil {
		log.Fatal(err)
	}
	k := "hello"
	v := ctx.Value(k)
	log.Println(v)

	dCtx, cancelFunc := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	t, _ = ctx.Deadline()
	log.Println(t)
	t, _ = dCtx.Deadline()
	log.Println(t)
	cancelFunc()

	cCtx, cancelFunc := context.WithCancel(ctx)
	t, _ = ctx.Deadline()
	log.Println(t)
	t, _ = cCtx.Deadline()
	log.Println(t)
	cancelFunc()

	tCtx, cancelFunc := context.WithTimeout(ctx, time.Minute)
	t, _ = ctx.Deadline()
	log.Println(t)
	t, _ = tCtx.Deadline()
	log.Println(t)
	cancelFunc()

	v = "world"
	vCtx := context.WithValue(ctx, contextKey(k), v)
	t, _ = ctx.Deadline()
	log.Println(t)
	t, _ = vCtx.Deadline()
	log.Println(t)
	v = vCtx.Value(contextKey(k))
	log.Println(v)
}

func goroutineContext() {
	ctx := context.Background()

	// 主动调用 cancelFunc
	cCtx, cancelFunc := context.WithCancel(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println("go 1 finish")
				return
			default:
				log.Println("go 1 doing something")
			}
		}
	}(cCtx)

	// 调用 cancelFunc2
	// 时间到达 Deadline
	dCtx, cancelFunc2 := context.WithDeadline(ctx, time.Now().Add(time.Second))
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println("go 2 finish")
				return
			default:
				log.Println("go 2 doing something")
			}
		}
	}(dCtx)

	tCtx, cancelFunc3 := context.WithTimeout(ctx, time.Second*2)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println("go 3 finish")
				return
			default:
				log.Println("go 3 doing  something")
			}
		}
	}(tCtx)

	time.Sleep(time.Second)
	cancelFunc()

	time.Sleep(time.Second * 2) // 1s 后，打印停止，说明 go 2 已停止
	cancelFunc2()

	time.Sleep(time.Second * 3)

	k := "hello"
	tCtx, cancelFunc3 = context.WithTimeout(ctx, time.Second*2)
	vCtx := context.WithValue(tCtx, contextKey(k), "world")
	go func(ctx context.Context) {
		v := ctx.Value(contextKey(k))
		log.Println(v)
		for {
			select {
			case <-ctx.Done():
				log.Println("go 4 finish")
				return
			default:
				log.Println("go 4 doing something")
			}
		}
	}(vCtx)

	time.Sleep(time.Second * 3)
	cancelFunc3()
}

func otherContext() {
	todoCtx := context.TODO()
	dCtx, cancelFunc := context.WithDeadline(todoCtx, time.Now().Add(time.Second))
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println("go finish")
				return
			default:
				log.Println("go doing something")
			}
		}
	}(dCtx)

	time.Sleep(time.Second * 2)
	cancelFunc()
}

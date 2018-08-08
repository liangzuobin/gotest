package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// go func() {
	// 	time.AfterFunc(time.Second, cancel)
	// }()
	sleepAndTalk(ctx, 3*time.Second, "hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	case <-time.After(d):
		fmt.Println(msg)
	}
}

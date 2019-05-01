package main

import (
	"context"
	"fmt"
	"time"
)

func doCalculation(ctx context.Context, i int) {
	defer func() {
		fmt.Println("go doCalculation() main ctx ended returning")
	}()
	fmt.Printf("go doCalculation() ID: %d doing some stuff\n", i)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("go doCalculation() <-ctx.Done()")
			return
		default:
			// do something
		}
	}
}

func hub(ctx context.Context) {
	defer func() {
		fmt.Println("go hub() main ctx exnded returning")
	}()
	go doCalculation(ctx, 1)
	go doCalculation(ctx, 2)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("go hub() <-ctx.Done()")
			return
		default:
			// do nothing
		}
	}
}
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("main() executing hub")
	go hub(ctx)
	time.Sleep(1 * time.Second)
	fmt.Println("main() calling cancel()")
	cancel()
	time.Sleep(10 * time.Second)
}

package main

import (
	"context"
	"fmt"
	"time"
)

func doCalculation(ctx context.Context, i int) {
	go func() {
		fmt.Printf("go doCalculation() ID: %d doing some stuff\n", i)
		time.Sleep(100 * time.Second)
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("go doCalculation() context timeout")
			return
		default:
			// do something
		}
	}
}

func hub(ctx context.Context) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer func() {
		fmt.Println("go hub() returning")
		cancel()
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
	fmt.Println("main() executing hub")
	go hub(ctx)
	time.Sleep(10 * time.Second)
}

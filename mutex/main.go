package main

import (
	"fmt"
	"sync"
	"time"
)

type sharedData struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *sharedData) Increment(key string, index int) {
	fmt.Printf("goroutine ID: %d\n", index)
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *sharedData) Value(key string) int {
	defer c.mux.Unlock()
	c.mux.Lock()
	return c.v[key]
}
func main() {
	fmt.Println("MUTEX Example")
	data := sharedData{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go data.Increment("klucz", i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(data.Value("klucz"))
}

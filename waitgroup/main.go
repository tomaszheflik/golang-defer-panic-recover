package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type httpCrawler struct{}

func (httpCrawler) Get(url string) {
	rand.Seed(time.Now().UnixNano())
	wait := rand.Intn(1e4)
	time.Sleep(time.Duration(wait) * time.Millisecond)
	fmt.Printf("Time: %d to parse %s\n", wait, url)
}

var http httpCrawler

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.someslltupidname.com/",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			http.Get(url)
		}(url)
	}
	fmt.Println("Before wait")
	wg.Wait()
	fmt.Println("After wait")
}

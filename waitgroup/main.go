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

func main() {
	var http httpCrawler
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.sap.com/",
		"http://www.hybris.com/",
		"http://www.linux.org/",
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

package main

import (
	"fmt"
	"time"
)
func say(s string) {
	var i int
	for {
		fmt.Printf("%s -> %d\n", s, i)
		i++
		time.Sleep(1*time.Second)
	}
}
func main() {
	fmt.Println("Hello GO")
	go say("Hello from routine")
	go say("Hello from seccond routine")
	fmt.Println("Routine fired")
	time.Sleep(10*time.Second)
}
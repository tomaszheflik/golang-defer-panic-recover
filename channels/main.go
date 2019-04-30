package main

import (
	"fmt"
	"time"
)

func pinger(times int, channel chan string, done chan string) {
	for i := 0; i <= times; i++ {
		fmt.Printf("PINGER: message ID:%d\n", i)
		channel <- fmt.Sprintf("message ID:%d\n", i)
		time.Sleep(1 * time.Second)
	}
	done <- "PINGER: done"
}

func reciver(channel chan string) {
	for {
		select {
		case message := <-channel:
			fmt.Printf("RECIVER: %s\n", message)
		default:
			// do nothing
		}
	}
}

func main() {
	fmt.Println("MASTER: GO started")
	d := make(chan string)
	c := make(chan string)

	go reciver(c)
	go pinger(5, c, d)

mainloop:
	for {
		select {
		case done := <-d:
			fmt.Println(done)
			break mainloop
		default:
			fmt.Println("MASTER: heartbeat")
			time.Sleep(2000 * time.Millisecond)
		}
	}
}

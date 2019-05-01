package main

import (
	"fmt"
	"time"
)

func p1(max int, channel chan int, done chan string) {

	for {
		select {
		case message := <-channel:
			fmt.Printf("P1: recived %d\n", message)
			if message < max {
				message++
				fmt.Printf("P1: sent %d\n", message)
				channel <- message
			} else {
				done <- "P1 is done"
			}

		default:
			// do nothing
		}
	}
}

func p2(channel chan int) {
	for {
		select {
		case message := <-channel:
			fmt.Printf("P2: recived %d\n", message)
			message++
			fmt.Printf("P2: sent %d\n", message)
			channel <- message
		default:
			// do nothing
		}
	}
}

func start(channel chan int) {
	var i = 0
	fmt.Printf("START: sent %d\n", i)
	channel <- i
	return
}
func main() {
	fmt.Println("MASTER: GO started")
	d := make(chan string)
	c := make(chan int)

	go p2(c)
	go p1(5, c, d)
	start(c)
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

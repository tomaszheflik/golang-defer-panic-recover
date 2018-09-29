package main

import (
	"fmt"
)

func defered(s string) {
	fmt.Println("Defer got executed in ", s)
}

func recoverFrom() {
	r := recover()
	if r != nil {
		fmt.Println("Recovered from ", r)
	}
}

func generatePanicWithDefer(err error) {
	defer recoverFrom()
	panic(err)
}

func generatePanic(err error) {
	panic(err)
}

func main() {
	defer recoverFrom()
	generatePanicWithDefer(fmt.Errorf("First panic() with defer"))

	defer defered("main()")
	generatePanic(fmt.Errorf("Second panic without defer"))
}

package main

import (
	"fmt"
)

type data1 struct {
	Name string
	Role int
}

func (c *data1) Close(f string) {
	fmt.Printf("Executed from %s defer is executed\n", f)
}

func f1(d *data1) {
	defer d.Close("f1()")
	fmt.Println("Eecuted in f1()")
	f2(d)
	return
}

func f2(d *data1) {
	defer d.Close("f2()")
	fmt.Println("Executed in f2()")
	return
}

func f3(d *data1) {
	defer d.Close("f3()")
	fmt.Println("Executed in f3()")
}
func main() {
	d := data1{
		Name: "Tomek",
		Role: 1,
	}
	defer d.Close("main()")
	f3(&d)
	f1(&d)
}

package main

import "fmt"

func adder() func() {
	a := 1
	return func() {
		fmt.Println(a)
		a++
	}
}

func main() {
	f := adder()
	f()
	f()
	f()
}

package main

import "fmt"

type hotdog int

type person struct {
	Name   int
	Age    int
	height int
}

func main() {
	var oink hotdog
	x := 7
	xi := []int{2, 3, 4, 5, 23}
	m := map[string]int{
		"John":   43,
		"Martin": 423,
	}
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", xi)
	fmt.Printf("%T\n", m)

	fmt.Printf("%T\n", oink)
}

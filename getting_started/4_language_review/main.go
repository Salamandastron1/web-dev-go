package main

import "fmt"

type hotdog int

type person struct {
	name   string
	age    int
	height int
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println("My name is", p.name)
	fmt.Println("My age is", p.age)
	fmt.Println("My height is", p.height)
}
func (sa secretAgent) speak() {
	fmt.Println("---------------")
	fmt.Println("I'm saying something")
	fmt.Println("My name is", sa.name)
	fmt.Println("My age is", sa.age)
	fmt.Println("My height is", sa.height)
	fmt.Printf("Do I have a license to kill?: %v\n", sa.licenseToKill)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	var oink hotdog
	x := 7
	xi := []int{2, 3, 4, 5, 23}
	m := map[string]int{
		"John":   43,
		"Martin": 423,
	}
	p := person{"Thony", 18, 18}
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", xi)
	fmt.Printf("%T\n", m)

	fmt.Printf("%T\n", oink)
	fmt.Printf("%T\n", p)
	fmt.Println(p)
	p.speak()

	sa := secretAgent{
		person{
			"Thony",
			29,
			180,
		},
		true,
	}
	fmt.Println(sa)
	sa.speak()
	saySomething(sa)
	saySomething(p)
}

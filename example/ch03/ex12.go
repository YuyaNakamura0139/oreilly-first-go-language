package main

import "fmt"

func main() {
	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}
	fmt.Println("f:", f)
	fmt.Println("g:", g)
	g = f
	fmt.Println("g（fの代入後）:", g)
	fmt.Println("f==g:", f == g)

}

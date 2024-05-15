package main

import "fmt"

func main() {
	type firstPerson struct {
		name string
		age  int
	}

	type secondPerson struct {
		name string
		age  int
	}

	type thirdPerson struct {
		age  int
		name string
	}

	x1 := firstPerson{
		"太郎",
		24,
	}
	fmt.Println("x1:", x1)

	var x2 secondPerson
	x2 = secondPerson(x1)
	fmt.Println("x2:", x2)
	// fmt.Println("x1==x2:", x1 == x2)

	x12 := firstPerson(x2)
	fmt.Println("x1==x2:", x1 == x12)
}

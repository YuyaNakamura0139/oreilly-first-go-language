package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}

	// var fred person
	bob := person{}
	// julia := person{
	// 	"ジュリア",
	// 	40,
	// 	"cat",
	// }
	// beth := person{
	// 	age:  30,
	// 	name: "ベス",
	// }
	bob.name = "ボブ"
	fmt.Println(bob.name)
}

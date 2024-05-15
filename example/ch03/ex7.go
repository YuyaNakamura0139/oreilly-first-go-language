package main

import "fmt"

func main() {
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println(m["hello"])
	delete(m, "hello")
	fmt.Println(m["hello"])
}

package main

import "fmt"

func main() {
	var s *string
	fmt.Println(s == nil) // ポインタ型のゼロ値はnil
	var i interface{}
	fmt.Println(i == nil)
	i = s
	fmt.Println(i == nil)
}

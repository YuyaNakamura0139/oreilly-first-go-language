package main

import "fmt"

func main() {
	// 長さ5, キャパシティ10のスライス
	x := make([]int, 5, 10)
	x = append(x, 5, 6, 7, 8)
	x[1] = 100
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

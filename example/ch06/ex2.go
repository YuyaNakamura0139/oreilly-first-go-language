package main

import "fmt"

func main() {
	x := 10
	// ポインタ型の変数を宣言する際は、そのポインタが指す領域に保存される値の型の前に*をつけて表す
	var pointerToX *int
	pointerToX = &x
	fmt.Println(pointerToX)
}

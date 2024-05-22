package main

import "fmt"

func main() {
	x := 10
	// &は変数の前につけると変数のアドレスを返す。
	// アドレス演算子
	// pointerToXは変数xのアドレス
	pointerToX := &x
	fmt.Println(pointerToX)
	// *は間接参照のための演算子
	// ポインタ型の変数の前につけるとポインタが参照するアドレスに保存されている値を返す
	fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z)
}

package main

import (
	"fmt"
)

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder10 := Adder{start: 10}
	fmt.Println(myAdder10.AddTo(5))

	f1 := myAdder10.AddTo // Adder型の変数myAdder10のメソッドAddToをf1に代入
	fmt.Println(f1(10))

	f2 := Adder.AddTo // 型Adderをレシーバとして定義されているメソッドAddToをf2に代入
	fmt.Println(f2(myAdder10, 15))
}

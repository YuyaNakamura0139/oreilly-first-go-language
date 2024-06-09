package main

import "fmt"

func main() {
	// デッドロックになってしまう例
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v // mainのゴルーチンによる読み込みが開始されるまで書き込まれない
		v2 := <-ch2
		fmt.Println(v, v2)
	}()

	v := 2
	ch2 <- v // 無名関数によるch2の読み込みが行われるまで書き込まれない
	v2 := <-ch1
	fmt.Println(v, v2)
}

package main

import "fmt"

type Score int
type HighScore Score

type Person struct {
	LastName  string
	FirstName string
	Age       int
}
type Employee Person

func (s Score) Double() Score {
	return s * 2
}

func main() {
	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	// hs = s // コンパイル時のエラー
	// s = i  // コンパイル時のエラー
	s = Score(i) // 型変換後に代入
	hs = HighScore(s)
	fmt.Println(s, hs)
	hhs := hs + 20 // 基底型(この場合int)に対して使える演算子は使える
	fmt.Println(hhs)

	s = 200
	hs = 300
	fmt.Println(s.Double())
	fmt.Println(Score(hs).Double())
	// fmt.Println(hs.Double()) // コンバイル時のエラー
}

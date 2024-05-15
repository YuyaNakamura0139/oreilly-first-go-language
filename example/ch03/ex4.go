package main

import "fmt"

func main() {
	totalWins := map[string]int{} // string->intのマップを要素なしで初期化
	fmt.Println(totalWins == nil)
	fmt.Println(totalWins["abc"])
	totalWins["abc"] = 3          // 書き込めるようになる
	fmt.Println(totalWins["abc"]) // 読み込みもできる
}

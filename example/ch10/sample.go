package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 現在のGoルーチンの数を取得
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())

	// ガベージコレクションを強制的に実行
	runtime.GC()

	// OSスレッドの数を取得
	fmt.Println("Number of CPUs:", runtime.NumCPU())
}

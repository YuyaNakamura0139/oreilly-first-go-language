package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(10)
	if n == 0 {
		fmt.Println("少し小さすぎます:", n)
	} else if n > 5 {
		fmt.Println("大きすぎます:", n)
	} else {
		fmt.Println("いい感じの数字です:", n)
	}
}

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	words := []string{"山", "sun", "微笑み", "人類学者", "モグラの穴", "mountain", "タコの足とイカの足", "antholopologist", "タコのあしは8本で以下の足は10本"}
	for _, word := range words {
		switch size := utf8.RuneCountInString(word); size {
		case 1, 2, 3, 4:
			fmt.Printf("「%s」の文字数は%dで、短い単語だ。\n", word, size)
		case 5:
			fmt.Printf("「%s」の文字数は%dで、これはちょうど良い長さだ。\n", word, size)
		case 6, 7, 8, 9:
		default:
			fmt.Printf("「%s」の文字数は%dで、とても長い。\n", word, size)
			if n := len(word); size < n {
				fmt.Printf("%dバイトもある！\n", n)
			} else {
				fmt.Println()
			}
		}
	}
}

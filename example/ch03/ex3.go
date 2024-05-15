package main

import "fmt"

func main() {
	var nilMap map[string]int  // 初期値はnil
	fmt.Println(len(nilMap))   // 0
	fmt.Println(nilMap["abc"]) // 0
	nilMap["abc"] = 3          // panic
}

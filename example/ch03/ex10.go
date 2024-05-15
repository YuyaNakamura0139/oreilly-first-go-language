package main

import "fmt"

func main() {
	// 無名構造体
	// Jsonのアンマーシャリング、マーシャリング時に使用s
	var person struct {
		name string
		age  int
		pet  string
	}
	person.name = "ボブ"
	person.age = 50
	person.pet = "dog"

	pet := struct {
		name string
		kind string
	}{
		name: "ポチ",
		kind: "dog",
	}

	fmt.Println(person.name)
	fmt.Println(pet.name)
}

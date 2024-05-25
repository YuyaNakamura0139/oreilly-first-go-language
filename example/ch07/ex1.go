package main

import "fmt"

// Personという名前のユーザー定義型を宣言。
// 構造体リテラルを規定型としてもつ
type Person struct {
	LastName  string
	FirstName string
	Age       int
}

// 型Personに付随するメソッドStringを定義
func (p Person) String() string { // (p Person)がレシーバの指定
	return fmt.Sprintf("%s %s : 年齢%d", p.LastName, p.FirstName, p.Age)
}

func main() {
	yuya := Person{
		LastName:  "中村",
		FirstName: "友哉",
		Age:       25,
	}
	fmt.Println(yuya.String())
}

// jsonへの変換はマーシャリング
// jsonからの変換はアンマーシャリングと呼ぶ

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	f := struct {
		Name string
		Age  int
	}{}

	err := json.Unmarshal([]byte(`{"name": "小野小町", "occupation": "歌人", "age": 20}`), &f)
	// 大文字小文字の違いを無視して、フィールドに対応付けてくれる
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", f) // フィールド名付きで出力される
	fmt.Println(f)
}

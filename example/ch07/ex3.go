package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

// ポインタレシーバ
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// 値レシーバ
func (c Counter) String() string {
	return fmt.Sprintf("合計: %d, 更新: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) { // 元の値は変更されない、間違いの例
	c.Increment() // mainのcのコピーに対してIncrementが行われる
	fmt.Println("NG:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("OK:", c.String())
}

func main() {
	var c Counter
	doUpdateWrong(c)
	fmt.Println("main:", c.String())
	doUpdateRight(&c)
	fmt.Println("main:", c.String())
}

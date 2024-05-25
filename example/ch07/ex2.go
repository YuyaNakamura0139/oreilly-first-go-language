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

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}

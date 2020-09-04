package a

import (
	"fmt"
)

type Container struct {
	Values [2]string
}

func (c *Container) SetValues(values [2]string) {
	// 引数のスライスで受け取ったスライスがそのままフィールドに保存されている
	c.Values = values
}

func main() {
	c := &Container{}
	list := [2]string{"hello", "world"}
	c.SetValues(list)
	fmt.Println(c)
	// その関数の引数に渡したスライスがあとで要素が変更されている
	list[1] = "tenntenn"
	fmt.Println(c)
}

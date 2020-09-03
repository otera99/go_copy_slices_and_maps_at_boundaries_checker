package a

import (
	"fmt"
)

type Container struct {
	Values []string
}

func (c *Container) SetValues(values []string) {
	// 引数のスライスで受け取ったスライスがそのままフィールドに保存されている
	c.Values = values
}

func (c *Container) CorrectSetValues(values []string) {
	// 正しい書き方をしている場合
	vs := make([]string, len(values))
	copy(vs, values)
	c.Values = vs
}

func main() {
	c := &Container{}
	d := &Container{}
	list := []string{"hello", "world"}
	list1 := []string{"hello", "world"}
	c.SetValues(list)
	d.CorrectSetValues(list1)
	fmt.Println(c)
	// その関数の引数に渡したスライスがあとで要素が変更されている
	list[1] = "tenntenn"
	list1[1] = "tenntenn"
	fmt.Println(c)
}

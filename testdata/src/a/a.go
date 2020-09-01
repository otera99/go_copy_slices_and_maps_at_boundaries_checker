package a

import (
	"fmt"
)

type Container struct {
	Values []string
}

func (c *Container) SetValues(values []string) {
	// 本当はこう書かないとだめ
	// vs := make([]string, len(values))
	// copy(vs, values)
	// c.Values = vs
	// 引数のスライスで受け取ったスライスがそのままフィールドに保存されている
	c.Values = values
}

func main() {
	c := &Container{}
	list := []string{"hello", "world"}
	c.SetValues(list)
	fmt.Println(c)
	// その関数の引数に渡したスライスがあとで要素が変更されている
	list[1] = "tenntenn"
	fmt.Println(c)
}

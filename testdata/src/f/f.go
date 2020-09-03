package f

import (
	"fmt"
)

type Container struct {
	Values map[string]int
}

func (c *Container) SetValues(values map[string]int) {
	// 引数のマップで受け取ったマップがそのままフィールドに保存されている
	c.Values = values
}

func (c *Container) CorrectSetValues(values map[string]int) {
	// 正しい書き方をしている場合
	vs := make(map[string]int, len(values))
	for k, v := range values {
		vs[k] = v
	}
	c.Values = vs
}

func main() {
	c := &Container{}
	d := &Container{}
	m := map[string]int{"Foo": 1, "Bar": 2}
	m1 := map[string]int{"Foo": 1, "Bar": 2}
	c.SetValues(m)
	d.CorrectSetValues(m1)
	fmt.Println(c)
	fmt.Println(d)
	// その関数の引数に渡したスライスがあとで要素が変更されている
	m["Foo"] = 3  // want "WARN"
	m1["Foo"] = 3
	fmt.Println(c)
	fmt.Println(d)
}

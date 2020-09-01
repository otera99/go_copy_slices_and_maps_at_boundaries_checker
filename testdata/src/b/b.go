package b

import (
	"fmt"
)

type Container struct {
	Values []string
}

func (c *Container) SetValues(values []string) {
	// 正しい書き方をしている場合
	vs := make([]string, len(values))
	copy(vs, values)
	c.Values = vs
	// c.Values = values
}

func main() {
	c := &Container{}
	list := []string{"hello", "world"}
	c.SetValues(list)
	fmt.Println(c)
	list[1] = "tenntenn"
	fmt.Println(c)
}

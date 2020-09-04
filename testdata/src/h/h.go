package a

import (
	"fmt"
)

type Container struct {
	Values [2]string
}

func (c *Container) SetValues(values [2]string) {
	c.Values = values
}

func main() {
	c := &Container{}
	list := [2]string{"hello", "world"}
	c.SetValues(list)
	fmt.Println(c)
	list[1] = "tenntenn"
	fmt.Println(c)
}

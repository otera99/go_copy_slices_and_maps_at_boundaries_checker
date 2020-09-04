package c

import (
	"fmt"
)

type Container struct {
	Values []string
}

func (c *Container) SetValues(values []string) string {
	c.Values = values
	return "hello"
}

func main() {
	c := &Container{}
	list := []string{"hello", "world"}
	u := c.SetValues(list)
	fmt.Println(u)
	fmt.Println(c)
	list[1] = "tenntenn" // want "WARN: A slice taken as an argument and stored in a field is rewritten."
	fmt.Println(c)
}
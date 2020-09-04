package d

import (
	"fmt"
)

type Container struct {
	Values []string
}

func (c *Container) SetValues(values []string, values2 []string) {
	values2[1] = "otera"
	c.Values = values
}

func main() {
	c := &Container{}
	list := []string{"hello", "world"}
	list2 := []string{"hello", "world"}
	c.SetValues(list, list2)
	fmt.Println(c)
	list[1] = "tenntenn" // want "WARN: Slice or map taken as an argument and stored in a field may be rewritten."
	list2[1] = "tenntenn"
	fmt.Println(c)
}
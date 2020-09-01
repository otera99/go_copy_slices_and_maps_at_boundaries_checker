package b

import (
	"fmt"
)

type Foo struct {
	Values []string
}

func (f *Foo) SetVal(s []string) {
	f.Values = func()[]string {
		return s
	}()
}

func main() {
	f := &Foo{}
	v := []string{"hello", "world"}
	f.SetVal(v)
	fmt.Println(f)
	v[1] = "tenntenn"
	fmt.Println(f)
}

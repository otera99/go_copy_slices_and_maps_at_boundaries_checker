package e

import (
	"fmt"
)

type Foo struct {
	Values []string
}

func (f *Foo) SetVal(s []string) {
	f.Values = func()[]string {
	        t := s
		return t
	}()
}

func main() {
	f := &Foo{}
	v := []string{"hello", "world"}
	f.SetVal(v)
	fmt.Println(f)
	// 未対応だけど、"WARN"をだしたい
	v[1] = "tenntenn"
	fmt.Println(f)
}

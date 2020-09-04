# go_copy_slices_and_maps_at_boundaries_checker
An individual project at mercari summer internship 2020.

メルカリサマーインターン2020での成果物です。

This project is inspired by [knsh14](https://github.com/knsh14)'s idea, who is an extreamely talented engineer at mercari inc.

このプロジェクトの大元のアイデアはメルカリのエンジニアである[knsh14](https://github.com/knsh14)氏によるものです。

## Description
https://github.com/knsh14/uber-style-guide-ja/blob/master/guide.md#copy-slices-and-maps-at-boundaries にあるような、スライスやマップは内部でデータへのポインタが含まれていることを考慮せずにコピーしているコードに対してデータが書き換わる可能性がある箇所で警告をだすツールをskeletonを用いて作りました。

### example
```
package main

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
	c.Values = values
}

func main() {
	c := &Container{}
	list := []string{"hello", "world"}
	c.SetValues(list)
	fmt.Println(c)
	list[1] = "tenntenn" //ここで警告を出す
	fmt.Println(c)
}
```

## install
```go get -u github.com/otera99/go_copy_slices_and_maps_at_boundaries_checker```
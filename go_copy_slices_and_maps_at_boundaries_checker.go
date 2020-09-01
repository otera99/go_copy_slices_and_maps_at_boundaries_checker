package go_copy_slices_and_maps_at_boundaries_checker

import (
	"go/ast"
	"fmt"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "go_copy_slices_and_maps_at_boundaries_checker is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "go_copy_slices_and_maps_at_boundaries_checker",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	mFunc := map[string]bool{}
	// mSlice := map[string]bool{}

	// 引数のスライスで受け取ったスライスもしくはマップがそのままフィールドに保存されている関数があるかを調べるパート
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch t := n.(type) {
		case *ast.FuncDecl:
			if t.Recv != nil {
				for _, rev := range t.Recv.List {
					// sliceがあったらメモメモ
					if rev == nil {
						continue
					}
					fmt.Println(rev)
				}
			}

			check := false
			for _, stmt := range t.Body.List {
				// fmt.Println(stmt)
				switch u := stmt.(type) {
				case *ast.AssignStmt:
					if u.Lhs != nil && u.Rhs != nil {
						fmt.Println(u.Lhs[0])
					}
				}
			}
			mFunc[t.Name.Name] = check
		}
	})


	// その関数の引数に渡したスライスもしくはマップがあとで要素が変更されてたら警告するパート
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch t := n.(type) {
		case *ast.CallExpr:
			
		case *ast.AssignStmt:
			fmt.Println(t)
			if t.Lhs != nil && t.Rhs != nil {
				fmt.Println(t.Lhs[0])
			}
		}
	})

	return nil, nil
}


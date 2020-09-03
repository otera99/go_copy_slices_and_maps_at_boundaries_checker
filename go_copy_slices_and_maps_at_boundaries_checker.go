package go_copy_slices_and_maps_at_boundaries_checker

import (
	"go/ast"
	"go/types"
	"go/token"
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

func under(t types.Type) types.Type {
	if named, _ := t.(*types.Named); named != nil {
		return under(named.Underlying())
	}
	return t
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.CallExpr)(nil),
		(*ast.AssignStmt)(nil),
	}

	mFunc := map[string]bool{}
	mSlice := map[string]bool{}

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
					//fmt.Println(rev)
				}
			}

			check := false
			for _, stmt := range t.Body.List {
				// fmt.Println(stmt)
				switch u := stmt.(type) {
				case *ast.AssignStmt:
					if u.Lhs != nil && u.Rhs != nil {
						switch v := y.Lhs[0].(type) {
							
						}
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
			funcName := ""
			switch u := t.Fun.(type) {
			case *ast.BasicLit :
				funcName = u.Value
			}
			if(mFunc[funcName]) {
				for _, arg := range t.Args {
					fmt.Println(arg)
				}
			}
		case *ast.AssignStmt:
			if t.Lhs != nil && t.Rhs != nil {
				switch u := t.Lhs[0].(type) {
				case *ast.IndexExpr :
					switch v := u.X.(type) {
					case *ast.BasicLit :
						if v.Kind == token.STRING {
							sliceName := v.Value
							if(mSlice[sliceName]) {
								pass.Reportf(v.Pos(), "WARN")
							}
						}
					}
				}
			}
		}
	})

	return nil, nil
}


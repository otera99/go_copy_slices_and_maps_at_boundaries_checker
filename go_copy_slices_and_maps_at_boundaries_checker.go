package go_copy_slices_and_maps_at_boundaries_checker

import (
	"go/ast"
	"go/types"

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

type Pair struct {
	Func   types.Object
	ArgNum int
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.CallExpr)(nil),
		(*ast.AssignStmt)(nil),
	}

	mFunc := map[types.Object]bool{}
	mSlice := map[types.Object]bool{}
	mPair := map[Pair]bool{}

	// 引数のスライスで受け取ったスライスもしくはマップがそのままフィールドに保存されている関数があるかを調べるパート
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch t := n.(type) {
		case *ast.FuncDecl:
			funcObj := pass.TypesInfo.ObjectOf(t.Name)

			var recvObj types.Object
			if t.Recv != nil && t.Recv.List != nil {
				if t.Recv.List[0].Names != nil {
					recvObj = pass.TypesInfo.ObjectOf(t.Recv.List[0].Names[0])
				}
			}
			mArgUsed := map[types.Object]bool{}
			mArgNum := map[types.Object]int{} 
			for i, arg := range t.Type.Params.List {
				if arg.Names != nil {
					argObj := pass.TypesInfo.ObjectOf(arg.Names[0])
					mArgUsed[argObj] = true
					mArgNum[argObj] = i
				}
			}

			check := false
			for _, stmt := range t.Body.List {
				switch u := stmt.(type) {
				case *ast.AssignStmt:
					if u.Lhs != nil && u.Rhs != nil {
						var stObj types.Object
						switch v := u.Lhs[0].(type) {
						case *ast.SelectorExpr:
							switch w := v.X.(type) {
							case *ast.Ident:
								stObj = pass.TypesInfo.ObjectOf(w)
							}
						}
						// u.Rhs[0] が *ast.CallExpr のエッヂケースにも対応する(testdate の b.go)
						var sliceOrMapObj types.Object
						switch v := u.Rhs[0].(type) {
						case *ast.Ident:
							switch pass.TypesInfo.TypeOf(v).(type) {
							case *types.Slice:
								sliceOrMapObj = pass.TypesInfo.ObjectOf(v)
							case *types.Map:
								sliceOrMapObj = pass.TypesInfo.ObjectOf(v)
							}
						}
						if stObj != nil && sliceOrMapObj != nil && recvObj == stObj && mArgUsed[sliceOrMapObj] {
							check = true
							mPair[Pair{funcObj, mArgNum[sliceOrMapObj]}] = true
						}
					}
				}
			}
			mFunc[funcObj] = check
		}
	})


	// その関数の引数に渡したスライスもしくはマップがあとで要素が変更されてたら警告するパート
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch t := n.(type) {
		case *ast.CallExpr:
			switch u := t.Fun.(type) {
			case *ast.SelectorExpr:
				funcObj := pass.TypesInfo.ObjectOf(u.Sel)
				if mFunc[funcObj] {
					for i, arg := range t.Args {
						if mPair[Pair{funcObj, i}] {
							switch v := arg.(type) {
							case *ast.Ident:
								sliceObj := pass.TypesInfo.ObjectOf(v)
								mSlice[sliceObj] = true
							}
						}
					}
				}
			}
		case *ast.AssignStmt:
			if t.Lhs != nil && t.Rhs != nil {
				switch u := t.Lhs[0].(type) {
				case *ast.IndexExpr:
					switch v := u.X.(type) {
					case *ast.Ident:
						obj :=  pass.TypesInfo.ObjectOf(v)
						if mSlice[obj] {
							pass.Reportf(u.Pos(), "WARN")
						}
					}
				}
			}
		}
	})

	return nil, nil
}


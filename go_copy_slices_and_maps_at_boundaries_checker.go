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

	/*
	receiving slices and maps at boundaries
	*/

	mFunc := map[types.Object]bool{}
	mSliceOrMap := map[types.Object]bool{}
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
								sliceOrMapObj := pass.TypesInfo.ObjectOf(v)
								mSliceOrMap[sliceOrMapObj] = true
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
						if mSliceOrMap[obj] {
							pass.Reportf(u.Pos(), "WARN: Slice or map taken as an argument and stored in a field may be rewritten.")
						}
					}
				}
			}
		}
	})

	/*
	returning slices and maps at boundaries
	*/

	mNotGoodFunc := map[types.Object]bool{}

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
			check := false
			unlock := false
			for _, stmt := range t.Body.List {
				switch u := stmt.(type) {
				case *ast.ExprStmt:
					switch v := u.X.(type) {
					case *ast.CallExpr:
						switch w := v.Fun.(type) {
						case *ast.SelectorExpr:
							switch x := w.X.(type) {
							case *ast.SelectorExpr:
								switch y := x.X.(type) {
								case *ast.Ident: 
									stObj := pass.TypesInfo.ObjectOf(y)
									if recvObj == stObj && w.Sel !=  nil && w.Sel.Name == "Lock" {
										unlock = true
									}
								}
							}
						}
					}
				case *ast.ReturnStmt:
					ret := u.Results
					if ret != nil && unlock {
						for _, res := range ret {
							switch v := res.(type) {
							case *ast.SelectorExpr :
								switch w := v.X.(type) {
								case *ast.Ident :
									stObj := pass.TypesInfo.ObjectOf(w)
									if recvObj == stObj && v.Sel != nil {
										switch pass.TypesInfo.TypeOf(v.Sel).(type) {
										case *types.Slice:
											check = true
										case *types.Map:
											check = true
										}
									}
								}
							}
						}
					}
				}
			}
			mNotGoodFunc[funcObj] = check
		}
	})

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch t := n.(type) {
		case *ast.CallExpr :
			switch u := t.Fun.(type) {
			case *ast.SelectorExpr:
				funcObj := pass.TypesInfo.ObjectOf(u.Sel)
				if(mNotGoodFunc[funcObj]) {
					pass.Reportf(u.Pos(), "WARN: Slices or maps that are kept internally without being made public may be changed.")
				}
			}
		}
	})

	return nil, nil
}


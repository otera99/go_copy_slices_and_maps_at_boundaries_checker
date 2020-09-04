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

	/*
	returning slices and maps
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
					pass.Reportf(u.Pos(), "WARN")
				}
			}
		}
	})

	return nil, nil
}


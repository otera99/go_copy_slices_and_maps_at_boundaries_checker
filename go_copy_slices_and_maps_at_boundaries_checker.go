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

	return nil, nil
}


package main

import (
	"github.com/otera99/go_copy_slices_and_maps_at_boundaries_checker"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(go_copy_slices_and_maps_at_boundaries_checker.Analyzer) }


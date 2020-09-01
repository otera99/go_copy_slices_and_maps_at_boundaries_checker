package go_copy_slices_and_maps_at_boundaries_checker_test

import (
	"testing"

	"github.com/otera99/go_copy_slices_and_maps_at_boundaries_checker"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, go_copy_slices_and_maps_at_boundaries_checker.Analyzer, "a")
	analysistest.Run(t, testdata, go_copy_slices_and_maps_at_boundaries_checker.Analyzer, "b")
	analysistest.Run(t, testdata, go_copy_slices_and_maps_at_boundaries_checker.Analyzer, "c")
	analysistest.Run(t, testdata, go_copy_slices_and_maps_at_boundaries_checker.Analyzer, "d")
}


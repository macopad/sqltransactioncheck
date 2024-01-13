package analyzer_test

import (
	"github.com/macopad/sqltransactioncheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestNewAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	checker := analyzer.NewAnalyzer()

	packages := []string{
		"github.com/macopad/sqltransactioncheck/pkg/analyzer/testdata/trans",
	}

	for _, pkg := range packages {
		pkg := pkg

		t.Run(pkg, func(t *testing.T) {
			t.Parallel()

			analysistest.Run(t, testdata, checker, pkg)
		})
	}
}

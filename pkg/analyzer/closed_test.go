package analyzer_test

import (
	"testing"

	"github.com/macopad/sqltransactioncheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func XTestClosedAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	checker := analyzer.NewClosedAnalyzer()

	packages := []string{
		"github.com/macopad/sqltransactioncheck/testdata/trans_samples",
	}

	for _, pkg := range packages {
		pkg := pkg

		t.Run(pkg, func(t *testing.T) {
			t.Parallel()

			analysistest.Run(t, testdata, checker, pkg)
		})
	}
}

package transanalyzer_test

import (
	transanalyzer "testdata/trans"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestTransanalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	checker := transanalyzer.NewTransanalyzer()

	packages := []string{
		"github.com/macopad/sqltransactioncheck/pkg/transanalyzer/testdata/trans",
	}

	for _, pkg := range packages {
		pkg := pkg

		t.Run(pkg, func(t *testing.T) {
			t.Parallel()

			analysistest.Run(t, testdata, checker, pkg)
		})
	}
}

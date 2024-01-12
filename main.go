package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"testdata/trans_examples/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.NewAnalyzer())
}

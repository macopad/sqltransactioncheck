package main

import (
	"github.com/macopad/sqltransactioncheck/pkg/transanalyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(transanalyzer.NewTransanalyzer())
}

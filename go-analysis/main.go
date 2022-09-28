package main

import (
	"github.com/luisacatherine/sample-golang-linter/go-analysis/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}

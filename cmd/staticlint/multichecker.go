package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"golang.org/x/tools/go/analysis/passes/deepequalerrors"
	"golang.org/x/tools/go/analysis/passes/loopclosure"
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"golang.org/x/tools/go/analysis/passes/unreachable"
	"honnef.co/go/tools/staticcheck"
)

func main() {
	var myChecks []*analysis.Analyzer
	for _, v := range staticcheck.Analyzers {
		myChecks = append(myChecks, v.Analyzer)
	}
	myChecks = append(myChecks,
		copylock.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		unreachable.Analyzer,
		printf.Analyzer,
		structtag.Analyzer,
		exitCheckAnalyzer,
		deepequalerrors.Analyzer,
		buildssa.Analyzer,
	)
	multichecker.Main(
		myChecks...,
	)
}

var exitCheckAnalyzer = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "checks for direct calls to os.Exit in main function of main package",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.Name == "main" {
			for _, decl := range file.Decls {
				if funcDecl, ok := decl.(*ast.FuncDecl); ok && funcDecl.Name.Name == "main" {
					ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
						if callExpr, ok := n.(*ast.CallExpr); ok {
							if fun, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
								if pkg, ok := fun.X.(*ast.Ident); ok && pkg.Name == "os" && fun.Sel.Name == "Exit" {
									pass.Reportf(callExpr.Pos(), "direct call to os.Exit in main function is prohibited")
								}
							}
						}
						return true
					})
				}
			}
		}
	}
	return nil, nil
}

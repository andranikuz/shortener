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
	var customAnalyzers []*analysis.Analyzer
	for _, v := range staticcheck.Analyzers {
		customAnalyzers = append(customAnalyzers, v.Analyzer)
	}
	customAnalyzers = append(customAnalyzers,
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
		customAnalyzers...,
	)
}

var exitCheckAnalyzer = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "checks for direct calls to os.Exit in main function of main package",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		checkMainPackage(pass, file)
	}
	return nil, nil
}

func checkMainPackage(pass *analysis.Pass, file *ast.File) {
	if file.Name.Name == "main" {
		checkMainFunction(pass, file)
	}
}

func checkMainFunction(pass *analysis.Pass, file *ast.File) {
	for _, decl := range file.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok && funcDecl.Name.Name == "main" {
			ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
				return checkOsExit(pass, n)
			})
		}
	}
}

func checkOsExit(pass *analysis.Pass, n ast.Node) bool {
	if callExpr, ok := n.(*ast.CallExpr); ok {
		if fun, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
			if pkg, ok := fun.X.(*ast.Ident); ok && pkg.Name == "os" && fun.Sel.Name == "Exit" {
				pass.Reportf(callExpr.Pos(), "direct call to os.Exit in main function is prohibited")
			}
		}
	}
	return true
}

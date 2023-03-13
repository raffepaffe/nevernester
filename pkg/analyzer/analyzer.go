package analyzer

import (
	"flag"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const (
	name              = "nevernester"
	doc               = "Checks the level of nesting in functions"
	defaultMaxNesting = 4
)

//nolint:gochecknoglobals
var (
	maxNesting     int
	skipTests      bool
	skipBenchmarks bool
	flagSet        flag.FlagSet
)

//nolint:gochecknoinits
func init() {
	flagSet.IntVar(&maxNesting, "maxNesting", defaultMaxNesting, "max nesting of the function can have")
	flagSet.BoolVar(&skipTests, "skipTests", false, "should functions starting with Test be checked")
	flagSet.BoolVar(&skipBenchmarks, "skipBenchmarks", false, "should functions starting with Benchmark be checked")
}

func New() *analysis.Analyzer {
	a := &analysis.Analyzer{
		Name:     name,
		Doc:      doc,
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    flagSet,
	}

	return a
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			fs := pass.Fset
			fd, ok := node.(*ast.FuncDecl)
			if !ok {
				return true
			}

			if skipTests && isFuncNamePrefix(fd, "test") {
				return true
			}
			if skipBenchmarks && isFuncNamePrefix(fd, "benchmark") {
				return true
			}

			totalNesting := nesting(fd, fs)
			if totalNesting > maxNesting {
				pass.Reportf(node.Pos(), "calculated nesting for function %s is %d, max is %d", fd.Name, totalNesting, maxNesting)
			}

			return true
		})
	}

	return nil, nil //nolint:nilnil
}

func isFuncNamePrefix(fd *ast.FuncDecl, prefix string) bool {
	funcName := strings.ToLower(fd.Name.Name)

	return strings.HasPrefix(funcName, prefix)
}

func nesting(fd *ast.FuncDecl, fs *token.FileSet) int {
	nv := newNesterVisitor(fd, fs)
	ast.Walk(&nv, fd)

	return nv.Indentation
}

type nesterVisitor struct {
	Indentation int
	Colum       int
	Line        int
	FileSet     *token.FileSet
	SkipRow     map[token.Position]bool
}

func newNesterVisitor(fd *ast.FuncDecl, fs *token.FileSet) nesterVisitor {
	absPosition := fs.PositionFor(fd.Pos(), false)
	sr := make(map[token.Position]bool, 0)
	nv := nesterVisitor{
		Indentation: 0,
		Colum:       absPosition.Column,
		Line:        absPosition.Line,
		FileSet:     fs,
		SkipRow:     sr,
	}

	return nv
}

// Visit checks if a Node should be counted as a new indentation.
// First, nodes that should be skipped are identified.
// After that, the position of the node is determined. Then, if the node is on a new line
// and the column count (indentation) of the node is larger than the
// last node, the indentation counter is added by one.
func (nv *nesterVisitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nv
	}

	nv.skipElseStatement(n)

	absPosition := nv.FileSet.PositionFor(n.Pos(), true)
	if absPosition.Line != nv.Line {
		nv.Line = absPosition.Line
		if absPosition.Column > nv.Colum {
			nv.addIndentation(absPosition)
		}
	}

	return nv
}

func (nv *nesterVisitor) addIndentation(absPosition token.Position) {
	_, found := nv.SkipRow[absPosition]
	if !found {
		nv.Indentation++
		nv.Colum = absPosition.Column
	}
}

// skipElseStatement finds else statements and excludes them from
// the indentation counting. The position of the else is saved.
func (nv *nesterVisitor) skipElseStatement(n ast.Node) {
	if ifBlock, ok := n.(*ast.IfStmt); ok {
		if ifBlock.Else != nil {
			nv.addRowToSkip(ifBlock.Else.Pos())
		}
	}
}

func (nv *nesterVisitor) addRowToSkip(pos token.Pos) {
	if pos.IsValid() {
		elsePosition := nv.FileSet.PositionFor(pos, true)
		nv.SkipRow[elsePosition] = true
	}
}

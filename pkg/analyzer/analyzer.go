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
)

func parseFlags() flag.FlagSet {
	flagSet := flag.NewFlagSet("nevernester", flag.ContinueOnError)
	flagSet.IntVar(&maxNesting, "max-nesting", defaultMaxNesting, "max nesting the function can have")
	flagSet.BoolVar(&skipTests, "skip-tests", false, "should functions starting with Test be checked")
	flagSet.BoolVar(&skipBenchmarks, "skip-benchmarks", false, "should functions starting with Benchmark be checked")

	return *flagSet
}

func New() *analysis.Analyzer {
	flagSet := parseFlags()
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
	SkipRow     map[int]bool
}

func newNesterVisitor(fd *ast.FuncDecl, fs *token.FileSet) nesterVisitor {
	absPosition := fs.PositionFor(fd.Pos(), false)
	sr := make(map[int]bool, 0)
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
// First, the position of the node is determined. Then, if the node is on a new line
// and the column count (indentation) of the node is larger than the
// last saved node, the indentation counter is added by one.
// After that, nodes that should be skipped are identified.
func (nv *nesterVisitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nv
	}

	startPosition := nv.FileSet.PositionFor(n.Pos(), true)
	if startPosition.Line != nv.Line {
		nv.Line = startPosition.Line
		if startPosition.Column > nv.Colum {
			nv.addIndentation(startPosition)
		}
		nv.addNodeToSkip(n)
	}

	return nv
}

func (nv *nesterVisitor) addIndentation(startPosition token.Position) {
	if nv.usePosition(startPosition) {
		nv.Indentation++
		nv.Colum = startPosition.Column
	}
}

func (nv *nesterVisitor) usePosition(positions ...token.Position) bool {
	for _, position := range positions {
		if _, found := nv.SkipRow[position.Line]; found {
			return false
		}
	}

	return true
}

// addNodeToSkip finds rows for the given node and excludes them from
// the indentation counting. If an expression or statement is
// spread out over several rows, all rows are excluded.
//
//nolint:all
func (nv *nesterVisitor) addNodeToSkip(n ast.Node) {
	rowsToSkip := make([]int, 0)
	switch n.(type) {
	case *ast.ExprStmt, *ast.AssignStmt:
		rowsToSkip = nv.findRowsToSkip(n)
	case *ast.IfStmt:
		ifStmt := n.(*ast.IfStmt)
		rowsToSkip = append(rowsToSkip, nv.FileSet.PositionFor(ifStmt.Pos(), true).Line)
		rowsToSkip = append(rowsToSkip, nv.FileSet.PositionFor(ifStmt.End(), true).Line)
		rowsToSkip = append(rowsToSkip, nv.findRowsToSkip(ifStmt.Init)...)
		rowsToSkip = append(rowsToSkip, nv.findRowsToSkip(ifStmt.Cond)...)
		if ifStmt.Else != nil {
			rowsToSkip = append(rowsToSkip, nv.FileSet.PositionFor(ifStmt.Else.Pos(), true).Line)
		}
	default:
		position := nv.FileSet.PositionFor(n.End(), true)
		rowsToSkip = append(rowsToSkip, position.Line)
	}

	nv.addRowsToSkip(rowsToSkip...)
}

func (nv *nesterVisitor) addRowsToSkip(positions ...int) {
	for _, position := range positions {
		nv.SkipRow[position] = true
	}
}

func (nv *nesterVisitor) findRowsToSkip(n ast.Node) []int {
	rowsToSkip := make([]int, 0)
	if n == nil {
		return rowsToSkip
	}
	startPosition := nv.FileSet.PositionFor(n.Pos(), true)
	endPosition := nv.FileSet.PositionFor(n.End(), true)
	for i := startPosition.Line; i <= endPosition.Line; i++ {
		rowsToSkip = append(rowsToSkip, i)
	}

	return rowsToSkip
}

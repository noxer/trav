package trav

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// escape is used to detect if the panic was intended to escape the traversal.
type escape struct{}

// TraverseSource traverses the AST of a Go source string, calling f with a
// slice containing the full path of nodes leading to the current node. The
// last element is the current node.
func TraverseSource(src string, f func(Path)) error {

	defer func() {

		r := recover()
		if r == nil {
			return
		}
		if _, ok := r.(escape); ok {
			return
		}
		panic(r)

	}()

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "trav", src, parser.ParseComments|parser.AllErrors)
	if err != nil {
		return err
	}

	TraverseNode(file, f)

	return nil

}

// TraverseNode traverses the subtree of a ast.Node, calling f with a
// slice containing the full path of nodes leading to the current node. The
// last element is the current node.
func TraverseNode(node ast.Node, f func(path Path)) {

	path := make([]ast.Node, 0, 100) // plenty of room
	ast.Inspect(node, func(n ast.Node) bool {

		if n == nil {

			path = path[:len(path)-1]
			return true

		}

		path = append(path, n)

		f(path)

		return true

	})

}

// Path represents a path of nodes in the AST.
type Path []ast.Node

// Current returns the last element in the path, which is the current node.
func (p Path) Current() ast.Node {

	if len(p) == 0 {
		return nil
	}

	return p[len(p)-1]

}

// Copy creates a copy of the path which can safely be used beyond the scope of
// the traversal function.
func (p Path) Copy() Path {

	o := make(Path, len(p))
	copy(o, p)

	return o

}

// Break terminates the traversal of the AST. Must not be called outside a
// traversal function.
func Break() {
	panic(escape{})
}

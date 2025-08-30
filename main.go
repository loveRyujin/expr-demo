package main

import (
	"fmt"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/ast"
)

type Visitor struct {
	Identifiers []string
}

func (v *Visitor) Visit(node *ast.Node) {
	if n, ok := (*node).(*ast.IdentifierNode); ok {
		v.Identifiers = append(v.Identifiers, n.Value)
	}
}

func main() {
	program, err := expr.Compile(`foo + bar`)
	if err != nil {
		panic(err)
	}

	node := program.Node()

	v := &Visitor{}
	ast.Walk(&node, v)
	fmt.Println(v)
}

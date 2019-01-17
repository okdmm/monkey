package evaluator

import (
	"github.com/okdmm/monkey/ast"
	"github.com/okdmm/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}

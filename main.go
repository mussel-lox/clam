package main

import (
	"encoding/json"
	"fmt"

	"github.com/mussel-lox/clam/ast"
)

func main() {
	expr := ast.BinaryExpression{
		Left: &ast.BinaryExpression{
			Left:     ast.NumberLiteral(114),
			Operator: ast.BinopAdd,
			Right:    ast.NumberLiteral(514),
		},
		Operator: ast.BinopDivide,
		Right:    ast.NumberLiteral(1919810),
	}

	serialized, err := json.Marshal(expr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(serialized))
}

package ast

type Expression interface {
	Accept(visitor ExpressionVisitor)
}

type ExpressionVisitor interface{}

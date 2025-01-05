//revive:disable

package ast

const (
	UopNegate UnaryOperator = iota
	UopLogicalNot
)

const (
	BinopLogicalOr BinaryOperator = iota
	BinopLogicalAnd
	BinopEqual
	BinopNotEqual
	BinopGreater
	BinopGreaterEqual
	BinopLess
	BinopLessEqual
	BinopAdd
	BinopSubtract
	BinopMultiply
	BinopDivide
)

type UnaryOperator byte
type BinaryOperator byte

type Expression interface {
	Accept(visitor ExpressionVisitor)
}

type ExpressionVisitor interface {
	VisitAssignment(a *AssignmentExpression)
	VisitBinary(b *BinaryExpression)
	VisitUnary(u *UnaryExpression)
	VisitInvocation(i *InvocationExpression)
	VisitPropertyAccess(p *PropertyAccessExpression)

	VisitBooleanLiteral(b BooleanLiteral)
	VisitNil(n Nil)
	VisitThis(t This)
	VisitNumberLiteral(n NumberLiteral)
	VisitStringLiteral(s StringLiteral)
	VisitIdentifier(i Identifier)
	VisitSuper(s Super)
}

type AssignmentExpression struct {
	Target Expression
	Value  Expression
}

type BinaryExpression struct {
	Left     Expression
	Right    Expression
	Operator BinaryOperator
}

type UnaryExpression struct {
	Operand  Expression
	Operator UnaryOperator
}

type InvocationExpression struct {
	Callee    Expression
	Arguments []Expression
}

type PropertyAccessExpression struct {
	Target   Expression
	Property Identifier
}

// Primary expressions

type BooleanLiteral bool
type Nil struct{}
type This struct{}
type NumberLiteral float64
type StringLiteral string
type Identifier string
type Super struct{}

func (a *AssignmentExpression) Accept(visitor ExpressionVisitor)     { visitor.VisitAssignment(a) }
func (b *BinaryExpression) Accept(visitor ExpressionVisitor)         { visitor.VisitBinary(b) }
func (u *UnaryExpression) Accept(visitor ExpressionVisitor)          { visitor.VisitUnary(u) }
func (i *InvocationExpression) Accept(visitor ExpressionVisitor)     { visitor.VisitInvocation(i) }
func (p *PropertyAccessExpression) Accept(visitor ExpressionVisitor) { visitor.VisitPropertyAccess(p) }
func (b BooleanLiteral) Accept(visitor ExpressionVisitor)            { visitor.VisitBooleanLiteral(b) }
func (n Nil) Accept(visitor ExpressionVisitor)                       { visitor.VisitNil(n) }
func (t This) Accept(visitor ExpressionVisitor)                      { visitor.VisitThis(t) }
func (s StringLiteral) Accept(visitor ExpressionVisitor)             { visitor.VisitStringLiteral(s) }
func (n NumberLiteral) Accept(visitor ExpressionVisitor)             { visitor.VisitNumberLiteral(n) }
func (i Identifier) Accept(visitor ExpressionVisitor)                { visitor.VisitIdentifier(i) }
func (s Super) Accept(visitor ExpressionVisitor)                     { visitor.VisitSuper(s) }

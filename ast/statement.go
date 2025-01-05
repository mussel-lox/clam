//revive:disable

package ast

type Statement interface {
	Accept(StatementVisitor)
}

type StatementVisitor interface {
	VisitExpressionStatement(*ExpressionStatement)
	VisitFor(*ForStatement)
	VisitIf(*IfStatement)
	VisitPrint(*PrintStatement)
	VisitReturn(*ReturnStatement)
	VisitWhile(*WhileStatement)
	VisitBlock(*BlockStatement)
}

type ExpressionStatement struct {
	Expression Expression
}

type ForStatement struct {
	ExpressionInitializer Expression
	VarInitializer        *VarDeclaration
	Condition             Expression
	Increment             Expression
	Body                  Statement
}

type IfStatement struct {
	Condition Expression
	Then      Statement
	Otherwise Statement
}

type PrintStatement struct {
	Expression Expression
}

type ReturnStatement struct {
	Expression Expression
}

type WhileStatement struct {
	Condition Expression
	Body      Statement
}

type BlockStatement struct {
	Declarations []Declaration
}

func (es *ExpressionStatement) Accept(visitor StatementVisitor) { visitor.VisitExpressionStatement(es) }
func (f *ForStatement) Accept(visitor StatementVisitor)         { visitor.VisitFor(f) }
func (i *IfStatement) Accept(visitor StatementVisitor)          { visitor.VisitIf(i) }
func (p *PrintStatement) Accept(visitor StatementVisitor)       { visitor.VisitPrint(p) }
func (r *ReturnStatement) Accept(visitor StatementVisitor)      { visitor.VisitReturn(r) }
func (w *WhileStatement) Accept(visitor StatementVisitor)       { visitor.VisitWhile(w) }
func (b *BlockStatement) Accept(visitor StatementVisitor)       { visitor.VisitBlock(b) }

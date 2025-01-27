package ast

type Declaration interface {
	Accept(DeclarationVisitor)
}

type DeclarationVisitor interface {
	VisitStatementDeclaration(*StatementDeclaration)
	VisitClass(*ClassDeclaration)
	VisitFun(*FunDeclaration)
	VisitVar(*VarDeclaration)
}

type StatementDeclaration struct {
	Statement Statement
}

type ClassDeclaration struct {
	Name      Identifier
	Baseclass *Identifier
	Methods   []FunDeclaration
}

type FunDeclaration struct {
	Name       Identifier
	Parameters []Identifier
	Body       *BlockStatement
}

type VarDeclaration struct {
	Name        Identifier
	Initializer Expression
}

func (s *StatementDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitStatementDeclaration(s)
}
func (c *ClassDeclaration) Accept(visitor DeclarationVisitor) { visitor.VisitClass(c) }
func (f *FunDeclaration) Accept(visitor DeclarationVisitor)   { visitor.VisitFun(f) }
func (v *VarDeclaration) Accept(visitor DeclarationVisitor)   { visitor.VisitVar(v) }

package ast

type Token struct {
	Token   int
	Literal string
}

// Statment
type Stmnt interface {
	Stmnt()
}

// Statment Implement
type StmntImpl struct {
	Expr Expr
}

func (si *StmntImpl) Stmnt() {}

// Expression
type Expr interface {
	Expr()
}

// Expression Implement
type ExprImpl struct {
	StmntImpl
}

func (ei *ExprImpl) Expr() {}

type NumExpr struct {
	ExprImpl
	Literal string
}

type IdenExpr struct {
	ExprImpl
	Literal string
}

type BinOpExpr struct {
	ExprImpl
	Left  Expr
	Op    rune
	Right Expr
}

func NewStmntByExpr(e Expr) Stmnt {
	return &StmntImpl{Expr: e}
}

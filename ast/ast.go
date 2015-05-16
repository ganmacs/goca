package ast

type Expression interface{}

type Token struct {
	Token   int
	Literal string
}

type NumExpr struct {
	Literal string
}

type IdenExpr struct {
	Literal string
}

type BinOpExpr struct {
	Left  Expression
	Op    rune
	Right Expression
}

package main

type Expression interface{}

type Token struct {
	token   int
	literal string
}

type NumExpr struct {
	literal string
}

type IdenExpr struct {
	literal string
}

type BinOpExpr struct {
	left  Expression
	op    rune
	right Expression
}

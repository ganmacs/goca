package parser

import (
	"github.com/ganmacs/goca/ast"
	"text/scanner"
)

type Lexer struct {
	scanner.Scanner
	stmnts []ast.Stmnt
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Int {
		token = NUMBER
	}
	if token == scanner.Ident {
		token = IDENT
	}
	lval.token = ast.Token{Token: token, Literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}

func Parse(l *Lexer) []ast.Stmnt {
	yyParse(l)
	return l.stmnts
}

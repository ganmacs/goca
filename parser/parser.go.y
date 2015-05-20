%{
package parser

import (
	"github.com/ganmacs/goca/ast"
)
%}

%type<stmnts> stmnts
%type<stmnt> stmnt
%type<expr> expr

%union{
	token ast.Token
	stmnts []ast.Stmnt
	stmnt ast.Stmnt
	expr ast.Expr
}

%token<token> IDENT NUMBER

%left '+', '-'
%left '*', '/'
%right '='
%%

stmnts
	: {                          // empty
	$$ = nil
		if l, ok := yylex.(*Lexer); ok {
		l.stmnts = $$
		}
	}
	| stmnt stmnts
	{
		$$ = append([]ast.Stmnt{$1}, $2...)
		if l, ok := yylex.(*Lexer); ok {
		l.stmnts = $$
		}
	}

stmnt
	: expr ';'
	{
		$$ = ast.NewStmntByExpr($1)
	}
	| IDENT '=' expr ';'
	{
		var ident ast.Expr = &ast.IdenExpr{Literal: $1.Literal}
		$$ = ast.NewStmntByExpr(&ast.BinOpExpr{Left: ident, Op: '=', Right: $3})
	}
	expr
	: NUMBER
	{
	$$ = &ast.NumExpr{Literal: $1.Literal}
	}
	| IDENT
	{
		$$ = &ast.IdenExpr{Literal: $1.Literal}
	}
	| expr '+' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Op: '+',  Right: $3}
	}
	| expr '-' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Op: '-',  Right: $3}
	}
	| expr '*' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Op: '*',  Right: $3}
	}
	| expr '/' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Op: '/',  Right: $3}
	}
%%

%{
package parser

import (
	"github.com/ganmacs/goca/ast"
)
%}

%type<stmtns> stmnts
%type<expr> stmnt
%type<expr> expr

%union{
	token ast.Token
	expr ast.Expression
	stmtns []ast.Expression
}

%token<token> IDENT NUMBER

%left '+', '-'
%left '*', '/'
%right '='
%%

stmnts : {
    $$ = nil
		if l, ok := yylex.(*Lexer); ok {
      l.result = $$
		}
  }
  | stmnt stmnts
  {
    $$ = append([]ast.Expression{$1}, $2...)
		if l, ok := yylex.(*Lexer); ok {
      l.result = $$
		}
  }

stmnt
  : expr ';' {
      $$ = $1
  }
  | IDENT '=' expr ';' {
      ident := ast.IdenExpr{Literal: $1.Literal}
      $$ = ast.BinOpExpr{Left: ident, Op: '=', Right: $3}
  }

expr
  : NUMBER {
      $$ = ast.NumExpr{Literal: $1.Literal}
  }
  | IDENT {
      $$ = ast.IdenExpr{Literal: $1.Literal}
  }
  | expr '+' expr {
      $$ = ast.BinOpExpr{Left: $1, Op: '+',  Right: $3}
    }
  | expr '-' expr {
      $$ = ast.BinOpExpr{Left: $1, Op: '-',  Right: $3}
    }
  | expr '*' expr {
      $$ = ast.BinOpExpr{Left: $1, Op: '*',  Right: $3}
    }
  | expr '/' expr {
      $$ = ast.BinOpExpr{Left: $1, Op: '/',  Right: $3}
    }
%%

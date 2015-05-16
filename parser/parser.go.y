%{
package parser

import (
	"github.com/ganmacs/goca/ast"
)
%}

%type<expr> program
%type<expr> expr
%type<expr> stmnt

%union{
	token ast.Token
	expr ast.Expression
}

%token<token> IDENT NUMBER

%left '+', '-'
%left '*', '/'
%right '='
%%

program
  : stmnt {
      $$ = $1
      yylex.(*Lexer).result = $$
  }
stmnt
  : expr
  | IDENT '=' expr {
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

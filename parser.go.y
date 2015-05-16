%{
package main
%}

%type<expr> program
%type<expr> expr
%type<expr> stmnt

%union{
	token Token
	expr Expression
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
      ident := IdenExpr{literal: $1.literal}
      $$ = BinOpExpr{left: ident, op: '=', right: $3}
  }
expr
  : NUMBER {
      $$ = NumExpr{literal: $1.literal}
  }
  | IDENT {
      $$ = IdenExpr{literal: $1.literal}
  }
  | expr '+' expr {
      $$ = BinOpExpr{left: $1, op: '+',  right: $3}
    }
  | expr '-' expr {
      $$ = BinOpExpr{left: $1, op: '-',  right: $3}
    }
  | expr '*' expr {
      $$ = BinOpExpr{left: $1, op: '*',  right: $3}
    }
  | expr '/' expr {
      $$ = BinOpExpr{left: $1, op: '/',  right: $3}
    }
%%

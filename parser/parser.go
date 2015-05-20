//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/ganmacs/goca/ast"
)

//line parser.go.y:13
type yySymType struct {
	yys    int
	token  ast.Token
	expr   ast.Expression
	stmtns []ast.Expression
}

const IDENT = 57346
const NUMBER = 57347

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'='",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:68

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 11
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 29

var yyAct = []int{

	3, 8, 9, 10, 11, 2, 19, 10, 11, 13,
	15, 16, 17, 18, 8, 9, 10, 11, 0, 7,
	12, 14, 5, 4, 5, 1, 0, 0, 6,
}
var yyPact = []int{

	19, -1000, 19, 8, 10, -1000, -1000, -1000, 17, 17,
	17, 17, 17, -1, -1000, -1, -1000, -1000, -5, -1000,
}
var yyPgo = []int{

	0, 25, 5, 0,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3,
}
var yyR2 = []int{

	0, 0, 2, 2, 4, 1, 1, 3, 3, 3,
	3,
}
var yyChk = []int{

	-1000, -1, -2, -3, 4, 5, -1, 11, 6, 7,
	8, 9, 10, -3, 4, -3, -3, -3, -3, 11,
}
var yyDef = []int{

	1, -2, 1, 0, 6, 5, 2, 3, 0, 0,
	0, 0, 0, 7, 6, 8, 9, 10, 0, 4,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 8, 6, 3, 7, 3, 9, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 11,
	3, 10,
}
var yyTok2 = []int{

	2, 3, 4, 5,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:26
		{
			yyVAL.stmtns = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmnts = yyVAL.stmtns
			}
		}
	case 2:
		//line parser.go.y:33
		{
			yyVAL.stmtns = append([]ast.Expression{yyS[yypt-1].expr}, yyS[yypt-0].stmtns...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmnts = yyVAL.stmtns
			}
		}
	case 3:
		//line parser.go.y:41
		{
			yyVAL.expr = yyS[yypt-1].expr
		}
	case 4:
		//line parser.go.y:44
		{
			ident := ast.IdenExpr{Literal: yyS[yypt-3].token.Literal}
			yyVAL.expr = ast.BinOpExpr{Left: ident, Op: '=', Right: yyS[yypt-1].expr}
		}
	case 5:
		//line parser.go.y:50
		{
			yyVAL.expr = ast.NumExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 6:
		//line parser.go.y:53
		{
			yyVAL.expr = ast.IdenExpr{Literal: yyS[yypt-0].token.Literal}
		}
	case 7:
		//line parser.go.y:56
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Op: '+', Right: yyS[yypt-0].expr}
		}
	case 8:
		//line parser.go.y:59
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Op: '-', Right: yyS[yypt-0].expr}
		}
	case 9:
		//line parser.go.y:62
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Op: '*', Right: yyS[yypt-0].expr}
		}
	case 10:
		//line parser.go.y:65
		{
			yyVAL.expr = ast.BinOpExpr{Left: yyS[yypt-2].expr, Op: '/', Right: yyS[yypt-0].expr}
		}
	}
	goto yystack /* stack new state and value */
}

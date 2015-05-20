package vm

import (
	"fmt"
	"github.com/ganmacs/goca/ast"
	"reflect"
	"strconv"
)

var NilValue = reflect.ValueOf((interface{})(nil))

func Run(stmnts []ast.Stmnt, e *Env) {
	for _, stmns := range stmnts {
		v, err := runSingleStmnt(stmns, e)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v.Int())
	}
}

func runSingleStmnt(stmnt ast.Stmnt, e *Env) (reflect.Value, error) {
	rv := NilValue
	rv, err := invokeExpr(stmnt.(*ast.StmntImpl).Expr, e)
	if err != nil {
		fmt.Println("error")
	}
	return rv, nil
}

func invokeExpr(expr ast.Expr, env *Env) (reflect.Value, error) {
	rv := NilValue
	var err error

	switch e := expr.(type) {
	case *ast.BinOpExpr:
		if e.Op == '=' {
			rv = assignValue(*e, env)
		} else {
			rv = computeValue(*e, env)
		}
	case *ast.NumExpr:
		i, err := toInt(*e)
		if err != nil {
			panic(err)
		}
		rv = reflect.ValueOf(i)
	case *ast.IdenExpr:
		rv, err = env.Get(e.Literal)
		if err != nil {
			fmt.Printf("Not found such keyword %s", rv)
		}
	}
	return rv, nil
}

func toInt(i ast.NumExpr) (int64, error) {
	return strconv.ParseInt(i.Literal, 10, 64)
}

func computeValue(expr ast.BinOpExpr, e *Env) reflect.Value {
	rv := NilValue

	left, err := invokeExpr(expr.Left, e)
	if err != nil {
		panic(err)
	}

	right, err := invokeExpr(expr.Right, e)
	if err != nil {
		panic(err)
	}
	switch expr.Op {
	case '+':
		rv = reflect.ValueOf(left.Int() + right.Int())
	case '-':
		rv = reflect.ValueOf(left.Int() - right.Int())
	case '/':
		rv = reflect.ValueOf(left.Int() / right.Int())
	case '*':
		rv = reflect.ValueOf(left.Int() * right.Int())
	default:
		fmt.Println("default")
	}
	return rv
}

func assignValue(expr ast.BinOpExpr, e *Env) reflect.Value {
	right, err := invokeExpr(expr.Right, e)
	if err != nil {
		panic(err)
	}

	e.Set(expr.Left.(*ast.IdenExpr).Literal, right)
	return right
}

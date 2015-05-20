package vm

import (
	"fmt"
	"github.com/ganmacs/goca/ast"
	"reflect"
	"strconv"
)

var NilValue = reflect.ValueOf((interface{})(nil))

func Run(exprs []ast.Expression, e *Env) {
	for _, expr := range exprs {

		v, err := runSingleExpr(expr, e)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v.Int())
	}
}

func runSingleExpr(expr ast.Expression, e *Env) (reflect.Value, error) {
	rv := NilValue

	switch expr := expr.(type) {
	case ast.BinOpExpr:
		if expr.Op == '=' {
			rv = assignValue(expr, e)
		} else {
			rv = computeValue(expr, e)
		}
	case ast.NumExpr:
		i, err := toInt(expr)
		if err != nil {
			panic(err)
		}
		rv = reflect.ValueOf(i)
	case ast.IdenExpr:
		_rv, err := e.Get(expr.Literal)
		rv = _rv
		if err != nil {
			fmt.Printf("Not found such keyword %s", rv)
		}
	default:
		fmt.Println("error")
	}

	return rv, nil
}

func toInt(i ast.NumExpr) (int64, error) {
	return strconv.ParseInt(i.Literal, 10, 64)
}

func computeValue(expr ast.BinOpExpr, e *Env) reflect.Value {
	rv := NilValue

	left, err := runSingleExpr(expr.Left, e)
	if err != nil {
		panic(err)
	}

	right, err := runSingleExpr(expr.Right, e)
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
	// rv := NilValue
	// rv = reflect.ValueOf(right)

	right, err := runSingleExpr(expr.Right, e)
	if err != nil {
		panic(err)
	}

	e.Set(expr.Left.(ast.IdenExpr).Literal, right)
	return right
}

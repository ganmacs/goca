package vm

import (
	"fmt"
	"github.com/ganmacs/goca/ast"
	"reflect"
	"strconv"
)

var NilValue = reflect.ValueOf((*interface{})(nil))

func Run(expr ast.Expression, e *Env) (reflect.Value, error) {
	rv := NilValue

	switch expr := expr.(type) {
	case ast.BinOpExpr:
		left, err := Run(expr.Left, e)
		if err != nil {
			panic(err)
		}

		right, err := Run(expr.Right, e)
		if err != nil {
			panic(err)
		}

		switch expr.Op {
		case '=':
			fmt.Println("=")
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

	case ast.NumExpr:
		var i int64
		i, err := strconv.ParseInt(expr.Literal, 10, 64)
		if err != nil {
			panic(err)
		}
		rv = reflect.ValueOf(i)
	case ast.IdenExpr:
		rv = reflect.ValueOf(expr.Literal)
	default:
		fmt.Println("error")
	}
	return rv, nil
}

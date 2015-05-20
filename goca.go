package main

import (
	// "fmt"
	"github.com/ganmacs/goca/parser"
	"github.com/ganmacs/goca/vm"
	"os"
	"strings"
)

func main() {
	l := new(parser.Lexer)
	l.Init(strings.NewReader(os.Args[1]))
	exprs := parser.Parse(l)

	env := vm.NewEnv()
	vm.Run(exprs, env)
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// fmt.Println(ret)
}

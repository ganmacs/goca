package main

import (
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
}

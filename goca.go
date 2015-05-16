package main

import (
	"fmt"
	"github.com/ganmacs/goca/parser"
	"github.com/ganmacs/goca/vm"
	"os"
	"strings"
)

func main() {
	l := new(parser.Lexer)
	l.Init(strings.NewReader(os.Args[1]))
	expr := parser.Parse(l)

	env := vm.NewEnv()
	ret, err := vm.Run(expr, env)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Int())
}

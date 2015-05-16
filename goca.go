package main

import (
	"fmt"
	"github.com/ganmacs/goca/parser"
	"os"
	"strings"
)

func main() {
	l := new(parser.Lexer)
	l.Init(strings.NewReader(os.Args[1]))
	ret := parser.Parse(l)
	fmt.Printf("%#v\n", ret)
}

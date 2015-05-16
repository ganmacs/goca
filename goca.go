package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	l := new(Lexer)
	l.Init(strings.NewReader(os.Args[1]))
	yyParse(l)
	fmt.Printf("%#v\n", l.result)

	// if len(os.Args) != 2 {
	// 	fmt.Fprintf(os.Stderr, "usage of %s: [file]\n", os.Args[0])
	// 	os.Exit(1)
	// }

	// b, err := ioutil.ReadFile(os.Args[1])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error raise while reading %s\n", os.Args[1])
	// 	os.Exit(1)
	// }
	// scanner := new(Scanner)
	// scanner.Int(string(b))
	// stmts, err := Parse(scanner)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error raise while reading %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(stmts)
}

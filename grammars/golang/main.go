package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	astFlag = flag.Bool("ast", false, "display AST")
)

func main() {
	flag.Parse()
	buffer, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	s := string(buffer)
	golang := &goGrammar{Buffer: s, Pretty: true}
	golang.Init()
	if err := golang.Parse(); err != nil {
		log.Fatal(err)
	}
	if *astFlag {
		golang.PrettyPrintSyntaxTree(s)
	} else {
		fmt.Println("total tokens:", len(golang.Tokens()))
	}
}

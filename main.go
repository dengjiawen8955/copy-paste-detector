package main

import (
	"copy-paste-detector/parser"
	"fmt"
)

func main() {
	conf := &parser.Conf{
		MinRepeatLine: 0,
		ParseFolder:   "../copy-paste-detector",
	}
	d := parser.New(conf)
	r := d.Parse()
	fmt.Printf("len(r.IDLines): %v\n", len(r.IDLines))
}

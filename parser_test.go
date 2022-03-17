package main

import (
	"fmt"
	"testing"
)

func TestProgram(t *testing.T) {
	p := Parser{}
	ast := p.Parse(`123"ABC"`)
	fmt.Println(ast)
}

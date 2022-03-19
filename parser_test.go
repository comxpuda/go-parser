package main

import (
	"fmt"
	"testing"
)

func TestProgram(t *testing.T) {
	p := Parser{}
	ast := p.Parse(`
	42;
	"hello";
	{
		45;
	}
	`)
	fmt.Println(ast)
}

func TestEmptyStatement(t *testing.T) {
	p := Parser{}
	ast := p.Parse(`
	;
	`)
	fmt.Println(ast)
}

func TestAdditiveExpression(t *testing.T) {
	p := Parser{}
	ast := p.Parse(`
	2 + 2;
	3 + 3 + 1;
	`)
	fmt.Println(ast)
}

func TestMultiExpression(t *testing.T) {
	p := Parser{}
	ast := p.Parse(`
	3 + 2 * 1;
	`)
	fmt.Println(ast)
}

func TestPrimaryExpression(t *testing.T) {
	p := Parser{}
	ast := p.Parse(`
	(3 + 2) * 1;
	`)
	fmt.Println(ast)
}

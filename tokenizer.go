package main

import (
	"fmt"
	"regexp"
)

var Spec [][2]string = [][2]string{
	{`^\s+`, "NULL"},
	{`^\/\/.*`, "NULL"},
	{`^\/\*[\s\S]*?\*\/`, "NULL"},
	{`^;`, ";"},
	{`^\{`, "{"},
	{`^\}`, "}"},

	{`^\(`, "("},
	{`^\)`, ")"},

	{`^[+\-]`, "ADDITIVE_OPERATOR"},
	{`^[*\/]`, "MULTIPLICATIVE_OPERATOR"},
	{`^\d+`, "NUMBER"},
	{`"[^"]*"`, "STRING"},
	{`'[^']*'`, "STRING"},
}

type Tokenizer struct {
	source string
	cursor int
}

type Token struct {
	TokenType  string
	TokenValue string
}

func (t *Tokenizer) isEOF() bool {
	return len(t.source) == t.cursor
}

func (t *Tokenizer) hasMoreTokens() bool {
	return t.cursor < len(t.source)
}

func (t *Tokenizer) match(regex, str string) string {
	r, _ := regexp.Compile(regex)
	matched := r.MatchString(str)

	if !matched {
		return ""
	}
	t.cursor += len(r.FindString(str))
	return r.FindString(str)
}

func (t *Tokenizer) getNextToken() (*Token, error) {
	if !t.hasMoreTokens() {
		return nil, nil
	}
	str := t.source[t.cursor:]

	for _, spec := range Spec {
		regex := spec[0]
		tokenType := spec[1]

		tokenValue := t.match(regex, str)
		if tokenValue == "" {
			continue
		}

		if tokenType == "NULL" {
			return t.getNextToken()
		}

		return &Token{
			TokenType:  tokenType,
			TokenValue: tokenValue,
		}, nil
	}
	return nil, fmt.Errorf("unexpected token:%s", str)
}

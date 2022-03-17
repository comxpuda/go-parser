package main

import "unicode"

// const Spec := []string{
// 	"",
// 	"",
// }

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

func (t *Tokenizer) getNextToken() *Token {
	if !t.hasMoreTokens() {
		return nil
	}
	str := t.source[t.cursor:]
	if unicode.IsNumber(rune(str[0])) {
		var number string
		for !t.isEOF() && unicode.IsNumber(rune(str[t.cursor])) {
			number += string(str[t.cursor])
			t.cursor++
		}
		return &Token{
			TokenType:  "NUMBER",
			TokenValue: number,
		}
	}

	if str[0] == '"' {
		var s string
		t.cursor++
		for str[t.cursor] != '"' && !t.isEOF() {
			s += string(str[t.cursor])
			t.cursor++
		}
		return &Token{
			TokenType:  "STRING",
			TokenValue: s,
		}
	}

	return nil
}

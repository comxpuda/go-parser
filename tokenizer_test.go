package main

import (
	"fmt"
	"testing"
)

func TestGetNextToken(t *testing.T) {
	tokenizer := Tokenizer{
		source: `"123"`,
	}

	token := tokenizer.getNextToken()
	fmt.Println(token)

}

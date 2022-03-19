package main

import (
	"fmt"
	"testing"
)

func TestGetNextToken(t *testing.T) {
	tokenizer := Tokenizer{
		source: `  
		/*
		* hell
		*/

		// NUMBER
		42;
		`,
	}

	token, err := tokenizer.getNextToken()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token)

}

func TestAdditive(t *testing.T) {
	tokenizer := Tokenizer{source: "2+3;"}
	token, _ := tokenizer.getNextToken()
	token, _ = tokenizer.getNextToken()
	token, _ = tokenizer.getNextToken()
	fmt.Println(token)
}

package algebra

import (
	"fmt"
	"regexp"
)

type Token struct {
	Name     string
	Contents string
}

type Lexer struct {
	source   string
	location int
}

func (lex *Lexer) lex() []Token {
	var stream []Token
	for !lex.End() {
		c := lex.source[lex.location]

		if lex.isNumber(c) {
			stream = append(stream, lex.number())
			continue
		}

		if lex.isBrackets(c) {
			stream = append(stream, lex.brackets())
			continue
		}

		if lex.isAddition(c) {
			stream = append(stream, lex.addition())
			continue
		}

		if lex.isWhitespace(c) {
			stream = append(stream, lex.whitespace())
			continue
		}

		fmt.Println("Invalid token!")
		lex.Next()
	}

	return stream
}

func (lex *Lexer) addition() Token {
	lex.Next()

	return Token{"Addition", "+"}
}

func (lex *Lexer) whitespace() Token {
	var content string
	for !lex.End() && lex.isWhitespace(lex.source[lex.location]) {
		content += string(lex.source[lex.location])
		lex.Next()
	}
	return Token{"Skip", content}
}

func (lex *Lexer) brackets() Token {
	c := lex.source[lex.location]

	lex.Next()
	if c == '(' {
		return Token{"RightBracket", "("}
	}

	return Token{"LeftBracket", ")"}
}

func (lex *Lexer) number() Token {
	var content string

	for !lex.End() && lex.isNumber(lex.source[lex.location]) {
		content += string(lex.source[lex.location])
		lex.Next()
	}

	return Token{"Number", content}
}

func (lex *Lexer) isAddition(c uint8) bool {
	if c == '+' {
		return true
	}
	return false
}

func (lex *Lexer) isWhitespace(c uint8) bool {
	res, err := regexp.Match(`\s`, []byte(string(c)))
	if err != nil {
		return false
	}
	return res
}

func (lex *Lexer) isBrackets(c uint8) bool {
	if c == '(' || c == ')' {
		return true
	}

	return false
}

func (lex *Lexer) isNumber(c uint8) bool {
	res, err := regexp.Match(`\d+`, []byte(string(c)))
	if err != nil {
		return false
	}
	return res
}

func (lex *Lexer) End() bool {
	if lex.location >= len(lex.source) {
		return true
	}

	return false
}

func (lex *Lexer) Next() {
	lex.location += 1
}

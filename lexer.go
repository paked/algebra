package algebra

import (
	"fmt"
	"regexp"
)

const (
	NumberToken         = "Number"
	SkipToken           = "Skip"
	LeftBracketToken    = "LeftBracket"
	RightBracketToken   = "RightBracket"
	AdditionToken       = "Addition"
	SubtractionToken    = "Subtraction"
	MultiplicationToken = "Multiplication"
)

type Token struct {
	Name     string
	Contents string
}

func (t Token) Type(n string) bool {
	if n == t.Name {
		return true
	}

	return false
}

type Lexer struct {
	source   string
	location int
}

func (lex *Lexer) Lex() []Token {
	return lex.lex()
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

		if lex.isSubtraction(c) {
			stream = append(stream, lex.subtraction())
			continue
		}

		if lex.isMultiplication(c) {
			stream = append(stream, lex.multiplication())
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

	return Token{AdditionToken, "+"}
}

func (lex *Lexer) subtraction() Token {
	lex.Next()
	return Token{SubtractionToken, "-"}
}

func (lex *Lexer) multiplication() Token {
	lex.Next()
	return Token{MultiplicationToken, "*"}
}

func (lex *Lexer) whitespace() Token {
	var content string
	for !lex.End() && lex.isWhitespace(lex.source[lex.location]) {
		content += string(lex.source[lex.location])
		lex.Next()
	}
	return Token{SkipToken, content}
}

func (lex *Lexer) brackets() Token {
	c := lex.source[lex.location]

	lex.Next()
	if c == '(' {
		return Token{RightBracketToken, "("}
	}

	return Token{LeftBracketToken, ")"}
}

func (lex *Lexer) number() Token {
	var content string

	for !lex.End() && lex.isNumber(lex.source[lex.location]) {
		content += string(lex.source[lex.location])
		lex.Next()
	}

	return Token{NumberToken, content}
}

func (lex *Lexer) isSubtraction(c uint8) bool {
	if c == '-' {
		return true
	}
	return false
}

func (lex *Lexer) isAddition(c uint8) bool {
	if c == '+' {
		return true
	}
	return false
}

func (lex *Lexer) isMultiplication(c uint8) bool {
	if c == '*' {
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

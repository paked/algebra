package algebra

import "regexp"

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
		}

		lex.Next()
	}

	return stream
}

func (lex *Lexer) number() Token {
	var content string

	for !lex.End() && lex.isNumber(lex.source[lex.location]) {
		content += string(lex.source[lex.location])
		lex.Next()
	}

	return Token{"Number", content}
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

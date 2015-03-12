package algebra

import (
	"errors"
)

type Parser struct {
	tokens   []Token
	location int
}

func (p *Parser) Parse() {

}

func (p *Parser) End() bool {
	if p.location >= len(p.tokens) {
		return false
	}
	return true
}

func (p *Parser) Peek() (Token, error) {
	if p.location+1 > len(p.tokens) {
		return errors.New("End of tokens...")
	}
	return p.tokens[p.location], nil
}

func (p *Parser) Next() {
	p.location += 1
}

type Node interface {
	Eval() int
}

type NumberNode struct {
	Number int
}

func (n NumberNode) Eval() int {
	return n.Number
}

type AdditionNode struct {
	Left, Right Node
}

func (a AdditionNode) Eval() int {
	return a.Left.Eval() + a.Right.Eval()
}

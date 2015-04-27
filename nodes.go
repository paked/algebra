package algebra

import (
	"math"
)

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

type SubtractionNode struct {
	Left, Right Node
}

func (s SubtractionNode) Eval() int {
	return s.Left.Eval() - s.Right.Eval()
}

type MultiplicationNode struct {
	Left, Right Node
}

func (m MultiplicationNode) Eval() int {
	return m.Left.Eval() * m.Right.Eval()
}

type DivisionNode struct {
	Left, Right Node
}

func (d DivisionNode) Eval() int {
	return d.Left.Eval() / d.Right.Eval()
}

type PowerNode struct {
	Left, Right Node
}

func (pn PowerNode) Eval() int {
	return int(math.Pow(float64(pn.Left.Eval()), float64(pn.Right.Eval())))
}

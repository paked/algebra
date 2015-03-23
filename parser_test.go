package algebra

import (
	"testing"
)

func TestNumbers(t *testing.T) {
	testSum(t, "1", 1)
	testSum(t, "-1", -1)
}

func TestAddition(t *testing.T) {
	testSum(t, "1 + 1", 2)
	testSum(t, "100 + 1", 101)
}

func TestSubtraction(t *testing.T) {
	testSum(t, "1 - 1", 0)
	testSum(t, "100 - 1", 99)
}

func TestMultiplication(t *testing.T) {
	testSum(t, "1 * 1", 1)
	testSum(t, "10 * 10", 100)
	testSum(t, "-10 * 10", -100)
}

func TestBrackets(t *testing.T) {
	testSum(t, "(2 + 2) * 2", 8)
	testSum(t, "(2 - 2) * 1000", 0)
}

func testSum(t *testing.T, source string, expected int) {
	if got := Evaluate(source); got != expected {
		t.Errorf("number error: %v should equal %d, got %d", source, expected, got)
	}
}

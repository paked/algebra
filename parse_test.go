package algebra

import (
	"testing"
)

func TestParse(t *testing.T) {
	println("==EVALUTAING==")

	Evaluate("1 + 1 + 100")
	Evaluate("1 - 1")
	Evaluate("10 + 10 - 5")

	Evaluate("10 * 10 + 10")
}

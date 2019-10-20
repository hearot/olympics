package olympics

import (
	"fmt"
	"testing"
)

func ExampleAreCongruent() {
	fmt.Print(AreCongruent(25, 35, 10)) // Check if 25 ≡ 35 (mod 10)
	// Output: true
}

func TestAreCongruent(t *testing.T) {
	if !AreCongruent(21, 31, 10) {
		t.Errorf("AreCongruent(21, 31, 10) returned false instead of true")
	}
}

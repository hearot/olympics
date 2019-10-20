package olympics

import (
	"fmt"
	"testing"
)

const result = 49376

func ExamplePow() {
	fmt.Print(Pow(2, 50, 10)) // Let's calculate 2^50 mod 10
	// Output: 4
}

func ExamplePow_sum() {
	// Let's calculate the first four digits of the
	// sum of 2^i^3 as i goes from 1 to 1183.
	sum := 0

	for i := 1; i <= 1183; i++ {
		sum += Pow(2, i*i*i, 10000)
		sum %= 10000
	}

	fmt.Print(sum)
	// Output: 1426
}

func TestPow(t *testing.T) {
	if res := Pow(2, 4000, 100000); res != result {
		t.Errorf("Pow(2, 4000, 100000) should return %d, but returned %d", result, res)
	}
}

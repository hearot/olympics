package olympics

import (
	"fmt"
	"testing"
)

const (
	evensResult = 10100
	gaussResult = 5050
	oddsResult = 10000
)

func ExampleEvensSummation() {
	fmt.Print(EvensSummation(15)) // Let's calculate the sum of the first 15 even numbers
	// Output: 240
}

func ExampleGaussSummation() {
	fmt.Print(GaussSummation(50)) // Let's calculate the sum of the first 50 integers.
	// Output: 1275
}

func ExampleOddsSummation() {
	fmt.Print(OddsSummation(15)) // Let's calculate the sum of the first 15 odd integers.
	// Output: 225
}

func TestEvensSummation(t *testing.T) {
	if sum := EvensSummation(100); sum != gaussResult {
		t.Errorf("EvensSummation(100) should have returned %d, not %d", evensResult, sum)
	}
}

func TestGaussSummation(t *testing.T) {
	if sum := GaussSummation(100); sum != gaussResult {
		t.Errorf("GaussSummation(100) should have returned %d, not %d", gaussResult, sum)
	}
}

func TestOddsSummation(t *testing.T) {
	if sum := OddsSummation(100); sum != oddsResult {
		t.Errorf("OddsSummation(100) should have returned %d, not %d", oddsResult, sum)
	}
}

package olympics

import (
	"fmt"
	"testing"
)

const (
	evensResult                 = 10100
	gaussResult                 = 5050
	leftRiemannSumResult        = 10
	oddsResult                  = 10000
	rightRiemannSumResult       = 15
	simpsonSumResult            = 8
	trapezoidalRiemannSumResult = 12.5
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
	if sum := EvensSummation(100); sum != evensResult {
		t.Errorf("EvensSummation(100) should have returned %d, not %d", evensResult, sum)
	}
}

func TestGaussSummation(t *testing.T) {
	if sum := GaussSummation(100); sum != gaussResult {
		t.Errorf("GaussSummation(100) should have returned %d, not %d", gaussResult, sum)
	}
}

func TestLeftRiemannSum(t *testing.T) {
	if sum := LeftRiemannSum(func(x float64) float64 { return x }, 0, 5, 5); sum != leftRiemannSumResult {
		t.Errorf("LeftRiemannSum(func(x float64)float64{return x}, 0, 5, 5) should have returned %d, not %f", leftRiemannSumResult, sum)
	}
}

func TestOddsSummation(t *testing.T) {
	if sum := OddsSummation(100); sum != oddsResult {
		t.Errorf("OddsSummation(100) should have returned %d, not %d", oddsResult, sum)
	}
}

func TestRightRiemannSum(t *testing.T) {
	if sum := RightRiemannSum(func(x float64) float64 { return x }, 0, 5, 5); sum != rightRiemannSumResult {
		t.Errorf("RightRiemannSum(func(x float64)float64{return x}, 0, 5, 5) should have returned %d, not %f", rightRiemannSumResult, sum)
	}
}

func TestSimpsonSum(t *testing.T) {
	if sum := SimpsonSum(func(x float64) float64 { return x }, 0, 4, 50); sum != simpsonSumResult {
		t.Errorf("SimpsonSum(func(x float64)float64{return x}, 0, 4, 50) should have returned %d, not %f", simpsonSumResult, sum)
	}
}

func TestTrapezoidalRiemannSum(t *testing.T) {
	if sum := TrapezoidalRiemannSum(func(x float64) float64 { return x }, 0, 5, 5); sum != trapezoidalRiemannSumResult {
		t.Errorf("TrapezoidalRiemannSum(func(x float64)float64{return x}, 0, 5, 5) should have returned %f, not %f", trapezoidalRiemannSumResult, sum)
	}
}

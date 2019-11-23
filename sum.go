package olympics

// GaussSummation returns the sum of first
// n integers. Not to be confused with
// Gauss sum (https://en.wikipedia.org/wiki/Gauss_sum).
func GaussSummation(n uint) uint {
	return n * (n + 1) / 2
}

// EvensSummation is used in order to sum
// the first n even numbers.
func EvensSummation(n uint) uint {
	return n * (n + 1)
}

// LeftRiemannSum is used in order to calculate
// the approximation of an integral. You can get
// more information on Wikipedia
// (https://en.wikipedia.org/wiki/Riemann_sum).
func LeftRiemannSum(f func(x float64) float64, a float64, b float64, n int) (sum float64) {
	dx := (b - a) / float64(n)

	for i := 0; i <= (n - 1); i++ {
		sum += dx * f(a+float64(i)*dx)
	}

	return
}

// OddsSummation is used in order to sum
// the first n odd numbers.
func OddsSummation(n uint) uint {
	return n * n
}

// RightRiemannSum is used in order to calculate
// the approximation of an integral. You can get
// more information on Wikipedia
// (https://en.wikipedia.org/wiki/Riemann_sum).
func RightRiemannSum(f func(x float64) float64, a float64, b float64, n int) (sum float64) {
	dx := (b - a) / float64(n)

	for i := 1; i <= n; i++ {
		sum += dx * f(a+float64(i)*dx)
	}

	return
}

// SimpsonSum is used in order to calculate
// the approximation of an integral. You can get
// more information on Wikipedia
// (https://en.wikipedia.org/wiki/Simpson%27s_rule).
func SimpsonSum(f func(x float64) float64, a float64, b float64, n int) (sum float64) {
	dx := (b - a) / float64(n)

	sum += f(a)

	for i := 1; i < n; i += 2 {
		sum += 4 * f(a+float64(i)*dx)
	}

	for i := 2; i < n; i += 2 {
		sum += 2 * f(a+float64(i)*dx)
	}

	sum += f(b)

	return dx / 3 * sum
}

// RightRiemannSum is used in order to calculate
// the approximation of an integral using the
// trapezoidal rule. You can get more information on
// Wikipedia (https://en.wikipedia.org/wiki/Trapezoidal_rule).
func TrapezoidalRiemannSum(f func(x float64) float64, a float64, b float64, n int) (sum float64) {
	dx := (b - a) / float64(n)

	for i := 1; i <= n; i++ {
		sum += dx * (f(a+float64(i-1)*dx) + f(a+float64(i)*dx)) / 2
	}

	return
}

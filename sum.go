package olympics

// GaussSummation returns the sum of first
// n integers. Not to be confused with
// Gauss sum (https://en.wikipedia.org/wiki/Gauss_sum).
func GaussSummation(n uint) uint {
	return n*(n+1)/2
}

// EvensSummation is used in order to sum
// the first n even numbers.
func EvensSummation(n uint) uint {
	return n*(n+1)
}

// OddsSummation is used in order to sum
// the first n odd numbers.
func OddsSummation(n uint) uint {
	return n*n
}
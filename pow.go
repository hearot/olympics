package olympics

// Pow is the representation of a^b mod modulo.
// It makes use of the Modular exponentiation
// (https://en.wikipedia.org/wiki/Modular_exponentiation).
func Pow(a int, b int, modulo int) int {
	if b == 0 {
		return 1
	}

	res := Pow(a, b/2, modulo)
	res = (res * res) % modulo

	if b % 2 == 1 {
		res = (res * a) % modulo
	}

	return res
}

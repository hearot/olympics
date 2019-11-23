package olympics

// AreCongruent is used in order to know if
// two integers are congruent modulo mod.
func AreCongruent(a int, b int, mod int) bool {
	return (a-b)%mod == 0
}

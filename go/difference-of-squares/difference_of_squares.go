package diffsquares

func SquareOfSum(n int) int {
	ret := (n + 1) * n / 2
	return ret * ret
}

func SumOfSquares(n int) int {
	ret := (2*n*n + 3*n + 1) * n / 6
	return ret
}

func Difference(n int) int {
	ret := SquareOfSum(n) - SumOfSquares(n)
	return ret
}

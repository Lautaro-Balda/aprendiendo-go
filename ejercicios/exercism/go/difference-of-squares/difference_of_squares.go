package diffsquares

func SquareOfSum(n int) int {
	var result int
	for i := range n {
		result += i + 1
	}
	return result * result
}

func SumOfSquares(n int) int {
	var result int

	for i := range n {
		result += (i + 1) * (i + 1)
	}
	return result
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

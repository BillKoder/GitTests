package mathematics

// Sum finds the sum of an array of ints
func Sum(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}

// Multiply returns the product of the passed ints
func Multiply(nums []int) int {
	result := 1
	for _, num := range nums {
		result = result * num
	}
	return result
}

// Difference returns the amount of the second int subtracted from the first int
func Difference(x, y int) int {
	return x - y
}

// Divide returns the quotient of x divided by y
func Divide(x, y int) int {
	return x / y
}

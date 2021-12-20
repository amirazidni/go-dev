package list

// Add all of the numbers in an array
func ListSum(numList []float64) float64 {
	sum := 0.0
	for _, num := range numList {
		sum += num
	}

	return sum
}

// Multiply all of the numbers in an array
func ListMul(numList []float64) float64 {
	mul := numList[0]
	for _, num := range numList[1:] {
		mul *= num
	}

	return mul
}

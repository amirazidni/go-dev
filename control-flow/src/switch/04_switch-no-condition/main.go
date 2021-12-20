package main

import "fmt"

func grading(point int) string {
	switch {
	case point >= 90:
		return "Excellent"
	case point >= 80 && point < 90:
		return "Good"
	case point >= 70 && point < 80:
		return "Pass"
	case point >= 50 && point < 70:
		return "Unsatisfactory"
	default:
		return "Fail"
	}
}

func main() {
	fmt.Println("91: ", grading(81))
	fmt.Println("72: ", grading(72))
	fmt.Println("50: ", grading(50))
	fmt.Println("10: ", grading(10))
}

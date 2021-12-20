package main

import "fmt"

func grading(point int) string {
	if point >= 90 {
		return "Excellent"
	} else if point >= 80 && point < 90 {
		return "Good"
	} else if point >= 70 && point < 80 {
		return "Pass"
	} else if point >= 50 && point < 70 {
		return "Unsatisfactory"
	} else {
		return "Fail"
	}
}

func main() {
	fmt.Println("91: ", grading(81))
	fmt.Println("72: ", grading(72))
	fmt.Println("50: ", grading(50))
	fmt.Println("10: ", grading(10))
}

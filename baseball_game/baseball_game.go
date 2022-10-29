package main

import (
	"fmt"
	"strconv"
)

func main() {
	//insert different test cases into this slice
	var ops3 = []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(ops3))
}

func calPoints(ops []string) int {
	var result int
	var stack1 []int

	for _, i := range ops {
		length := len(stack1)
		if i == "+" {
			var p = stack1[length-1] + stack1[length-2]
			stack1 = append(stack1, p)
		} else if i == "C" {
			stack1 = stack1[:len(stack1)-1]
		} else if i == "D" {
			var n int = stack1[length-1] * 2
			stack1 = append(stack1, n)
		} else {
			s, _ := strconv.Atoi(i)
			stack1 = append(stack1, s)
		}
	}
	for _, f := range stack1 {
		result = result + f
	}
	return result
}

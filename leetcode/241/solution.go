package main

import (
	"fmt"
	"strconv"
	"strings"
)

func diffWaysToCompute(expression string) []int {
	num, err := strconv.Atoi(strings.TrimSpace(expression))
	if err == nil {
		return []int{num}
	}

	var result []int
	for index, char := range expression {
		curStr := string(char)
		if curStr == "+" || curStr == "-" || curStr == "*" {
			left := expression[0:index]
			right := expression[index+1:]

			leftResultArr := diffWaysToCompute(left)
			rightResultArr := diffWaysToCompute(right)

			for _, leftResult := range leftResultArr {
				for _, rightResult := range rightResultArr {
					num := 0
					if curStr == "+" {
						num = leftResult + rightResult
					} else if curStr == "-" {
						num = leftResult - rightResult
					} else {
						num = leftResult * rightResult
					}
					result = append(result, num)
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(diffWaysToCompute("1 + 2 * 3"))
}

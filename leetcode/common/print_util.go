package common

import "fmt"

func PrintArr[T any](arr []T) {
	if arr == nil {
		fmt.Printf("arr is nil \n")
		return
	}

	if len(arr) == 0 {
		fmt.Printf("arr len is 0 \n")
		return
	}

	for index, data := range arr {
		fmt.Printf("No.%v,\t Value: %v\n", index, data)
	}
}

package main

import "fmt"

type Node struct {
	V    int
	Next *Node
}

func createTargetArray(nums []int, index []int) []int {
	HeadNode := &Node{}

	for numIndex, insertIndex := range index {
		insertNum := nums[numIndex]

		postInsertNode := HeadNode
		for linkIndex := 0; linkIndex < insertIndex; linkIndex++ {
			postInsertNode = postInsertNode.Next
		}

		insertNode := &Node{V: insertNum}

		t := postInsertNode.Next
		postInsertNode.Next = insertNode
		insertNode.Next = t
	}

	result := make([]int, 0)

	curNode := HeadNode.Next
	for {
		result = append(result, curNode.V)
		if curNode.Next == nil {
			break
		}
		curNode = curNode.Next

	}

	return result
}

func main() {
	arr := make([]int, 6)
	arr = []int{1, 2, 3, 4, 5, 0}
	copy(arr[4:], arr[3:])
	arr[3] = 10
	fmt.Println(arr)
}

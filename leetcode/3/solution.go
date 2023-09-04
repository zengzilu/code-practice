package _3

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	leftIndex, rightIndex := 0, -1
	strLen := len(s)
	charMap := make(map[uint8]int)

	for rightIndex+1 < strLen {
		toAdd := s[rightIndex+1]

		if index, ok := charMap[toAdd]; ok {
			leftIndex = index + 1
		}

		charMap[toAdd] = rightIndex + 1
		rightIndex++

		if rightIndex-leftIndex+1 > maxLen {
			maxLen = rightIndex - leftIndex + 1
		}

	}

	return maxLen
}

func isConflict(leftIndex, rightIndex int, s string, toAdd uint8) (int, bool) {
	conflictIndex, conflict := -1, false
	for index := leftIndex; index <= rightIndex; index++ {
		if s[index] == toAdd {
			conflict = true
			conflictIndex = index
			break
		}
	}

	return conflictIndex, conflict
}

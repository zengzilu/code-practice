package _40

// 双指针
func trap(height []int) int {
	heightLen := len(height)
	if heightLen <= 2 {
		return 0
	}

	leftMax := height[0]
	rightMax := height[heightLen-1]

	leftP := 1
	rightP := heightLen - 1

	total := 0
	for leftP <= rightP {
		if leftMax <= rightMax {
			// 左指针右移动
			if height[leftP] < leftMax {
				total += leftMax - height[leftP]
			} else {
				leftMax = height[leftP]
			}
			leftP++
		} else {
			// 右指针左移动
			if height[rightP] < rightMax {
				total += rightMax - height[rightP]
			} else {
				rightMax = height[rightP]
			}
			rightP--
		}
	}
	return total
}

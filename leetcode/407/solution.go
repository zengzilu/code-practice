package _07

import "math"

// 为每一个垂直柱子设置一个值 wateredHeight ，表示此位置（x,y）接雨水后的最高值
// * 这个值有一个特点，预期邻接的柱子最大的 wateredHeight 不会大于 此柱子的 wateredHeight, 整体的雨量 sum (wateredHeight[x][y] - height[x][y])
// 最外圈无法接水，可以确认最外圈 wateredHeight == 其柱子height

// 遍历 x,y 并标记其是否已经被计算过水量
// wateredHeight(x,y) = max(height(x,y) , min(wateredHeight(x-1,y), wateredHeight(x+1,y), wateredHeight(x,y-1), wateredHeight(x,y+1)))

// 按照怎么样的逻辑去遍历 x,y
// 1. 确认最外圈的waterHeight，找到最低点
// 2. 以最低点为延伸点 s ，向内探测 a
//		2.1 如果 a.高度 <= s.wateredHeight, 则不管a点其他三面如何，一定可以接水且a.wateredHeight == s.wateredHeight
//		2.2 如果 a.高度 > s.wateredHeight, 则a一定无水可接，a.wateredHeight = a.height
// 3. 在探测完a点后，已经计算过wateredHeight所有节点仍然是一个封闭的圈，可以继续在圈边寻找最低点重复#2

var waterSum = 0

func trapRainWater(height [][]int) int {
	xLen, yLen := len(height), len(height[0])
	if xLen <= 2 || yLen <= 2 {
		return 0
	}

	// 初始化wateredHeight
	wateredHeight := make([][]int, xLen)
	for x := 0; x < xLen; x++ {
		wateredHeight[x] = make([]int, yLen)

		for y := 0; y < yLen; y++ {
			if x == 0 || x == xLen-1 || y == 0 || y == yLen-1 {
				wateredHeight[x][y] = height[x][y]
			} else {
				// 还没有被计算过
				wateredHeight[x][y] = -1
			}
		}
	}

	// 寻找统计水量
	for visitNextPoint(wateredHeight, height) {
	}
	return waterSum
}

func visitNextPoint(wateredHeight [][]int, height [][]int) bool {
	xLen, yLen := len(wateredHeight), len(wateredHeight[0])
	minHeight := math.MaxInt

	minX, minY := -1, -1
	neighboorX, neighboorY := -1, -1
	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {

			// visited && 小于当前最小值
			if wateredHeight[x][y] >= 0 && wateredHeight[x][y] < minHeight {
				// 周围有尚未visited的点
				px, py := getVisitablePoint(wateredHeight, x, y)
				if px >= 0 && py >= 0 {
					// 当前的最矮柱子的邻接点
					minHeight = min(minHeight, wateredHeight[x][y])
					minX, minY = px, py
					neighboorX, neighboorY = x, y
				}
			}
		}
	}

	if minX >= 0 && minY >= 0 {
		// 选中
		// wateredHeight[minX][minY] = max(wateredHeight[x][y], height[x][y])
		if wateredHeight[neighboorX][neighboorY] > height[minX][minY] {
			waterSum += wateredHeight[neighboorX][neighboorY] - height[minX][minY]
			wateredHeight[minX][minY] = wateredHeight[neighboorX][neighboorY]
		} else {
			wateredHeight[minX][minY] = height[minX][minY]
		}
		return true
	}

	return false
}

func min(intNums ...int) int {
	minValue := math.MaxInt
	for _, num := range intNums {
		if num < minValue {
			minValue = num
		}
	}

	return minValue
}

func max(intNums ...int) int {
	maxValue := math.MinInt
	for _, num := range intNums {
		if num > maxValue {
			maxValue = num
		}
	}

	return maxValue
}

func getVisitablePoint(wateredHeight [][]int, x, y int) (px, py int) {
	// 上
	px, py = x-1, y
	if isVisitable(wateredHeight, px, py) {
		return
	}

	// 下
	px, py = x+1, y
	if isVisitable(wateredHeight, px, py) {
		return
	}

	// 左
	px, py = x, y-1
	if isVisitable(wateredHeight, px, py) {
		return
	}

	// 右
	px, py = x, y+1
	if isVisitable(wateredHeight, px, py) {
		return
	}

	return -1, -1
}

func isVisitable(wateredHeight [][]int, x, y int) bool {
	if x >= 0 && x < len(wateredHeight) && y >= 0 && y < len(wateredHeight[0]) && wateredHeight[x][y] == -1 {
		return true
	}

	return false
}

package main

import (
	"container/heap"
	"math"
)

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

func trapRainWater(height [][]int) int {
	water := 0

	xLen, yLen := len(height), len(height[0])
	if xLen <= 2 || yLen <= 2 {
		return 0
	}

	// 初始化墙壁
	wallHeap := &walls{}
	visited := make([][]bool, xLen)
	for x := 0; x < xLen; x++ {
		visited[x] = make([]bool, yLen)
		for y := 0; y < yLen; y++ {
			if x == 0 || x == xLen-1 || y == 0 || y == yLen-1 {
				heap.Push(wallHeap, wall{x, y, height[x][y]})
				visited[x][y] = true
			} else {
				// 还没有被计算过
				visited[x][y] = false
			}
		}
	}

	// 开始遍历, 只有还有有效的墙壁即循环遍历 => 直到所有的墙壁旁边的墙壁都被visited
	for wallHeap.Len() > 0 {
		// 取最矮墙壁
		shortestWall := heap.Pop(wallHeap).(wall)

		// 向4个方向遍历
		directions := []int{-1, 0, 1, 0, -1}
		for dirIndex := 0; dirIndex < 4; dirIndex++ {
			newX, newY := shortestWall.x+directions[dirIndex], shortestWall.y+directions[dirIndex+1]
			if newX >= 0 && newX < xLen && newY >= 0 && newY < yLen && visited[newX][newY] == false {
				if height[newX][newY] < shortestWall.z {
					water += shortestWall.z - height[newX][newY]
				}
				heap.Push(wallHeap, wall{newX, newY, max(height[newX][newY], shortestWall.z)})
				visited[newX][newY] = true
			}
		}
	}

	return water
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

type wall struct {
	x, y, z int
}
type walls []wall

func (w walls) Len() int { return len(w) }

func (w walls) Less(i, j int) bool { return w[i].z < w[j].z }

func (w walls) Swap(i, j int) { w[i], w[j] = w[j], w[i] }

func (w *walls) Push(x interface{}) { *w = append(*w, x.(wall)) }

func (w *walls) Pop() interface{} { r := (*w)[len(*w)-1]; *w = (*w)[:len(*w)-1]; return r }

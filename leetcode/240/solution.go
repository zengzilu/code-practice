package _40

// 从左至右递增, 从上至下递增
// 方法1: 暴力搜索, one by one , 时间复杂度 O(m*n)
// 方法2: 二分法, 有几种不同的二分法.
//       a. 逐行二分法
//       b. 二位二分法, 在二位中选一点v, 将二维数组分为四个象限
//          -- 第一象限和第三象限可能大于或小于v.
//          -- 第二象限一定小于v, 第四象限一定大于v. 将其与target比对,快速排除不满足的象限递归查询
//          -- 这个算法一次只能排除1/4. 并没有二分的效率
// 方法3: 从右上角开始寻找, target大于v则向下, target小于v则向左, target等于v则找到结果. 时间复杂度为O(m+n)
//       思路来源:
//       -- 已经发现了图中选择一点4个象限的划分
//       -- 开始想从左上角开始,向对角线寻址,每次可以排除左上角或者右下角数据区域. 这个想法类似于图的"四分法", 第一次误判了第一,第三象限的取值
//       -- 嫖了一下题解, 可以从右上角开始, 因为向下和向左大小方向是不一样的, 可以单调寻址, 如果从左上角开始,向下和向右都是增加的, 需要双向寻址
//       -- 从右上角寻址这个太妙了, 对应的如果从左下角寻址也是一样的效果

// 右上角寻址
func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	if row == 0 || col == 0 {
		return false
	}

	curRow, curCol := 0, col-1
	for curRow < row && curCol >= 0 {
		if matrix[curRow][curCol] > target {
			// 向左
			curCol--
		} else if matrix[curRow][curCol] < target {
			// 向下
			curRow++
		} else {
			return true
		}
	}

	return false
}

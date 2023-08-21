## 双指针为什么是正确的？

1. 为什么通过比对左右指针的两边最大值,可以判断移动哪个指针
2. 指针移动代表着什么
3. 计算结束的条件是什么
---
1. 以leftMax, rightMax, leftRealMax, rightRealMax代表左指针左边记录的最大值,右指针右边记录的最大值,左指针左边实际的最大值,右指针右边实际的最大值
   对于左指针格子有: leftRealMax == leftMax, rightRealMax >= rightMax
   对于右指针格子有: leftRealMax >= leftMax, rightRealMax == rightMax
   
   如果 leftMax < rightMax , 则有 leftRealMax == leftMax < rightMax <= rightRealMax, 即leftRealMax < rightRealMax.
   翻译过来就是, 如果左边的最大值小于右边的最大值, 则左指针列能存储的最大雨量可以确认, 为leftRealMax - height[left]. (负数则不用存)
   这就是指针移动的逻辑

2. 指针移动代表指针列的雨量已经被确认并计算到总值中

3. 计算结束的条件是左指针>右指针, 重合时仍然需要计算
   
---

一列能存多少水取决于min(maxLeft, maxRight)

左右两个指针分别记录做左右两边正在等待处理的列, 并向中间靠齐, 直至相交

在指针移动的过程中

对于左指针的格子, 左边的最高墙一定等于leftMax, 因为对于左指针左边所有的格子都计算过. 右边的最高墙的最小值等于rightMax, 之间可能有更高墙.
此时如果leftMax < rightMax 则左指针接水以leftMax为准. 当左指针的水量被确认就可以将左指针向右移动

右指针也是相同情况
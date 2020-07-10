package sort_

import (
	"fmt"
)

// 分治法(分解,解决,合并)(递归情况 recursive case; 基本情况 base case)
// 增量法

// 各种排序算法
// 1.最值排序/简单选择排序		n*n
// 2.双最值排序/二元选择排序	n*n
// 3.插入排序 	n*n
// 4.堆排序		n*logN
// 5.冒泡排序		n*n
// 6.快速排序		n*logN
// 7.归并排序		n*logN
// 8.桶排序		n
// 9.hash排序	n
// 10.基数排序	n

// 冒泡排序
// 遍历,反复交换相邻的未按照次序排列的元素,效果与选择排序类似,保证最值在一端

// 插入排序
// 遍历,使 nums[0:1] 有序,使 nums[0:2] 有序,使 nums[0:3] 有序......使 nums[0:n] 有序

// 快速排序
// 递归:选择一个基准值,使基准值成为中位数,切割左右序列,重复 基准--中位--分割 的流程

// 归并排序
// 与快速排序原理相反,分割为两两一组的子序列,对所有的子序列排序
// 重复 两两结合--双指针排序--合并 的流程

// 基数排序
// 将所有的元素转换为[]bit,高位对齐,低位补 0,之后桶排序或者交换

func printSort(v []int, f func([]int) []int)  {
	fmt.Print(v, " => ")
	fmt.Print(f(v))
	fmt.Println()
}

func swap(a, b int) (int, int){
	//return b^a^a, a^b^b
	a = a^b
	b = b^a
	a = a^b

	return a, b
}
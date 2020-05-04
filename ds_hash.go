package leetcode

import "math"

// 散列表
type hashTable struct {
	buf []data
}

func hash(key interface{}) int {
	return 0
}
// 除法散列法:size 最好为一个素数,不要为 2 的 n 次方
func hashMod(k int) int {
	size := 701
	return k%size
}
// a % b == a & (b - 1)
//前提：b 为 2^n

// 乘法散列法:key*A 取小数部分kA,kA乘以m向下取整
func hashXX(k int) int {
	A := (math.Sqrt(5)-1)/2
	m := float64(1 << 8)
	return  int(math.Trunc(math.Mod(float64(k)*A, 1)*m))
}

//// 使用链接法解决冲突
//type hashTable struct {
//	buf []doubleLinkedList
//}
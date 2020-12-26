package pkg

import (
	"fmt"
	"sort"
	"testing"
)

// 56 合并区间
//给出一个区间的集合，请合并所有重叠的区间。
//
//示例 1:
//
//输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2:
//
//输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
type tag struct {
	start bool
	value int
}
type tags []tag

func (p tags) Len() int           { return len(p) }
func (p tags) Less(i, j int) bool { return p[i].value < p[j].value }
func (p tags) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Test_merge(t *testing.T) {
	ss := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 9}, {4, 5}},
	}

	for _, v := range ss {
		fmt.Println(v, merge(v))
	}
}

func makeTags(list [][]int) tags {
	tags := make([]tag, 0, 10)
	for _, v := range list {
		tags = append(tags, tag{start: true, value: v[0]})
		tags = append(tags, tag{start: false, value: v[1]})
	}

	//// 同向取其一,异向都不取
	//for k, v := range tags {
	//
	//}

	return tags
}

// current_start current_end list   to
//     Y             Y        S     re + start S
//     Y             Y        E     set E
//     Y             N        S     continue
//     Y             N        E     set E
//     N             Y        S     error
//     N             Y        E     error
//     N             N        S     set S
//     N             N        E     error
func merge(intervals [][]int) [][]int {
	ts := makeTags(intervals)
	sort.Sort(ts)
	fmt.Println("tags", ts)

	//
	res := make([][]int, 0)
	startF := false
	endF := false
	var start int
	var end int
	for i := 0; i < len(ts); i++ {
		v := ts[i].value
		flag := ts[i].start
		if startF && endF && flag {
			if start != end {
				res = append(res, []int{start, end})
				endF = false
				start = v
			}
			continue
		}
		if startF && endF && !flag {
			end = v
		}
		if startF && !endF && flag {
			continue
		}
		if startF && !endF && !flag {
			endF = true
			end = v
			continue
		}
		if !startF && endF && flag {
			panic("N Y S")
		}
		if !startF && endF && !flag {
			panic("N Y E")
		}
		if !startF && !endF && flag {
			startF = true
			start = v
			continue
		}
		if !startF && !endF && !flag {
			panic("N N E")
		}

		//if startF {
		//	if endF {
		//		if flag {
		//			res = append(res, []int{start, end})
		//			endF = false
		//			start = v
		//		}
		//	} else {
		//		if !flag {
		//			endF = true
		//			end = v
		//		}
		//	}
		//
		//} else {
		//	if !endF && flag {
		//		start = v
		//	}
		//}
	}
	res = append(res, []int{start, end})

	return res
}

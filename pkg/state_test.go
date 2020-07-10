package pkg

import (
	"fmt"
	"math/rand"
	"testing"
)

// 过河问题
func find(states [][]int, res [][]int) {
	if len(states) == 0 {
		return
	}
	i := rand.Intn(len(states))
	v := states[i]
	res = append(res, v)
	fmt.Println(len(res),len(append(states[:i], states[i+1:]...)))
	find(append(states[:i], states[i+1:]...), res)
}

func TestFmtInSlice(t *testing.T){
	fmt.Print("\rbbc")
	fmt.Print("\rac")
	//fmt.Print("\r")
	return
	states := make([][]int, 0)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			states = append(states, []int{i, j})
		}
	}
	fmt.Println(states)
	find(states, [][]int{})

}


func TestState(t *testing.T) {
	oxAndTiger(2, 2)

}

type stateCT struct {
	ac int //A cattle
	at int //B tiger
}

// 过河问题
// 三牛三虎过河,任意一岸的虎的数量大于牛的数量,牛会被吃掉,船最多能装 2 只动物
func oxAndTiger(cattle,tiger int){
	tr := [][]int{{0,1},{0,2},{1,0},{1,1},{2,0}}
	states := make([]stateCT, 0)
	for i := 0; i <= cattle; i++ {//牛
		for j := 0; j <= tiger; j++ {//虎
			//fmt.Println([]int{i, j, 3-i, 3-j})
			if (j > i && i != 0) || (tiger-j > cattle-i && cattle-i!=0) {// 排除牛被吃的情况
				continue
			}
			states = append(states, stateCT{i, j})
		}
	}

	start := states[0]
	end := states[len(states)-1]
	//fmt.Println(start, end)

	for _, v := range route(start, end, states[1:], tr, []stateCT{}) {
		fmt.Println(v)
	}

	//fmt.Println(states)
	//fmt.Println(tr)
}

func route(start stateCT, end stateCT, states []stateCT, tr [][]int, sol []stateCT) (res [][]stateCT) {
	sol = append(sol, start)
	if start == end{
		temp := make([]stateCT, len(sol))
		copy(temp, sol)
		res = append(res, temp)
		//fmt.Println(temp)
		return [][]stateCT{temp}
	}
	for i, v := range states {
		x := v.ac - start.ac
		y := v.at - start.at
		if x < 0 || y < 0 || x+y > 2 {
			continue
		}
		for j := 0; j < len(tr); j++ {
			if x == tr[j][0] && y == tr[j][1] {
				fmt.Println(start, "===>", v)
				temp := make([]stateCT, 0, len(states))
				temp = append(temp, states[:i]...)
				temp = append(temp, states[i+1:]...)
				//route(v, end, temp, tr, sol)
				res = append(res, route(v, end, temp, tr, sol)...)
			}
		}
	}
	return res
}
package pkg

import (
	"bytes"
	"fmt"
	"testing"
)

//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//上图为 8 皇后问题的一种解法。
//
//给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
//
//每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//示例:
//
//输入: 4
//输出: [
//[".Q..",  // 解法 1
//"...Q",
//"Q...",
//"..Q."],
//
//["..Q.",  // 解法 2
//"Q...",
//"...Q",
//".Q.."]
//]
//解释: 4 皇后问题存在两个不同的解法。
//
//
//提示：
//
//皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一到七步，可进可退。（引用自 百度百科 - 皇后 ）
func Test_solveNQueens(t *testing.T) {
	fmt.Println(solveNQueens(4))
}

func solveNQueensMine(n int) board {
	l = n
	b := make([][]bool, l)
	for k, _ := range b {
		b[k] = make([]bool, l)
	}

	for i := 0; i < n; i++ {
		bn, ok := board(b).place(0, 0)
		board(bn).place(0, 0)
		if ok {

		}
	}

	return b
}

func solveNQueens(n int) [][]string {
	l = n
	b := make([][]bool, l)
	for k, _ := range b {
		b[k] = make([]bool, l)
	}

	bn, ok := board(b).place(0, 0)
	if ok {
		board(bn).place(0, 0)
	}

	return solveNQueens(0)
}

var l = 4 //棋盘大小
var a = 7 //攻击范围
// 棋盘, true 为被占用;false 未必占用
type board [][]bool

func (b board) String() string {
	var buffer bytes.Buffer
	for _, v := range b {
		for _, vv := range v {
			if vv {
				buffer.WriteByte('X')
			} else {
				buffer.WriteByte(' ')
			}
			buffer.WriteByte(' ')
		}
		buffer.WriteByte('\n')
	}
	return buffer.String()
}

func TestBoard_String(t *testing.T) {
	b := make([][]bool, l)
	for k, _ := range b {
		b[k] = make([]bool, l)
	}

	fmt.Println(board(b))
}

func toXY(xB, yB int) (int, int) {
	return xB - 1, yB - 1
}

func (b board) place(x, y int) (bn board, ok bool) {
	if b[y][x] {
		return
	}
	// 新棋盘
	bn = make([][]bool, l)
	for k, _ := range bn {
		bn[k] = make([]bool, l)
	}
	copy(bn, b)
	fmt.Println(bn)
	defer func() {
		fmt.Println(bn)
	}()
	// 放置
	bn[y][x] = true //本身
	// 行
	for i := 0; i < l; i++ {
		bn[i][x] = true
	}
	// 列
	for i := 0; i < l; i++ {
		bn[y][i] = true
	}
	// 主对角线
	// 最上方坐标
	for i := 1; i <= a; i++ {
		if xi, yi := x+i, y+i; xi < l && yi < l {
			bn[yi][xi] = true
		}
		if xi, yi := x-i, y-i; xi >= 0 && yi >= 0 {
			bn[yi][xi] = true
		}
	}
	// 次对角线
	// 最上方坐标
	for i := 1; i <= a; i++ {
		if xi, yi := x+i, y-i; xi < l && yi < l && xi >= 0 && yi >= 0 {
			bn[yi][xi] = true
		}
		if xi, yi := x-i, y+i; xi < l && yi < l && xi >= 0 && yi >= 0 {
			bn[yi][xi] = true
		}
	}

	return
}

func TestBoard_Place(t *testing.T) {
	b := make([][]bool, l)
	for k, _ := range b {
		b[k] = make([]bool, l)
	}
	board(b).place(toXY(6, 7))
}

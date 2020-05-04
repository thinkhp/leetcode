package leetcode

import (
	"fmt"
	"testing"
)

//给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
//示例:
//
//board =
//[
//['A','B','C','E'],
//['S','F','C','S'],
//['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true
//给定 word = "SEE", 返回 true
//给定 word = "ABCB", 返回 false
//
//
//提示：
//
//board 和 word 中只包含大写和小写英文字母。
//1 <= board.length <= 200
//1 <= board[i].length <= 200
//1 <= word.length <= 10^3

func TestExist(t *testing.T) {
	board := [][]byte{
		{'S', 'A', 'B', 'C', 'E'},
		{'E', 'S', 'F', 'C', 'S'},
		{'F', 'A', 'D', 'E', 'E'},
	}

	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCB"))
}

func dpsBoard(board [][]byte, directions [][]int, mark [][]bool, x, y int, str []byte) bool {
	mark[y][x] = true
	if len(str) == 0 {
		return true
	}
	for _, v := range directions {
		xn, yn := x+v[0], y+v[1]
		if yn >= 0 && yn < len(board) && xn >= 0 && xn < len(board[yn]) && board[yn][xn] == str[0] && !mark[yn][xn]{
			if dpsBoard(board, directions, mark, xn, yn, str[1:]) {
				return true
			}
		}
	}
	mark[y][x] = false
	return false
}

type xy struct {
	x int
	y int
}

func exist(board [][]byte, word string) bool {
	directions := [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	mark := make([][]bool, 0)
	xys := make([]xy, 0)
	for y, v := range board {
		markRow := make([]bool, 0)
		for x, vv := range v {
			markRow = append(markRow, false)
			if vv == word[0] {
				xys = append(xys, xy{x, y})
			}
		}
		mark = append(mark, markRow)
	}

	for _, v := range xys {
		if dpsBoard(board, directions, mark, v.x, v.y, []byte(word)[1:]) {
			fmt.Println("mark", mark)
			return true
		}
	}

	return false
}

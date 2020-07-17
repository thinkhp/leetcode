package pkg

import (
	"fmt"
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	//grid := [][]int{
	//	{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	//	{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
	//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	//}
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	}

	fmt.Println(maxAreaOfIsland(grid))

}

func maxAreaOfIsland(grid [][]int) int {
	//direction := [][]int{
	//	{0, -1},
	//	{0, 1},
	//	{-1, 0},
	//	{1, 0},
	//}
	max := 0
	for y, v := range grid {
		for x, vv := range v {
			if vv != 0 {
				s := dfsGrid(grid, x, y)
				if max < s {
					max = s
				}
			}
		}
	}
	return max
}

func dfsGrid(grid [][]int, x, y int) int {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) || grid[y][x] == 0 {
		return 0
	}
	grid[y][x] = 0
	sum := 1
	sum += dfsGrid(grid, x, y+1)
	sum += dfsGrid(grid, x, y-1)
	sum += dfsGrid(grid, x+1, y)
	sum += dfsGrid(grid, x-1, y)
	return sum
}

package main

import "fmt"

func main() {
	var (
		grid [][]byte
		gridorig [][]string
		islands int
	) 

	grid = [][]byte{
		{'1','1','1','1','0'}, 
		{'1','1','0','1','0'}, 
		{'1','1','0','0','0'}, 
		{'0','0','0','0','0'},
	}
	gridorig = [][]string{
		{"1","1","1","1","0"}, 
		{"1","1","0","1","0"}, 
		{"1","1","0","0","0"}, 
		{"0","0","0","0","0"},
	}
	islands = numIslands(grid)
	fmt.Printf("The number of islands in the map %v is %d\n", gridorig, islands)

	grid = [][]byte{
		{'1','1','0','0','0'}, 
		{'1','1','0','0','0'}, 
		{'0','0','1','0','0'}, 
		{'0','0','0','1','1'},
	}
	gridorig = [][]string{
		{"1","1","0","0","0"}, 
		{"1","1","0","0","0"}, 
		{"0","0","1","0","0"}, 
		{"0","0","0","1","1"},
	}
	islands = numIslands(grid)
	fmt.Printf("The number of islands in the map %v is %d\n", gridorig, islands)

	grid = [][]byte{
		{'1','1','0','0','0'}, 
		{'1','1','0','0','0'}, 
		{'0','0','1','0','0'}, 
		{'0','0','1','1','1'},
	}
	gridorig = [][]string{
		{"1","1","0","0","0"}, 
		{"1","1","0","0","0"}, 
		{"0","0","1","0","0"}, 
		{"0","0","1","1","1"},
	}
	islands = numIslands(grid)
	fmt.Printf("The number of islands in the map %v is %d\n", gridorig, islands)
}

// depth first search (DFS)
func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	rows, cols, islands := len(grid), len(grid[0]), 0

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows {
			return
		}
		if col < 0 || col >= cols {
			return
		}
		if grid[row][col] == '0' {
			return
		}

		// marks as visited
		grid[row][col] = '0'
		dfs(row - 1, col)
		dfs(row + 1, col)
		dfs(row, col - 1)
		dfs(row, col + 1)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				islands++
				dfs(row, col)
			}
		}
	}

	return islands
}

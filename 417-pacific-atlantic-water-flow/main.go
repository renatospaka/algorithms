package main

import "fmt"

func main() {
	var (
		hights [][]int
		reachBoth [][]int
	)

	hights = [][]int{{1,2,2,3,5},{3,2,3,4,4},{2,4,5,3,1},{6,7,1,4,5},{5,1,1,2,4}}
	reachBoth = pacificAtlantic(hights)
	fmt.Printf("The coordinates that reach both Atlantic and Pacific oceans from %v are: %v\n", hights, reachBoth)

	hights = [][]int{{1}}
	reachBoth = pacificAtlantic(hights)
	fmt.Printf("The coordinates that reach both Atlantic and Pacific oceans from %v are: %v\n", hights, reachBoth)

	hights = [][]int{{1, 3, 4}, {3, 3, 3}}
	reachBoth = pacificAtlantic(hights)
	fmt.Printf("The coordinates that reach both Atlantic and Pacific oceans from %v are: %v\n", hights, reachBoth)
}

// depth first search (DFS)
func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0]) 
	both := [][]int{}
	pacific, atlantic := make(map[int]bool), make(map[int]bool)

	var dfs func(int, int, map[int]bool, int)
	dfs = func(row int, col int, ocean map[int]bool, previousHeight int) {
		idx := (row * cols) + col
		if row < 0 || row == rows || col < 0 || col == cols {
			return
		}
		if heights[row][col] < previousHeight || ocean[idx] {
			return
		}

		ocean[idx] = true
		dfs(row + 1, col, ocean, heights[row][col])
		dfs(row - 1, col, ocean, heights[row][col])
		dfs(row, col + 1, ocean, heights[row][col])
		dfs(row, col - 1, ocean, heights[row][col])
	}

	for col := 0; col < cols; col++ {
		dfs(0, col, pacific, heights[0][col])									// pacific north boundary
		dfs(rows - 1, col, atlantic, heights[rows - 1][col])	// atlantic south boundary
	}

	for row := 0; row < rows; row++ {
		dfs(row, 0, pacific, heights[row][0])									// pacific west boundary
		dfs(row, cols - 1, atlantic, heights[row][cols - 1])	// atlantic east boundary
	}

	//checks weather the coordinate reaches both oceans
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			idx := (row * cols) + col
			// fmt.Printf("idx: %d || atlantic: %t || pacific: %t\n", idx, atlantic[idx], pacific[idx])
			if atlantic[idx] && pacific[idx] {
				both = append(both, []int{row, col})
			}
		}
	}

	return both
}

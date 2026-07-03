package questions

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	directions := [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	bfs := func(row, col int) {
		grid[row][col] = '0'
		queue := [][2]int{{row, col}}
		for head := 0; head < len(queue); head++ {
			r, c := queue[head][0], queue[head][1]
			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != '1' {
					continue
				}
				grid[nr][nc] = '0'
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}

	result := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				bfs(i, j)
				result++
			}
		}
	}
	return result
}

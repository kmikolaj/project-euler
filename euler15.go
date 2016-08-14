package main

import "fmt"

func lattice_path_count(width int, height int) int {
	var grid = make([][]int, width+1)
	for i, _ := range grid {
		grid[i] = make([]int, height+1)
	}

	for i := 0; i <= width; i++ {
		for j := 0; j <= height; j++ {
			if i == 0 || j == 0 {
				grid[i][j] = 1
			} else {
				grid[i][j] = grid[i-1][j] + grid[i][j-1]
			}
		}
	}

	return grid[width][height]
}

func main() {
	fmt.Println("Result:", lattice_path_count(20, 20))
}

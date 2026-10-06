package main

import "testing"

func TestNumIslands(t *testing.T) {
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	if got := NumIslands(grid1); got != 1 {
		t.Errorf("NumIslands(grid1) = %d; want 1", got)
	}

	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	if got := NumIslands(grid2); got != 3 {
		t.Errorf("NumIslands(grid2) = %d; want 3", got)
	}
}

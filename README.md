# Number of Islands

Depth-First Search (DFS) graph traversal in Go to count connected 4-directional land components.

## Problem Description
Given an $m \times n$ 2D binary grid `grid` which represents a map of `'1'`s (land) and `'0'`s (water), return the number of islands.

### Example
- Input:
  ```text
  1 1 0 0 0
  1 1 0 0 0
  0 0 1 0 0
  0 0 0 1 1
  ```
- Output: `3`

## Approach & Complexity
Iterate over every cell. Upon encountering `'1'`, increment count and recursively sink all adjacent land cells using DFS.

- **Time Complexity:** $O(M \times N)$
- **Space Complexity:** $O(M \times N)$ recursive stack in worst case.

## How to Run & Test
```bash
go run islands.go
go test -v ./...
```

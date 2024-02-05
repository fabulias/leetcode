package solution

/*
https://leetcode.com/problems/number-of-islands/

Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input: grid = [

	["1","1","1","1","0"],
	["1","1","0","1","0"],
	["1","1","0","0","0"],
	["0","0","0","0","0"]

]
Output: 1
Example 2:

Input: grid = [

	["1","1","0","0","0"],
	["1","1","0","0","0"],
	["0","0","1","0","0"],
	["0","0","0","1","1"]

]
Output: 3

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*/

type Queue struct {
	elements [][]int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value []int) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() []int {
	if q.IsEmpty() {
		return nil
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Front() []int {
	if q.IsEmpty() {
		return nil
	}
	return q.elements[0]
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func bfs(grid [][]byte, r, c, m, n int) {
	queue := NewQueue()
	queue.Enqueue([]int{r, c})
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	for !queue.IsEmpty() {
		curr := queue.Dequeue()

		for _, dir := range dirs {
			nr := curr[0] + dir[0]
			nc := curr[1] + dir[1]
			if 0 <= nr && nr < m && 0 <= nc && nc < n && grid[nr][nc] == '1' {
				grid[nr][nc] = '0'
				queue.Enqueue([]int{nr, nc})
			}
		}
	}
}

func numIslands1(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	count := 0

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '1' {
				count++
				bfs(grid, r, c, m, n)
			}
		}
	}

	return count
}

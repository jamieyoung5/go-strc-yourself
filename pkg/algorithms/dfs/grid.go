package dfs

type Point struct {
	X, Y int
}

// directions all eight possible movements from a cell.
var directions = []Point{
	{-1, 0},  // up
	{1, 0},   // down
	{0, -1},  // left
	{0, 1},   // right
	{-1, -1}, // up-left
	{-1, 1},  // up-right
	{1, -1},  // down-left
	{1, 1},   // down-right
}

// SearchGrid searches for occurrences of a 'pattern' in a grid and returns variations of the pattern found
func SearchGrid[T comparable](grid [][]T, pattern []T) []map[T]Point {
	rows := len(grid)
	if rows == 0 {
		return nil
	}
	cols := len(grid[0])
	if cols == 0 || len(pattern) == 0 {
		return nil
	}

	var occurrences []map[T]Point
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			// check all directions
			for _, dir := range directions {
				if occurrence := searchFromCell(grid, pattern, i, j, dir.X, dir.Y); occurrence != nil {
					occurrences = append(occurrences, occurrence)
				}
			}
		}
	}

	return occurrences
}

func searchFromCell[T comparable](grid [][]T, pattern []T, x, y, dx, dy int) map[T]Point {
	rows := len(grid)
	cols := len(grid[0])

	occurrence := make(map[T]Point, rows)
	for k := 0; k < len(pattern); k++ {
		nx := x + k*dx
		ny := y + k*dy
		if nx < 0 || ny < 0 || nx >= rows || ny >= cols {
			return nil
		}
		if grid[nx][ny] != pattern[k] {
			return nil
		}
		occurrence[pattern[k]] = Point{nx, ny}
	}
	return occurrence
}

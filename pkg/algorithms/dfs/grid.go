package dfs

type Point struct {
	X, Y int
}

type PatternOccurrence[T comparable] struct {
	Points    map[T]Point
	Direction Direction
}

type Direction Point

var (
	Up        = Direction(Point{-1, 0})
	Down      = Direction(Point{1, 0})
	Left      = Direction(Point{0, -1})
	Right     = Direction(Point{0, 1})
	UpLeft    = Direction(Point{-1, -1})
	UpRight   = Direction(Point{-1, 1})
	DownLeft  = Direction(Point{1, -1})
	DownRight = Direction(Point{1, 1})

	directions = []Direction{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}
)

// SearchGrid searches for occurrences of a 'pattern' in a grid and returns variations of the pattern found
func SearchGrid[T comparable](grid [][]T, pattern []T) []PatternOccurrence[T] {
	rows := len(grid)
	if rows == 0 {
		return nil
	}
	cols := len(grid[0])
	if cols == 0 || len(pattern) == 0 {
		return nil
	}

	var occurrences []PatternOccurrence[T]
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			// check all directions
			for _, dir := range directions {
				if points := searchFromCell(grid, pattern, i, j, dir.X, dir.Y); points != nil {
					occurrences = append(occurrences, PatternOccurrence[T]{
						Points:    points,
						Direction: dir,
					})
				}
			}
		}
	}

	return occurrences
}

func searchFromCell[T comparable](grid [][]T, pattern []T, x, y, dx, dy int) map[T]Point {
	rows := len(grid)
	cols := len(grid[0])

	point := make(map[T]Point, rows)
	for k := 0; k < len(pattern); k++ {
		nx := x + k*dx
		ny := y + k*dy
		if nx < 0 || ny < 0 || nx >= rows || ny >= cols {
			return nil
		}
		if grid[nx][ny] != pattern[k] {
			return nil
		}
		point[pattern[k]] = Point{nx, ny}
	}

	return point
}

package data_structures

const (
	Rows = 3
	Cols = 4
)

type Puzzle struct {
	Buffer [Rows][Cols]byte
	BlankX uint8
	BlankY uint8
}

func NewPuzzle() Puzzle {
	return Puzzle{}
}

func (puzzle *Puzzle) IsSorted() bool {
	for i := 0; i < Rows; i++ {
		for j := 1; j < Cols; j++ {
			if puzzle.Buffer[i][j-1] > puzzle.Buffer[i][j] {
				return false
			}
			if j == Cols-1 && i < Rows-1 {
				if puzzle.Buffer[i][j] > puzzle.Buffer[i+1][0] {
					return false
				}
			}
		}
	}
	return true
}

func (puzzle *Puzzle) MoveBlank(direction Direction) {
	switch direction {
	case Up:
		x, y := puzzle.BlankX, puzzle.BlankY
		temp := puzzle.Buffer[x-1][y]
		puzzle.Buffer[x-1][y] = puzzle.Buffer[x][y]
		puzzle.Buffer[x][y] = temp
		puzzle.BlankX--
	case Down:
		x, y := puzzle.BlankX, puzzle.BlankY
		temp := puzzle.Buffer[x+1][y]
		puzzle.Buffer[x+1][y] = puzzle.Buffer[x][y]
		puzzle.Buffer[x][y] = temp
		puzzle.BlankX++
	case Right:
		x, y := puzzle.BlankX, puzzle.BlankY
		temp := puzzle.Buffer[x][y+1]
		puzzle.Buffer[x][y+1] = puzzle.Buffer[x][y]
		puzzle.Buffer[x][y] = temp
		puzzle.BlankY++
	case Left:
		x, y := puzzle.BlankX, puzzle.BlankY
		temp := puzzle.Buffer[x][y-1]
		puzzle.Buffer[x][y-1] = puzzle.Buffer[x][y]
		puzzle.Buffer[x][y] = temp
		puzzle.BlankY--
	}
}

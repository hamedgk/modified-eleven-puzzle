package data_structures

import "sort"

type PuzzleBuffer = [Rows][Cols]byte

type Puzzle struct {
	Buffer PuzzleBuffer
	BlankX uint8
	BlankY uint8
}

func FromBuffer(buffer PuzzleBuffer) Puzzle {
	for i, row := range buffer {
		for j, column := range row {
			if column == Blank {
				x, y := uint8(i), uint8(j)
				return Puzzle{buffer, x, y}
			}
		}
	}
	return Puzzle{}
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

func SortPuzzle(buffer PuzzleBuffer) PuzzleBuffer{
	var flattened []byte
	for _, row := range buffer {
		flattened = append(flattened, row[:]...)
	}

	sort.Slice(flattened, func(i, j int) bool {
		return flattened[i] < flattened[j]
	})

	sortedArray := [Rows][Cols]byte{}
	for i := 0; i < Rows; i++ {
		copy(sortedArray[i][:], flattened[i*Cols:(i+1)*Cols])
	}
	return sortedArray
}
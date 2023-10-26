package data_structures

type Direction = uint8

func (puzzle *Puzzle) possibleBlankMoves() []Direction {
	switch {
	case puzzle.BlankX == 0:
		switch {
		case puzzle.BlankY == 0:
			return []uint8{Right, Down}
		case puzzle.BlankY == Cols-1:
			return []uint8{Left, Down}
		default:
			return []uint8{Right, Left, Down}
		}
	case puzzle.BlankX == Rows-1:
		switch {
		case puzzle.BlankY == 0:
			return []uint8{Right, Up}
		case puzzle.BlankY == Cols-1:
			return []uint8{Left, Up}
		default:
			return []uint8{Right, Left, Up}
		}
	default:
		switch {
		case puzzle.BlankY == 0:
			return []uint8{Right, Down, Up}
		case puzzle.BlankY == Cols-1:
			return []uint8{Left, Down, Up}
		default:
			return []uint8{Right, Left, Down, Up}
		}

	}
}

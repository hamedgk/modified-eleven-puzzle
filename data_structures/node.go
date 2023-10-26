package data_structures

import "fmt"

type Node struct {
	Parent    *Node
	Direction Direction
	Puzzle    Puzzle
}

func (node *Node) Expand(queue Queue, explored map[PuzzleBuffer]bool) {
	possibleMoves := node.Puzzle.possibleBlankMoves()
	for _, direction := range possibleMoves {
		copyNode := *node
		copyNode.Direction = direction
		copyNode.Parent = node
		copyNode.Puzzle.MoveBlank(direction)
		if explored[copyNode.Puzzle.Buffer] {
			continue
		}
		queue.Enqueue(copyNode)
		explored[copyNode.Puzzle.Buffer] = true
	}
}

func (node *Node) IsGoal(buffer PuzzleBuffer) bool {
	return node.Puzzle.Buffer == buffer
}

func TraceBack(node Node) {
	if node.Parent == nil {
		fmt.Print("init")
		return
	}

	TraceBack(*node.Parent)
	switch node.Direction {
	case Up:
		fmt.Print(" -> Up")
	case Down:
		fmt.Print(" -> Down")
	case Right:
		fmt.Print(" -> Right")
	case Left:
		fmt.Print(" -> Left")
	default:
		return
	}
}

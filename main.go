package main

import (
	"eleven-puzzle/data_structures"
	"eleven-puzzle/data_structures/puzzle"
	"fmt"
)

var examplePuzzle = puzzle.PuzzleBuffer{
	{1, 2, 11, 4},
	{5, 255, 7, 8},
	{9, 3, 10, 11},
}

func main() {
	sortedArray := puzzle.SortPuzzle(examplePuzzle)

	explored := map[puzzle.PuzzleBuffer]bool{}
	frontier := data_structures.NewQueue()
	frontier.Enqueue(
		data_structures.Node{
			Parent:    nil,
			Direction: puzzle.None,
			Puzzle:    puzzle.FromBuffer(examplePuzzle),
		},
	)

	for {
		if frontier.IsEmpty() {
			fmt.Println("empty frontier")
			return
		}
		node, ok := frontier.Dequeue()
		if ok {
			if node.IsGoal(sortedArray) {
				data_structures.TraceBack(node)
				return
			}
			node.Expand(frontier, explored)
		}
	}
}

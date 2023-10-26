package main

import (
	"eleven-puzzle/data_structures"
	"fmt"
)

var examplePuzzle = data_structures.PuzzleBuffer{
	{1, 2, 11, 4},
	{5, 255, 7, 8},
	{9, 3, 10, 11},
}

func main() {
	sortedArray := data_structures.SortPuzzle(examplePuzzle)

	explored := map[data_structures.PuzzleBuffer]bool{}
	frontier := data_structures.NewQueue()
	frontier.Enqueue(
		data_structures.Node{
			Parent:    nil,
			Direction: data_structures.None,
			Puzzle:    data_structures.FromBuffer(examplePuzzle),
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

package main

import (
	"eleven-puzzle/data_structures"
	"fmt"
)

const (
	rows = data_structures.Rows
	cols = data_structures.Cols
)

var examplePuzzle = [rows][cols]byte{
	//{1, 2, 255, 1},
	//{4, 4, 3, 8},
	//{6, 7, 7, 9},
	{1, 2, 11, 4},
	{5, 255, 7, 8},
	{9, 3, 10, 11},
}

func main() {
	sortedArray := data_structures.SortPuzzle(examplePuzzle)

	explored := map[[rows][cols]byte]bool{}
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

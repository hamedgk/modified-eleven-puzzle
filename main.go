package main

import (
	"eleven-puzzle/data_structures"
	"fmt"
	"sort"
)

const (
	rows = data_structures.Rows
	cols = data_structures.Cols
)

var examplePuzzle = [rows][cols]byte{
	//{1, 2, 255, 1},
	//{4, 4, 3, 8},
	//{6, 7, 7, 9},
	{1, 2, 255, 4},
	{5, 6, 7, 8},
	{9, 3, 10, 11},
}

func main() {
	// Flatten the 3x4 array into a 1D slice
	var flattened []byte
	for _, row := range examplePuzzle {
		flattened = append(flattened, row[:]...)
	}

	// Sort the flattened slice
	sort.Slice(flattened, func(i, j int) bool {
		return flattened[i] < flattened[j]
	})

	// Convert the sorted 1D slice back to a 3x4 array
	sortedArray := [rows][cols]byte{}
	for i := 0; i < rows; i++ {
		copy(sortedArray[i][:], flattened[i*cols:(i+1)*cols])
	}

	frontier := data_structures.NewQueue()
	frontier.Enqueue(data_structures.Node{
		Parent:    nil,
		Direction: data_structures.None,
		Puzzle:    data_structures.Puzzle{Buffer: examplePuzzle, BlankX: 0, BlankY: 2},
	})

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
			node.Expand(frontier)
		}
	}
}

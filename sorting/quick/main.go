package quick

import (
	"fmt"
)

type dividor struct {
	pivot int
	lte   []int
	gt    []int
}

func (div *dividor) reassemble() []int {
	output := []int{}

	output = append(output, div.lte...)

	output = append(output, div.pivot)

	output = append(output, div.gt...)

	return output
}

func (div dividor) String() string {
	return fmt.Sprintf("pivot: %v \nlte: %v \ngt: %v \n", div.pivot, div.lte, div.gt)
}

// Sort run quck sort on array of int
func Sort(arr []int) []int {
	return quickSort(arr)
}

func quickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	div := dividor{pivot: arr[len(arr)-1]}

	// TODO -- enhancement rather than appending to two different arrays
	// append to a single array keeping track of the pivot
	for _, value := range arr[:len(arr)-1] {
		if value <= div.pivot {
			div.lte = append(div.lte, value)
		} else {
			div.gt = append(div.gt, value)
		}
	}

	div.gt = quickSort(div.gt)
	div.lte = quickSort(div.lte)

	return div.reassemble()
}

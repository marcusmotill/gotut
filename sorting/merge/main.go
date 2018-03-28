package merge

// Sort run merge sort on array of int
func Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	input := [][]int{}
	input = append(input, arr)

	unsortedSplit := split(arr, input)
	sortedSplit := sortSplit(unsortedSplit)

	sorted := merge(sortedSplit)

	return sorted
}

// chunk up the input array to size of 2 chunks
func split(arr []int, slicer [][]int) [][]int {

	output := [][]int{}

	for _, value := range slicer {
		chunkSize := len(value) / 2

		if chunkSize < 2 {
			chunkSize = 2
		}

		split1, split2 := value[0:chunkSize], value[chunkSize:]

		output = append(output, split1)

		if len(split2) > 0 {
			output = append(output, split2)
		}
	}

	desiredLength := len(arr) / 2

	if len(arr)%2 != 0 {
		desiredLength++
	}

	if len(output) >= desiredLength {
		return output
	}

	return split(arr, output)
}

// sort the size of 2 chunks chunk by chunk
func sortSplit(unsortedSplit [][]int) [][]int {
	output := [][]int{}
	for _, value := range unsortedSplit {
		if len(value) == 1 {
			output = append(output, value)
		} else {
			val1, val2 := value[0], value[1]

			if val1 > val2 {
				output = append(output, []int{val2, val1})
			} else {
				output = append(output, value)
			}
		}
	}

	return output
}

//chunk the sorted splits by 2 again, sort agaisnt eachother
func merge(sortedSplit [][]int) []int {

	if len(sortedSplit) == 1 {
		return sortedSplit[0]
	}

	chunkSize := len(sortedSplit) / 2

	if chunkSize < 2 {
		chunkSize = 2
	}

	output := [][]int{}

	for i := 0; i < len(sortedSplit); i += 2 {

		if i+1 >= len(sortedSplit) {
			output = append(output, sortedSplit[i])
			continue
		}

		chunk := make([][]int, 2)

		chunk[0] = sortedSplit[i]
		chunk[1] = sortedSplit[i+1]

		output = append(output, mergeArrs(chunk[0], chunk[1]))
	}

	return merge(output)
}

//sort two arrays
func mergeArrs(arr1 []int, arr2 []int) []int {
	merged := []int{}
	i := 0
	j := 0

	for i < len(arr1) || j < len(arr2) {
		if i >= len(arr1) {
			merged = append(merged, arr2[j])
			j++
		} else if j >= len(arr2) {
			merged = append(merged, arr1[i])
			i++
		} else {
			val1 := arr1[i]
			val2 := arr2[j]

			if val1 <= val2 {
				merged = append(merged, val1)
				i++
			} else {
				merged = append(merged, val2)
				j++
			}
		}
	}

	return merged
}

package insertion

func Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	arr = insertionSort(arr, 0)

	return arr
}

func insertionSort(arr []int, pos int) []int {

	if pos >= len(arr) {
		return arr
	}

	if pos == 0 {
		return insertionSort(arr, pos+1)
	}

	posVal := arr[pos]

	for i := pos - 1; i >= 0; i-- {
		refVal := arr[i]

		if refVal >= posVal {
			if i-1 >= 0 {
				nextRef := arr[i-1]
				if nextRef >= posVal {
					continue
				}
			}
			arr = move(arr, pos, i)
			break
		}
	}

	pos++

	return insertionSort(arr, pos)
}

func move(arr []int, from int, to int) []int {
	diff := from - to
	fromVal := arr[from]

	for i := 0; i <= diff; i++ {
		if i == diff {
			arr[from-i] = fromVal
		} else {
			arr[from-i] = arr[from-(i+1)]
		}
	}

	return arr
}

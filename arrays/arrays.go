package arrays

import (
	"math/rand"
	"time"
)

func Generate(x int, options ...int) (arr []int) {
	arr = make([]int, x)
	max := options[0]
	if max == 0 {
		max = 10000
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := range arr {
		arr[i] = r1.Intn(max)
	}

	return
}

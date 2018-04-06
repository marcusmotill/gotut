package main

import (
	"fmt"
	"gotut/arrays"
	"gotut/sorting/heap"
	"gotut/sorting/insertion"
	"gotut/sorting/merge"
	"gotut/sorting/quick"
	"math/rand"
	"sort"
	"time"
)

type sortFunc func([]int) []int

func main() {
	test(quick.Sort, "quick")
	test(merge.Sort, "merge")
	test(heap.Sort, "heap")
	test(insertion.Sort, "insertion")
}

func test(sorter sortFunc, name string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 1; i <= 101; i++ {
		arr := arrays.Generate(r1.Intn(101), r1.Intn(101))
		tmp := make([]int, len(arr))
		copy(tmp, arr)

		fmt.Printf("Running %s sort array length %d: ", name, len(arr))
		outArr := sorter(tmp)
		fmt.Print("verifying... ")
		sort.Ints(arr)

		var err = ""
		for i, value := range arr {
			if value != outArr[i] {
				err = fmt.Sprintf("error:\nGot:\n%v\n\nExpected:\n%v", outArr, arr)
				break
			}
		}

		if len(err) > 0 {
			fmt.Println(err)
			break
		} else {
			fmt.Print("passed\n")
		}
	}

	fmt.Printf("\n---- Done testing %s ----\n\n", name)
}

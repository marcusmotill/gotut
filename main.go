package main

import (
	"fmt"
	"gotut/arrays"
	"gotut/sorting/heap"
	"math/rand"
	"sort"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 1; i <= 100; i++ {
		arr := arrays.Generate(r1.Intn(1000), r1.Intn(1000))
		tmp := make([]int, len(arr))
		copy(tmp, arr)

		fmt.Printf("Running heap sort array length %d: ", len(arr))
		outArr := heap.Sort(tmp)
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

}

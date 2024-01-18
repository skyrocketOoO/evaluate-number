package main

import (
	"fmt"
	"math"

	"github.com/mohae/deepcopy"
)

func main() {
	fmt.Println(Evaluate_numbers(5, 10, Distribution{
		[]int{90, 10},
	}))
}

// rounding

type Distribution struct {
	Proportion []int
}

func Evaluate_numbers(lowerBound int, upperBound int, distribution Distribution) [][]int {
	propMap := make(map[int]bool, len(distribution.Proportion))
	for _, v := range distribution.Proportion {
		propMap[v] = true
	}

	var result [][]int
	for n := lowerBound; n <= upperBound; n++ {
		composition := []int{}
		copyPropMap := deepcopy.Copy(propMap).(map[int]bool)
		curMembers := 0
		for curMembers < n {
			find := false
			for i := 0; i < n; i++ {
				p := int(math.Round(float64(i*100) / float64(n)))
				if _, ok := copyPropMap[p]; ok {
					delete(copyPropMap, p)
					composition = append(composition, i)
					find = true
					break
				}
			}
			if !find {
				break
			}
		}
		sum := 0
		for _, v := range composition {
			sum += v
		}
		if sum == n {
			result = append(result, composition)
		}
	}
	return result
}

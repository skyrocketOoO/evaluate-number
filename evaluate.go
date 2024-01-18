package main

import (
	"fmt"
	"math"

	"github.com/mohae/deepcopy"
)

type Distribution struct {
	Proportion []int
}

func Evaluate_numbers(lowerBound int, upperBound int, distributions ...Distribution) []int {
	var possibleResult [][][]int
	for _, distribution := range distributions {
		possibleResult = append(possibleResult, evaluate_numbers(lowerBound, upperBound, distribution))
	}

	fmt.Println(possibleResult)
	finalResult := Intersect(possibleResult...)
	return finalResult
}

func evaluate_numbers(lowerBound int, upperBound int, distribution Distribution) [][]int {
	propMap := make(map[int]int, len(distribution.Proportion))
	for _, v := range distribution.Proportion {
		if _, ok := propMap[v]; !ok {
			propMap[v] = 0
		}
		propMap[v]++
	}

	var result [][]int
	for n := lowerBound; n <= upperBound; n++ {
		composition := []int{}
		copyPropMap := deepcopy.Copy(propMap).(map[int]int)
		curMembers := 0
		for curMembers < n {
			find := false
			for i := 0; i < n; i++ {
				p := int(math.Round(float64(i*100) / float64(n)))
				if _, ok := copyPropMap[p]; ok {
					copyPropMap[p]--
					if v := copyPropMap[p]; v == 0 {
						delete(copyPropMap, p)
					}
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

func Intersect(allSlices ...[][]int) []int {
	if len(allSlices) == 0 {
		return nil
	}

	countMap := make(map[int]int)

	for _, slices := range allSlices {
		for _, slice := range slices {
			sum := 0
			for _, v := range slice {
				sum += v
			}
			countMap[sum]++
		}
	}

	// Find slices that occurred in all slices
	var common []int
	for k, c := range countMap {
		if c == len(allSlices) {
			common = append(common, k)
		}
	}

	return common
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Evaluate_numbers(5, 10,
		Distribution{
			[]int{33, 67},
		},
		Distribution{
			[]int{17, 17, 17, 50},
		},
	))
}

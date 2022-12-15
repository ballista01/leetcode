package main

import "fmt"

func main() {
	// getNumWitnesses([]int{3, 6, 3, 4, 1})
	str := "bad"
	fmt.Println(str[2] - 'a')
}

// func getNumWitnesses(heights []int) int {
// 	numWitness := 0
// 	maxHeight := heights[len(heights)-1]
// 	for i := len(heights) - 1; i >= 0; i-- {
// 		if heights[i] >= maxHeight {
// 			maxHeight = heights[i]
// 			numWitness++
// 		}
// 	}
// 	return numWitness
// }

package main

import "fmt"

func findPairs(arr []int, target int) [][2]int {
	seen := make(map[int]bool)
	pairs := make(map[[2]int]bool)
	result := make([][2]int, 0)

	for _, num := range arr {
		complement := target - num
		if seen[complement] {
			a, b := num, complement
			if a > b {
				a, b = b, a
			}
			pair := [2]int{a, b}
			if !pairs[pair] {
				result = append(result, pair)
				pairs[pair] = true
			}
		}
		seen[num] = true
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 3, 5, -1, 0}
	target := 4
	pairs := findPairs(arr, target)

	fmt.Println("Pairs that sum to", target, ":")
	for _, pair := range pairs {
		fmt.Println(pair)
	}
}
package main

import (
	"fmt"
)

var level int = -1

func print(slice []int, size int) {
	if slice != nil {
		for i := 0; i < size; i++ {
			fmt.Printf("%d ", slice[i])
		}
		fmt.Printf("\n")
	}
}

func visit(slice []int, N, k int) {
	level = level + 1
	slice[k] = level
	if level == N {
		print(slice, N)
	} else {
		for i := 0; i < N; i++ {
			if slice[i] == 0 {
				visit(slice, N, i)
			}
		}
	}
	level = level - 1
	slice[k] = 0
}

func main() {
	N := 4
	slice := make([]int, N)
	for i := 0; i < N; i++ {
		slice[i] = 0
	}

	fmt.Println(slice)
	visit(slice, N, 0)
}

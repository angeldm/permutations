package main

import (
	"fmt"
)

func print(slice []int, size int) {
	if slice != nil {
		for i := 0; i < size; i++ {
			fmt.Printf("%d ", slice[i])
		}
		fmt.Printf("\n")
	}
}

func swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func rotateLeft(slice []int, start, n int) {
	tmp := slice[start]
	for i := start; i < n-1; i++ {
		slice[i] = slice[i+1]
	}
	slice[n-1] = tmp
}

func permute(slice []int, start, n int) {
	print(slice, n)
	if start < n {
		var i, j int
		for i = n - 2; i >= start; i-- {
			for j = i + 1; j < n; j++ {
				swap(slice, i, j)
				permute(slice, i+1, n)
			}
			rotateLeft(slice, i, n)
		}
	}
}

func main() {
	N := 4
	slice := make([]int, N)
	for i := 0; i < N; i++ {
		slice[i] = 0
	}

	permute(slice, 0, N)
}

package main

import (
	"sort"
	"fmt"
)

// Generate the first permutation of data.
func PermutationFirst(data sort.Interface) {
	sort.Sort(data)
}

// Generate the next permutation of data if possible and return true.
// If there is no more permutation left return false.
// Based on the algorithm described here:
// http://en.wikipedia.org/wiki/Permutation#Systematic_generation_of_all_permutations 
func PermutationNext(data sort.Interface) bool {
	var k, l int
	for k = data.Len() - 2; ; k-- { // 1.
		if k < 0 {
			return false
		}

		if data.Less(k, k+1) {
			break
		}
	}
	for l = data.Len() - 1; !data.Less(k, l); l-- { // 2.
	}
	data.Swap(k, l)                             // 3.
	for i, j := k+1, data.Len()-1; i < j; i++ { // 4.
		data.Swap(i, j)
		j--
	}
	return true
}

func main() {
	slice := make(sort.IntSlice,4)
	slice[0]=1
	slice[1]=3
	slice[2]=4
	slice[3]=2

	fmt.Println(slice)
	PermutationFirst(slice)
	fmt.Println(slice)
	for PermutationNext(slice) {
		fmt.Println(slice)
	}
}

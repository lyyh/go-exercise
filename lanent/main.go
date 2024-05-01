package main

import (
	"fmt"
)

func Union[T comparable](slices ...[]T) []T {
	result := []T{}
	contain := map[T]struct{}{}

	for _, slice := range slices {
		for _, item := range slice {
			if _, ok := contain[item]; !ok {
				contain[item] = struct{}{}
				result = append(result, item)
			}
		}
	}

	return result
}
func main() {
	a := []int{1, 2, 3}
	b := []int{3, 4, 5}
	fmt.Printf("%+v", Union(a, b))
}

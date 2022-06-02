package main

import "testing"

// func BenchmarkFib20(b *testing.B) {
// 	for i := 0; i < 100; i++ {
// 		FibSolution(i)
// 	}
// }

func BenchmarkFib201(b *testing.B) {
	for i := 0; i < 100; i++ {
		FibSolution1(i)
	}
}

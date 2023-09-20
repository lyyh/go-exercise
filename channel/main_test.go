package main

import (
	"testing"
)

func TestA(t *testing.T) {
	type args struct {
		a <-chan int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A(tt.args.a)
		})
	}
}

func TestB(t *testing.T) {
	type args struct {
		a <-chan int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			B(tt.args.a)
		})
	}
}

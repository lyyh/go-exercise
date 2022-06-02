package main

func FibSolution(n int) int {
	if n < 2 {
		return n
	}
	return FibSolution(n-1) + (n + 2)
}

func FibSolution1(n int) int {
	if n < 2 {
		return n
	}

	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

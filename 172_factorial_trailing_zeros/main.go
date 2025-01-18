package main

func trailingZeroes(n int) int {
	a := n / 5
	b := n / 25
	c := n / 125
	d := n / 625
	e := n / 3125
	return a + b + c + d + e
}

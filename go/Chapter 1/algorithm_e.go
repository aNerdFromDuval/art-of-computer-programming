/*
Euclid's Algorithm
Compute the greatest common divisor of two integers

Algorithm:
 - E0: [Ensure m >= n] If m < n, exchange m and n
    - E1: [Find the Remainder] Divide m by n and let r be the remainder. (We have m = nq + r, where q is the quotient and r is the remainder)
    - E2: [Is it zero?] If r = 0, the algorithm terminates; n is the answer
    - E3: [Reduce] Set m <- n, n <- r, and go back to step E1

*/

package main

import "fmt"

func gcd(m, n int) int {
	if m < n {
		m, n = n, m
	}
	for {
		r := m % n
		if r == 0 {
			return n
		}
		m, n = n, r
	}
}

func gcd_recursive(m, n int) int {
	if m < n {
		m, n = n, m
	}
	if r := m % n; r == 0 {
		return n
	} else {
		return gcd_recursive(n, r)
	}
}

func main() {
	fmt.Println(gcd(1989, 1590))
	fmt.Println(gcd_recursive(1989, 1590))
}

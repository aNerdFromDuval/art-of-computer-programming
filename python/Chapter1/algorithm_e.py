# pyre-strict

"""
Euclid's Algorithm
Compute the greatest common divisor of two integers

Algorithm:
 - E0: [Ensure m >= n] If m < n, exchange m and n
    - E1: [Find the Remainder] Divide m by n and let r be the remainder. (We have m = nq + r, where q is the quotient and r is the remainder)
    - E2: [Is it zero?] If r = 0, the algorithm terminates; n is the answer
    - E3: [Reduce] Set m <- n, n <- r, and go back to step E1


Input: Two integers
Output: The greatest common divisor of the two integers
"""


def algorithm_e(m: int, n: int) -> int:
    if m < n:
        m, n = n, m
    while True:
        r: int = m % n
        if r == 0:
            return n
        m, n = n, r


def algorithm_e_recursive(m: int, n: int) -> int:
    if m < n:
        m, n = n, m
    r: int = m % n
    if r == 0:
        return n
    return algorithm_e_recursive(n, r)

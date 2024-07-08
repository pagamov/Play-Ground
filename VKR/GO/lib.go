package main

import "math/big"

func GCD(m int, n int) int {
	var mult int = 1
	if m > n {
		m = m % n
	} else if n > m {
		n = n % m
	}
	for {
		if m == 0 || n == 0 || n == m {
			return mult * max(n, m)
		}
		if m == 1 || n == 1 {
			return mult
		}
		mm2 := m % 2
		nm2 := n % 2
		if mm2 == 0 && nm2 == 0 {
			mult *= 2
			m = m / 2
			n = n / 2
		} else if mm2 == 0 && nm2 != 0 {
			m = m / 2
		} else if mm2 != 0 && nm2 == 0 {
			n = n / 2
		} else if mm2 != 0 && nm2 != 0 {
			if n > m {
				piv := (n - m) / 2
				n = m
				m = piv
			} else if n < m {
				m = (m - n) / 2
			}
		}
	}
}

type Q struct {
	n int
	m int
}

func (q *Q) create(n big.Int) {
	q.n = n
	q.m = int(bigin)
}

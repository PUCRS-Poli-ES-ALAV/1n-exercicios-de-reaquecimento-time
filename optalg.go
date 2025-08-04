package optalg

// 1
func Multiply(num, times int) int {
	switch times {
	case 0:
		return 0
	case 1:
		return num
	default:
		return num + Multiply(num, times-1)
	}
}

// 2
func Increment(n1, n2 int) int {
	switch {
	case n1 == 0 && n2 == 0:
		return 0
	case n1 == 0:
		return 1 + Increment(0, n2-1)
	default:
		return 1 + Increment(n1-1, n2)
	}
}

// 3
func Series(n int) float64 {
	switch n {
	case 0:
		panic("invalid division by zero")
	case 1:
		return 1
	default:
		return 1/float64(n) + Series(n-1)
	}
}

// 4
func ReverseString(s string) string {
	switch {
	case len(s) == 0:
		return ""
	default:
		return string(s[len(s)-1]) + ReverseString(s[:len(s)-1])

	}
}

// 5
func Ackerman(m, n int) int {
	switch {
	case m == 0:
		return n + 1
	case n == 0:
		return Ackerman(m-1, 1)
	default:
		return Ackerman(m-1, Ackerman(m, n-1))
	}
}

// 6
func VetSumProduct(v []int) (int, int) {
	switch len(v) {
	case 0:
		return 0, 1
	default:
		n, m := VetSumProduct(v[1:])
		return v[0] + n, v[0] * m
	}
}

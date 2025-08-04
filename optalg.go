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
func Sequence(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	default:
		return 2 * Sequence(n-1) * Sequence(n-2)
	}
}

// 6
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

// 7
func VetSumProduct(v []int) (int, int) {
	switch len(v) {
	case 0:
		return 0, 1
	default:
		n, m := VetSumProduct(v[1:])
		return v[0] + n, v[0] * m
	}
}

// 8
func IsPalindrome(s string) bool {
	switch len(s) {
	case 0:
		return false
	case 1:
		return true
	case 2:
		return s[0] == s[1]
	default:
		return (s[0] == s[len(s)-1]) && IsPalindrome(s[1:len(s)-1])
	}
}

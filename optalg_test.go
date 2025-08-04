package optalg_test

import (
	"testing"

	"optalg"
)

func TestMutiply(t *testing.T) {
	testCases := []struct {
		num, times int
		want       int
	}{
		{1234, 0, 0},
		{6, 4, 24},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := optalg.Multiply(tc.num, tc.times)
			if got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func TestIncrement(t *testing.T) {
	testCases := []struct {
		n1, n2 int
		want   int
	}{
		{3, 2, 5},
		{2, 2, 4},
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1000, 999, 1999},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := optalg.Increment(tc.n1, tc.n2)
			if got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func TestSeries(t *testing.T) {
	testCases := []struct {
		n    int
		want float64
	}{
		{2, 1 + float64(1)/2},
		{4, 1 + float64(1)/2 + float64(1)/3 + float64(1)/4},
		{15, 1 + float64(1)/2 + float64(1)/3 + float64(1)/4 + float64(1)/5 + float64(1)/6 + float64(1)/7 + float64(1)/8 + float64(1)/9 + float64(1)/10 + float64(1)/11 + float64(1)/12 + float64(1)/13 + float64(1)/14 + float64(1)/15},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := optalg.Series(tc.n)
			if got != tc.want {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	testCases := []struct {
		s    string
		want string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"abcdefghijk", "kjihgfedcba"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := optalg.ReverseString(tc.s)
			if got != tc.want {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}

package ch15

import (
	"strings"
	"testing"

	"cp/internal/check"
)

func TestEx15_1(t *testing.T) {
	check.Eq(t, SumFscan(strings.NewReader("3\n10 20 30")), 60)
}

func TestEx15_2(t *testing.T) {
	check.Eq(t, SquaresOutput(3), "1 1\n2 4\n3 9\n")
}

func TestEx15_3(t *testing.T) {
	check.Eq(t, Evens(9), []int{0, 2, 4, 6, 8})
	check.Eq(t, cap(Evens(1000)), 501)
}

func TestEx15_4(t *testing.T) {
	check.Eq(t, CountsSlice([]int{2, 0, 2, 3}, 3), []int{1, 0, 2, 1})
}

func TestEx15_5(t *testing.T) {
	check.Eq(t, SumFlat([]int{1, 2, 3, 4, 5, 6}, 2, 3), 21)
}

func TestEx15_6(t *testing.T) {
	check.Eq(t, []string{Repeat("ab", 3), Repeat("x", 0)}, []string{"ababab", ""})
}

func TestEx15_7(t *testing.T) {
	start, to := CSR(3, [][2]int{{0, 1}, {0, 2}, {2, 1}})
	check.Eq(t, []any{start, to}, []any{[]int{0, 2, 2, 3}, []int{1, 2, 1}})
}

func TestEx15_8(t *testing.T) {
	check.Eq(t, []string{Verdict(200_000, "nlogn"), Verdict(20_000, "n2"), Verdict(100_000, "n2")}, []string{"fast", "tight", "too slow"})
}

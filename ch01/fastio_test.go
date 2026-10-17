package ch01

import (
	"strings"
	"testing"

	"cp/internal/check"
)

func TestEx1_1(t *testing.T) {
	check.Eq(t, ReadInts(strings.NewReader("3 -4\n10")), []int{3, -4, 10})
	long := strings.Repeat("7 ", 100000)
	check.Eq(t, len(ReadInts(strings.NewReader(long))), 100000)
}

func TestEx1_2(t *testing.T) {
	f := NewFastReader(strings.NewReader("3\n10 -20\n30"))
	check.Eq(t, []int{f.Int(), f.Int(), f.Int(), f.Int()}, []int{3, 10, -20, 30})
}

func TestEx1_3(t *testing.T) {
	check.Eq(t, ReadArray("4\n5 1 9 2"), []int{5, 1, 9, 2})
	check.Eq(t, ReadArray("3\n1\n\n2 3"), []int{1, 2, 3})
}

func TestEx1_4(t *testing.T) {
	check.Eq(t, ReadGrid("2 3\n#.#\n..#"), [][]byte{[]byte("#.#"), []byte("..#")})
}

func TestEx1_5(t *testing.T) {
	check.Eq(t, SumCases("2\n1 2\n1000000000000 1"), "3\n1000000000001\n")
	check.Eq(t, SumCases("1\n1000000000000000000 -1"), "999999999999999999\n")
}

func TestEx1_6(t *testing.T) {
	check.Eq(t, []string{Fmt(2.0 / 3), FmtRow("ann", 42)}, []string{"0.666667", "ann          42"})
}

func TestEx1_7(t *testing.T) {
	check.Eq(t, []string{JoinLine([]int{3, 1, 2}), JoinLine(nil), JoinLine([]int{-5})}, []string{"3 1 2", "", "-5"})
}

func TestEx1_8(t *testing.T) {
	check.Eq(t, []string{Run("5\n3 9 -1 4 7"), Run("2\n-3 -9")}, []string{"9\n", "-3\n"})
}

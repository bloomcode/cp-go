package ch05

import (
	"testing"

	"cp/internal/check"
)

func TestEx5_1(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := BumpArray(a)
	check.Eq(t, []int{a[0], b[0]}, []int{1, 2})
	s := []int{1, 2, 3}
	BumpSlice(s)
	check.Eq(t, s[0], 2)
}

func TestEx5_2(t *testing.T) {
	x, y := AliasTrap()
	check.Eq(t, []int{x, y}, []int{5, 5})
	x, y = AliasFixed()
	check.Eq(t, []int{x, y}, []int{4, 5})
}

func TestEx5_3(t *testing.T) {
	g := [][]int{{1, 2}, {3, 4}}
	d := Clone2D(g)
	d[0][0] = 9
	check.Eq(t, []any{g[0][0], d}, []any{1, [][]int{{9, 2}, {3, 4}}})
}

func TestEx5_4(t *testing.T) {
	grid := MakeGrid(2, 3, 7)
	check.Eq(t, grid, [][]int{{7, 7, 7}, {7, 7, 7}})
	grid[0] = append(grid[0], 8)
	check.Eq(t, grid[1][0], 7) // rows don't share storage
}

func TestEx5_5(t *testing.T) {
	p := PrefixSums([]int{3, 1, 4, 1, 5})
	check.Eq(t, p, []int{0, 3, 4, 8, 9, 14})
	check.Eq(t, RangeSum([]int{0, 3, 4, 8, 9, 14}, 1, 3), 6)
}

func TestEx5_6(t *testing.T) {
	check.Eq(t, RemoveAt([]int{1, 2, 9, 3}, 2), []int{1, 2, 3})
	orig := []int{1, 2, 9, 3}
	check.Eq(t, RemoveAtCopy(orig, 2), []int{1, 2, 3})
	check.Eq(t, orig, []int{1, 2, 9, 3})
}

func TestEx5_7(t *testing.T) {
	check.Eq(t, RotateRight([]int{1, 2, 3, 4, 5}, 2), []int{4, 5, 1, 2, 3})
	check.Eq(t, RotateRight([]int{1, 2, 3}, 7), []int{3, 1, 2})
}

func TestEx5_8(t *testing.T) {
	in := []int{3, -1, 4, 0, 5}
	out := KeepPositive(in)
	check.Eq(t, out, []int{3, 4, 5})
	check.True(t, &out[0] == &in[0], "same backing array")
}

func TestEx5_9(t *testing.T) {
	check.Eq(t, Spiral([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})
	check.Eq(t, Spiral([][]int{{1, 2}, {3, 4}, {5, 6}}), []int{1, 2, 4, 6, 5, 3})
	check.Eq(t, Spiral([][]int{{1, 2, 3}}), []int{1, 2, 3})
}

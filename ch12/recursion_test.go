package ch12

import (
	"testing"

	"cp/internal/check"
)

func TestEx12_1(t *testing.T) {
	check.Eq(t, []int{Power(3, 13), Power(2, 0)}, []int{1594323, 1})
}

func TestEx12_2(t *testing.T) {
	check.Eq(t, []int{Fib(10), Fib(90)}, []int{55, 2880067194370816120})
}

func TestEx12_3(t *testing.T) {
	check.Eq(t, []int{GridPaths(3, 3), GridPaths(17, 17)}, []int{6, 601080390})
}

func TestEx12_4(t *testing.T) {
	check.Eq(t, Subsets([]int{1, 2}), [][]int{{}, {2}, {1}, {1, 2}})
	check.Eq(t, len(Subsets([]int{1, 2, 3, 4})), 16)
	check.Eq(t, Subsets([]int{1, 2, 3})[7], []int{1, 2, 3})
}

func TestEx12_5(t *testing.T) {
	check.Eq(t, Permutations("abc"), []string{"abc", "acb", "bac", "bca", "cab", "cba"})
}

func TestEx12_6(t *testing.T) {
	check.Eq(t, CountIslands([][]byte{[]byte("#.#"), []byte("#.."), []byte("..#")}), 3)
}

func TestEx12_7(t *testing.T) {
	check.Eq(t, DeepSum(1_000_000), 500000500000)
}

func TestEx12_8(t *testing.T) {
	check.Eq(t, []int{NQueens(4), NQueens(8)}, []int{2, 92})
}

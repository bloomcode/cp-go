package ch16

import (
	"math"
	"testing"

	"cp/internal/check"
)

func TestEx16_1(t *testing.T) {
	check.Eq(t, SolveAll("2\n3\n1 2 3\n1\n-5\n"), "6\n-5\n")
}

func TestEx16_2(t *testing.T) {
	check.Eq(t, CountPairsAtMost([]int{1, 5, 3, 2}, 5), 3)
}

func TestEx16_3(t *testing.T) {
	check.Eq(t, []int{MinCapacity([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5), MinCapacity([]int{3, 2, 2, 4, 1, 4}, 3)}, []int{15, 6})
}

func TestEx16_4(t *testing.T) {
	check.Eq(t, ApplyRanges(5, [][3]int{{1, 3, 2}, {0, 1, 5}}), []int{5, 7, 2, 2, 0})
}

func TestEx16_5(t *testing.T) {
	check.Eq(t, Compress([]int{100, -5, 100, 42}), []int{2, 0, 2, 1})
}

func TestEx16_6(t *testing.T) {
	d := NewDSU(5)
	check.True(t, d.Union(0, 1) && d.Union(3, 4) && !d.Union(1, 0), "unions")
	check.Eq(t, d.Components, 3)
	check.True(t, d.Find(0) == d.Find(1) && d.Find(0) != d.Find(3), "find")
}

func TestEx16_7(t *testing.T) {
	check.Eq(t, Dijkstra(4, [][3]int{{0, 1, 4}, {0, 2, 1}, {2, 1, 2}, {1, 3, 5}}, 0), []int{0, 3, 1, 8})
	check.Eq(t, Dijkstra(2, nil, 0)[1], math.MaxInt)
}

func TestEx16_8(t *testing.T) {
	check.Eq(t, DepthsFromRoot(5, [][2]int{{0, 1}, {0, 2}, {1, 3}, {3, 4}}), []int{0, 1, 1, 2, 3})
	n := 1_000_000
	path := make([][2]int, n-1)
	for i := range path {
		path[i] = [2]int{i, i + 1}
	}
	check.Eq(t, DepthsFromRoot(n, path)[n-1], n-1)
}

func TestEx16_9(t *testing.T) {
	check.Eq(t, []int{MinCoins([]int{1, 3, 4}, 6), MinCoins([]int{2}, 3)}, []int{2, -1})
}

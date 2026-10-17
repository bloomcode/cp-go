package ch07

import (
	"testing"

	"cp/internal/check"
)

func TestEx7_1(t *testing.T) {
	in := []int{3, 1, 2}
	check.Eq(t, SortedCopy(in), []int{1, 2, 3})
	check.Eq(t, SortedDesc(in), []int{3, 2, 1})
	check.Eq(t, in, []int{3, 1, 2})
}

func TestEx7_2(t *testing.T) {
	check.Eq(t, Leaderboard([]Player{{"bo", 80}, {"al", 95}, {"cy", 80}}), []Player{{"al", 95}, {"bo", 80}, {"cy", 80}})
}

func TestEx7_3(t *testing.T) {
	names := []string{}
	for _, p := range ByDeptThenAge([]Person{{"a", "ops", 40}, {"b", "dev", 30}, {"c", "ops", 25}}) {
		names = append(names, p.Name)
	}
	check.Eq(t, names, []string{"b", "c", "a"})
}

func TestEx7_4(t *testing.T) {
	check.Eq(t, Argsort([]int{30, 10, 20}), []int{1, 2, 0})
}

func TestEx7_5(t *testing.T) {
	sorted := []int{1, 3, 3, 5, 8}
	check.Eq(t, []int{CountInRange(sorted, 3, 5), CountInRange(sorted, 6, 7), CountInRange(sorted, 0, 100)}, []int{3, 0, 5})
}

func TestEx7_6(t *testing.T) {
	check.Eq(t, []int{IntSqrt(27), IntSqrt(36), IntSqrt(0), IntSqrt(1), IntSqrt(1_000_000_000_000_000_000)}, []int{5, 6, 0, 1, 1_000_000_000})
}

func TestEx7_7(t *testing.T) {
	check.Eq(t, []string{LargestNumber([]int{3, 30, 34, 5, 9}), LargestNumber([]int{0, 0})}, []string{"9534330", "0"})
}

func TestEx7_8(t *testing.T) {
	check.Eq(t, MergeIntervals([][2]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}}), [][2]int{{1, 6}, {8, 10}, {15, 18}})
	check.Eq(t, MergeIntervals([][2]int{{1, 4}, {4, 5}}), [][2]int{{1, 5}})
}

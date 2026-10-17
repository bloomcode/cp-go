package ch06

import (
	"fmt"
	"testing"

	"cp/internal/check"
)

func set(xs ...int) map[int]struct{} {
	m := map[int]struct{}{}
	for _, x := range xs {
		m[x] = struct{}{}
	}
	return m
}

func TestEx6_1(t *testing.T) {
	check.Eq(t, WordCounts("To be or not to be"), map[string]int{"to": 2, "be": 2, "or": 1, "not": 1})
}

func TestEx6_2(t *testing.T) {
	m := map[string]int{"a": 0}
	v, ok := Lookup(m, "a")
	check.Eq(t, []any{v, ok}, []any{0, true})
	v, ok = Lookup(m, "b")
	check.Eq(t, []any{v, ok}, []any{0, false})
}

func TestEx6_3(t *testing.T) {
	check.Eq(t, Unique([]int{3, 1, 3, 2, 1}), []int{3, 1, 2})
}

func TestEx6_4(t *testing.T) {
	check.Eq(t, fmt.Sprint(GroupByFirst([]string{"apple", "bat", "avocado"})), "map[97:[apple avocado] 98:[bat]]")
}

func TestEx6_5(t *testing.T) {
	check.Eq(t, SortedKeys(map[string]int{"c": 1, "a": 2, "b": 3}), []string{"a", "b", "c"})
	for i := 0; i < 20; i++ { // map order is random; the report must not be
		check.Eq(t, Report(map[string]int{"b": 2, "a": 1, "z": -3}), "a=1 b=2 z=-3")
	}
}

func TestEx6_6(t *testing.T) {
	check.Eq(t, Busiest([]Point{{1, 2}, {0, 0}, {1, 2}, {0, 0}, {3, 3}}), Point{0, 0})
}

func TestEx6_7(t *testing.T) {
	i, j, ok := TwoSum([]int{2, 7, 11, 15}, 9)
	check.Eq(t, []any{i, j, ok}, []any{0, 1, true})
	_, _, ok = TwoSum([]int{1, 2}, 9)
	check.Eq(t, ok, false)
}

func TestEx6_8(t *testing.T) {
	a, b := set(1, 2, 3, 4), set(3, 4, 5)
	check.Eq(t, Intersect(a, b), set(3, 4))
	check.Eq(t, Intersect(b, a), set(3, 4))
	check.Eq(t, Difference(a, b), set(1, 2))
	check.Eq(t, a, set(1, 2, 3, 4))
}

func TestEx6_9(t *testing.T) {
	check.Eq(t, DistinctInWindows([]int{1, 2, 1, 3, 4, 2, 3}, 4), []int{3, 4, 4, 3})
}

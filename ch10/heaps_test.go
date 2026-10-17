package ch10

import (
	"testing"

	"cp/internal/check"
)

func TestEx10_1(t *testing.T) {
	check.True(t, Balanced("{[()]}") && !Balanced("([)]") && !Balanced("((") && !Balanced(")"), "balanced")
}

func TestEx10_2(t *testing.T) {
	check.Eq(t, []int{ShortestPath([]string{"..#", "#..", "#.."}), ShortestPath([]string{".#", "#."})}, []int{4, -1})
}

func TestEx10_3(t *testing.T) {
	check.Eq(t, DrainAscending([]int{5, 1, 4}), []int{1, 4, 5})
}

func TestEx10_4(t *testing.T) {
	in := []int{5, 1, 4}
	check.Eq(t, DrainDescending(in), []int{5, 4, 1})
	check.Eq(t, in, []int{5, 1, 4})
}

func TestEx10_5(t *testing.T) {
	check.Eq(t, KLargest([]int{7, 2, 9, 4, 9, 1}, 3), []int{9, 9, 7})
}

func TestEx10_6(t *testing.T) {
	check.Eq(t, ProcessByCost([]Job{{5, 1}, {2, 7}, {5, 0}}), []int{7, 0, 1})
}

func TestEx10_7(t *testing.T) {
	check.Eq(t, RunningMedian([]int{5, 15, 1, 3}), []float64{5, 10, 5, 4})
}

func TestEx10_8(t *testing.T) {
	check.Eq(t, WindowMax([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3), []int{3, 3, 5, 5, 6, 7})
}

func TestEx10_9(t *testing.T) {
	c := NewLRU(2)
	c.Put(1, 10)
	c.Put(2, 20)
	v, ok := c.Get(1)
	check.Eq(t, []any{v, ok}, []any{10, true})
	c.Put(3, 30)
	_, ok = c.Get(2)
	check.Eq(t, ok, false)
	c.Put(1, 11)
	v, _ = c.Get(1)
	check.Eq(t, v, 11)
	v, _ = c.Get(3)
	check.Eq(t, v, 30)
}

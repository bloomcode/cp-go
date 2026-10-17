package ch09

import (
	"testing"

	"cp/internal/check"
)

func TestEx9_1(t *testing.T) {
	check.Eq(t, MaxOf(3, 9, 2), 9)
	check.Eq(t, MaxOf("pear", "apple"), "pear")
	check.Eq(t, MaxOf(2.5), 2.5)
}

func TestEx9_2(t *testing.T) {
	isEven := func(x int) bool { return x%2 == 0 }
	square := func(x int) int { return x * x }
	add := func(a, x int) int { return a + x }
	check.Eq(t, Reduce(Map(Filter([]int{1, 2, 3, 4}, isEven), square), 0, add), 20)
	check.Eq(t, Map([]int{1, 2}, func(x int) string { return string(rune('a' + x)) }), []string{"b", "c"})
}

func TestEx9_3(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	v, ok := s.Pop()
	check.Eq(t, []any{v, ok}, []any{2, true})
	v, ok = s.Peek()
	check.Eq(t, []any{v, ok, s.Len()}, []any{1, true, 1})
	s.Pop()
	v, ok = s.Pop()
	check.Eq(t, []any{v, ok}, []any{0, false})
}

func TestEx9_4(t *testing.T) {
	check.Eq(t, Swap(Pair[string, int]{"x", 1}), Pair[int, string]{1, "x"})
}

func TestEx9_5(t *testing.T) {
	set := SetOf(1, 2, 2, 3)
	check.Eq(t, set.Len(), 3)
	check.True(t, set.Has(2) && !set.Has(5), "set")
}

func TestEx9_6(t *testing.T) {
	check.Eq(t, []int64{Sum([]int32{1, 2, 3}), Sum([]Celsius{10, 20})}, []int64{6, 30})
}

func TestEx9_7(t *testing.T) {
	check.Eq(t, []int{LowerBound([]int{10, 20, 20, 30}, 20), LowerBound([]string{"a", "c"}, "b"), LowerBound([]int{}, 5)}, []int{1, 1, 0})
}

func TestEx9_8(t *testing.T) {
	var q Queue[int]
	q.buf = make([]int, 2)
	for i := 1; i <= 5; i++ {
		q.Push(i)
	}
	var got []int
	for q.Len() > 0 {
		x, _ := q.Pop()
		got = append(got, x)
	}
	check.Eq(t, got, []int{1, 2, 3, 4, 5})
	var q2 Queue[string]
	for round := 0; round < 100; round++ {
		q2.Push("a")
		q2.Push("b")
		x, _ := q2.Pop()
		check.Eq(t, x, map[bool]string{true: "a", false: "b"}[round%2 == 0])
	}
	check.Eq(t, q2.Len(), 100)
}

package ch11

import (
	"strings"
	"testing"

	"cp/internal/check"
)

func TestEx11_1(t *testing.T) {
	check.Eq(t, Apply([]int{1, 2, 3}, func(x int) int { return x * 10 }), []int{10, 20, 30})
}

func TestEx11_2(t *testing.T) {
	c1, c2 := MakeCounter(), MakeCounter()
	c1()
	c1()
	check.Eq(t, []int{c1(), c2()}, []int{3, 1})
}

func TestEx11_3(t *testing.T) {
	check.Eq(t, []int{Sum(1, 2, 3), Sum([]int{4, 5}...), Sum()}, []int{6, 9, 0})
}

func TestEx11_4(t *testing.T) {
	v, err := Eval("max", 3, 8)
	check.Eq(t, []any{v, err}, []any{8, nil})
	_, err = Eval("^", 1, 2)
	check.True(t, err != nil, "unknown op")
}

func TestEx11_5(t *testing.T) {
	check.Eq(t, []int{Fib(10), Fib(90)}, []int{55, 2880067194370816120})
}

func TestEx11_6(t *testing.T) {
	sq, calls := Memoize(func(x int) int { return x * x })
	check.Eq(t, []int{sq(1), sq(2), sq(1)}, []int{1, 4, 1})
	check.Eq(t, *calls, 2)
}

func TestEx11_7(t *testing.T) {
	q, err := SafeDiv(7, 2)
	check.Eq(t, []any{q, err}, []any{3, nil})
	q, err = SafeDiv(1, 0)
	check.Eq(t, q, 0)
	check.True(t, err != nil && strings.Contains(err.Error(), "divide by zero"), "recovered")
}

func TestEx11_8(t *testing.T) {
	w := []string{"ccc", "a", "bb", "dd"}
	SortByKey(w, func(s string) int { return len(s) })
	check.Eq(t, w, []string{"a", "bb", "dd", "ccc"})
}

package ch04

import (
	"errors"
	"strconv"
	"strings"
	"testing"

	"cp/internal/check"
)

func TestEx4_1(t *testing.T) {
	check.Eq(t, FizzBuzz(5), []string{"1", "2", "Fizz", "4", "Buzz"})
	check.Eq(t, FizzBuzz(15)[14], "FizzBuzz")
}

func TestEx4_2(t *testing.T) {
	r, c, ok := FindInGrid([][]int{{1, 2}, {3, 4}}, 3)
	check.Eq(t, []any{r, c, ok}, []any{1, 0, true})
	_, _, ok = FindInGrid([][]int{{1, 2}, {3, 4}}, 9)
	check.Eq(t, ok, false)
}

func TestEx4_3(t *testing.T) {
	check.Eq(t, []int{CollatzSteps(6), CollatzSteps(1)}, []int{8, 0})
}

func TestEx4_4(t *testing.T) {
	v, err := Calc('*', 6, 7)
	check.Eq(t, []any{v, err}, []any{42, nil})
	v, _ = Calc('/', 7, 2)
	check.Eq(t, v, 3)
	_, err = Calc('/', 1, 0)
	check.True(t, err != nil, "div zero")
	_, err = Calc('%', 1, 2)
	check.True(t, err != nil && strings.Contains(err.Error(), "'%'"), "unknown op")
}

func TestEx4_5(t *testing.T) {
	check.Eq(t, Squares(4), []int{0, 1, 4, 9})
	check.Eq(t, Closures(3), []int{0, 1, 2})
}

func TestEx4_6(t *testing.T) {
	check.Eq(t, DeferOrder(), []string{"third", "second", "first"})
}

func TestEx4_7(t *testing.T) {
	nums, err := ParseAll([]string{"1", "22"})
	check.Eq(t, []any{nums, err}, []any{[]int{1, 22}, nil})
	_, err = ParseAll([]string{"1", "x"})
	check.True(t, err != nil && strings.Contains(err.Error(), "word 1"), "wrapped")
	check.True(t, errors.Is(err, strconv.ErrSyntax), "errors.Is")
}

func TestEx4_8(t *testing.T) {
	check.Eq(t, []int{SumUntilNegative([]int{4, 0, 3, -1, 10}), SumUntilNegative(nil)}, []int{7, 0})
}

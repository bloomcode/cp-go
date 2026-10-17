package ch08

import (
	"fmt"
	"math"
	"testing"

	"cp/internal/check"
)

func TestEx8_1(t *testing.T) {
	var c Counter
	c.Add(2)
	c.Add(5)
	check.Eq(t, c.Value(), 7)
	check.True(t, c.Add(-1) != nil, "negative rejected")
	check.Eq(t, c.Value(), 7)
}

func TestEx8_2(t *testing.T) {
	b := Box{}
	b.IncValue()
	b.IncPointer()
	check.Eq(t, b.N, 1)
}

func TestEx8_3(t *testing.T) {
	check.Eq(t, Dog{Animal{"Rex"}, "lab"}.Describe(), "I am Rex")
}

func TestEx8_4(t *testing.T) {
	check.True(t, math.Abs(TotalArea([]Shape{Rect{2, 3}, Circle{1}})-(6+math.Pi)) < 1e-9, "area")
}

func TestEx8_5(t *testing.T) {
	check.Eq(t, fmt.Sprint(Vec{1, 2}), "(1, 2)")
	check.Eq(t, fmt.Sprintf("%v", []Vec{{3, 4}}), "[(3, 4)]")
}

func TestEx8_6(t *testing.T) {
	check.Eq(t, []string{Kind(3), Kind("s"), Kind([]int{1}), Kind(Rect{1, 1}), Kind(2.5)}, []string{"int", "string", "[]int", "Shape", "other"})
}

func TestEx8_7(t *testing.T) {
	m := map[string]int{"a": 1}
	v, err := Find(m, "a")
	check.Eq(t, []any{v, err}, []any{1, nil})
	_, err = Find(m, "zz")
	key, ok := KeyOf(fmt.Errorf("lookup: %w", err))
	check.Eq(t, []any{key, ok}, []any{"zz", true})
	_, ok = KeyOf(fmt.Errorf("other"))
	check.Eq(t, ok, false)
}

func TestEx8_8(t *testing.T) {
	a, _ := NewInterval(1, 5)
	b, _ := NewInterval(5, 9)
	check.True(t, a.Overlaps(b), "overlap")
	check.Eq(t, a.Len(), 4)
	_, err := NewInterval(3, 1)
	check.True(t, err != nil, "bad interval")
	check.True(t, a == Interval{1, 5}, "comparable")
}

// Package ch09 - Generics.
package ch09

import "cmp"

// ==== EX 9.1 | A generic Max | Core
// PROMPT: Write `func MaxOf[T cmp.Ordered](first T, rest ...T) T` that works for ints, floats and strings. Requiring one argument avoids deciding what the maximum of nothing is.
// EXAMPLE: `MaxOf(3, 9, 2)` -> `9`; `MaxOf("pear", "apple")` -> `"pear"`
// BIGO: O(N). Go 1.21+ also has the built-in `max`, but it needs a fixed number of arguments.
func MaxOf[T cmp.Ordered](first T, rest ...T) T {
	panic("todo")
}

// ==== END

// ==== EX 9.2 | Map, Filter, Reduce | Core
// PROMPT: Write `func Map[T, U any](xs []T, f func(T) U) []U`, `func Filter[T any](xs []T, keep func(T) bool) []T` and `func Reduce[T, A any](xs []T, init A, f func(A, T) A) A`.
// EXAMPLE: `Reduce(Map(Filter([]int{1, 2, 3, 4}, isEven), square), 0, add)` -> `20`
// BIGO: O(N) each. In hot contest loops a plain `for` is just as fast and often clearer.
func Map[T, U any](xs []T, f func(T) U) []U {
	panic("todo")
}

func Filter[T any](xs []T, keep func(T) bool) []T {
	panic("todo")
}

func Reduce[T, A any](xs []T, init A, f func(A, T) A) A {
	panic("todo")
}

// ==== END

// ==== EX 9.3 | A generic stack | Core
// PROMPT: Write `type Stack[T any] struct{ items []T }` with `Push`, `Pop() (T, bool)`, `Peek() (T, bool)` and `Len()`. Return the zero value and false when empty.
// EXAMPLE: push 1, push 2 -> `Pop()` is `2, true`; `Peek()` is `1, true`; on an empty stack `Pop()` is `0, false`
// BIGO: O(1) amortized per operation.
// NOTE: `var zero T` gives the zero value of any type parameter.
type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(x T) { panic("todo") }

func (s *Stack[T]) Pop() (T, bool) {
	panic("todo")
}

func (s *Stack[T]) Peek() (T, bool) {
	panic("todo")
}

func (s *Stack[T]) Len() int { panic("todo") }

// ==== END

// ==== EX 9.4 | A generic Pair | Core
// PROMPT: Write `type Pair[A, B any] struct{ First A; Second B }` and `func Swap[A, B any](p Pair[A, B]) Pair[B, A]`.
// EXAMPLE: `Swap(Pair[string, int]{"x", 1})` -> `{1 x}`
// BIGO: O(1).
type Pair[A, B any] struct {
	First  A
	Second B
}

func Swap[A, B any](p Pair[A, B]) Pair[B, A] {
	panic("todo")
}

// ==== END

// ==== EX 9.5 | A generic set | Core
// PROMPT: Write `type Set[T comparable] map[T]struct{}` with methods `Add`, `Has` and `Len`, and `func SetOf[T comparable](xs ...T) Set[T]`.
// EXAMPLE: `s := SetOf(1, 2, 2, 3)` -> `s.Len()` is `3`; `s.Has(2)` is `true`
// BIGO: O(1) average per operation. `comparable` admits every type that `==` works on.
type Set[T comparable] map[T]struct{}

func SetOf[T comparable](xs ...T) Set[T] {
	panic("todo")
}

func (s Set[T]) Add(x T) { panic("todo") }

func (s Set[T]) Has(x T) bool {
	panic("todo")
}

func (s Set[T]) Len() int { panic("todo") }

// ==== END

// ==== EX 9.6 | Constraints with ~ | Core
// PROMPT: Write `type Integer interface{ ~int | ~int32 | ~int64 }` and `func Sum[T Integer](xs []T) int64`. Define `type Celsius int` and show it is accepted because of the `~`.
// EXAMPLE: `Sum([]int32{1, 2, 3})` -> `6`; `Sum([]Celsius{10, 20})` -> `30`
// BIGO: O(N). Without `~`, only the exact type `int` would match, not types defined on top of it.
type Integer interface{ ~int | ~int32 | ~int64 }

type Celsius int

func Sum[T Integer](xs []T) int64 {
	panic("todo")
}

// ==== END

// ==== EX 9.7 | Generic binary search | Core
// PROMPT: Write `func LowerBound[T cmp.Ordered](sorted []T, target T) int` returning the first index whose value is >= target.
// EXAMPLE: `LowerBound([]int{10, 20, 20, 30}, 20)` -> `1`; `LowerBound([]string{"a", "c"}, "b")` -> `1`
// BIGO: O(log N). `slices.BinarySearch` is the library version.
func LowerBound[T cmp.Ordered](sorted []T, target T) int {
	panic("todo")
}

// ==== END

// ==== EX 9.8 | A generic ring-buffer queue | Practice
// PROMPT: Write `type Queue[T any] struct` on a slice used as a ring with `head` and `size`, doubling the slice when full. Methods: `Push`, `Pop() (T, bool)` and `Len()`.
// EXAMPLE: push 1..5 into a queue that starts with capacity 2 -> pops come out `1 2 3 4 5`
// BIGO: O(1) amortized. Popping with `q = q[1:]` also works, but the memory before `q[0]` is never reused while the slice lives.
type Queue[T any] struct {
	buf        []T
	head, size int
}

func (q *Queue[T]) Push(x T) {
	panic("todo")
}

func (q *Queue[T]) Pop() (T, bool) {
	panic("todo")
}

func (q *Queue[T]) Len() int { panic("todo") }

// ==== END

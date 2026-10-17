// Package ch11 - Functions and closures.
package ch11


// ==== EX 11.1 | Functions are values | Core
// PROMPT: Write `func Apply(xs []int, f func(int) int) []int` and call it with a named function and with a function literal.
// EXAMPLE: `Apply([]int{1, 2, 3}, func(x int) int { return x * 10 })` -> `[10 20 30]`
// BIGO: O(N) calls of f.
func Apply(xs []int, f func(int) int) []int {
	panic("todo")
}

// ==== END

// ==== EX 11.2 | Closures keep their own state | Core
// PROMPT: Write `func MakeCounter() func() int` returning a function that yields 1, 2, 3, ... on successive calls. Two counters must not share state.
// EXAMPLE: `c1, c2 := MakeCounter(), MakeCounter(); c1(); c1()` -> `c1()` is `3`, `c2()` is `1`
// BIGO: O(1) per call. The captured variable lives on as long as the closure does.
func MakeCounter() func() int {
	panic("todo")
}

// ==== END

// ==== EX 11.3 | Variadic functions | Core
// PROMPT: Write `func Sum(nums ...int) int`, then call it with separate arguments and with a slice spread as `Sum(s...)`.
// EXAMPLE: `Sum(1, 2, 3)` -> `6`; `Sum([]int{4, 5}...)` -> `9`; `Sum()` -> `0`
// BIGO: O(N). `s...` passes the slice itself, not a copy, so the callee could modify the caller's elements.
func Sum(nums ...int) int {
	panic("todo")
}

// ==== END

// ==== EX 11.4 | A table of operations | Core
// PROMPT: Write `type Op func(a, b int) int` and `func Ops() map[string]Op` returning the operations "+", "-", "*" and "max", then `func Eval(op string, a, b int) (int, error)` that looks the operation up.
// EXAMPLE: `Eval("max", 3, 8)` -> `8, nil`; `Eval("^", 1, 2)` -> an error
// BIGO: O(1). A map of functions replaces a long `switch`.
type Op func(a, b int) int

func Ops() map[string]Op {
	panic("todo")
}

func Eval(op string, a, b int) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 11.5 | Recursive closures | Core
// PROMPT: Write `func Fib(n int) int` whose body declares `var fib func(int) int` first and then assigns a closure that calls `fib` and caches results in a slice it captures.
// EXAMPLE: `Fib(10)` -> `55`; `Fib(90)` -> `2880067194370816120`
// BIGO: O(N). `fib := func(...) { ... fib(...) }` does not compile, because `fib` is not in scope inside its own initializer.
func Fib(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 11.6 | Memoize any function | Practice
// PROMPT: Write `func Memoize[K comparable, V any](f func(K) V) (func(K) V, *int)` returning a cached version of f and a pointer to a counter of real calls.
// EXAMPLE: memoize a slow square; call it with 1, 2, 1 -> results `1 4 1`, real calls `2`
// BIGO: O(1) bookkeeping per call; memory grows with the number of distinct arguments.
func Memoize[K comparable, V any](f func(K) V) (func(K) V, *int) {
	panic("todo")
}

// ==== END

// ==== EX 11.7 | Turn a panic into an error | Core
// PROMPT: Write `func SafeDiv(a, b int) (q int, err error)` that divides normally but uses `defer` and `recover()` to turn the divide-by-zero panic into an error.
// EXAMPLE: `SafeDiv(7, 2)` -> `3, nil`; `SafeDiv(1, 0)` -> `0` and an error mentioning "divide by zero"
// BIGO: O(1). `recover` only works inside a deferred function; in normal code, return errors rather than panicking.
func SafeDiv(a, b int) (q int, err error) {
	panic("todo")
}

// ==== END

// ==== EX 11.8 | Closures as sort keys | Core
// PROMPT: Write `func SortByKey(words []string, key func(string) int)` that sorts in place with `sort.SliceStable`, using the closure to compute each word's key.
// EXAMPLE: `SortByKey(w, func(s string) int { return len(s) })` on `[ccc a bb dd]` -> `[a bb dd ccc]`
// BIGO: O(N log N) comparisons, each calling `key` twice; precompute keys into a slice if `key` is expensive.
func SortByKey(words []string, key func(string) int) {
	panic("todo")
}

// ==== END

// Package ch04 - Control flow.
package ch04


// ==== EX 4.1 | FizzBuzz | Core
// PROMPT: Write `func FizzBuzz(n int) []string` for 1..n with a `switch` that has no condition (`switch { case i%15 == 0: ... }`).
// EXAMPLE: `FizzBuzz(5)` -> `[1 2 Fizz 4 Buzz]`
// BIGO: O(N).
func FizzBuzz(n int) []string {
	panic("todo")
}

// ==== END

// ==== EX 4.2 | Labeled break out of nested loops | Core
// PROMPT: Write `func FindInGrid(g [][]int, target int) (r, c int, ok bool)` that leaves both loops with one `break outer`.
// EXAMPLE: `FindInGrid([][]int{{1, 2}, {3, 4}}, 3)` -> `1, 0, true`
// BIGO: O(R * C).
func FindInGrid(g [][]int, target int) (r, c int, ok bool) {
	panic("todo")
}

// ==== END

// ==== EX 4.3 | for is also while | Core
// PROMPT: Write `func CollatzSteps(n int) int` with a condition-only `for n != 1 { ... }`: halve even n, map odd n to 3n + 1, count steps.
// EXAMPLE: `CollatzSteps(6)` -> `8`; `CollatzSteps(1)` -> `0`
// BIGO: No proven bound, hence a while-style loop. Go has only `for`.
func CollatzSteps(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 4.4 | switch, with an error result | Core
// PROMPT: Write `func Calc(op byte, a, b int) (int, error)` with a `switch op`, returning an error for an unknown operator or division by zero instead of panicking.
// EXAMPLE: `Calc('*', 6, 7)` -> `42, nil`; `Calc('/', 7, 2)` -> `3, nil`; `Calc('/', 1, 0)` -> an error
// BIGO: O(1). Go's `case` does not fall through; write `fallthrough` explicitly in the rare case you want it.
func Calc(op byte, a, b int) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 4.5 | range over an integer, and per-iteration variables | Core
// PROMPT: Write `func Squares(n int) []int` with `for i := range n` (Go 1.22+), and `func Closures(n int) []int` that stores `func() int { return i }` for each i, then calls them all.
// EXAMPLE: `Squares(4)` -> `[0 1 4 9]`; `Closures(3)` -> `[0 1 2]`
// BIGO: O(N).
// NOTE: Since Go 1.22 each iteration has its own `i`, so the closures see 0, 1, 2. In older Go they all shared one variable and printed `[3 3 3]`.
func Squares(n int) []int {
	panic("todo")
}

func Closures(n int) []int {
	panic("todo")
}

// ==== END

// ==== EX 4.6 | defer runs last-in, first-out | Core
// PROMPT: Write `func DeferOrder() (log []string)` that defers three appends ("first", "second", "third") to the named result, and predict the order.
// EXAMPLE: `DeferOrder()` -> `[third second first]`
// BIGO: O(1). Deferred calls run when the function returns, newest first, and can modify named results.
func DeferOrder() (log []string) {
	panic("todo")
}

// ==== END

// ==== EX 4.7 | Multiple results and wrapped errors | Core
// PROMPT: Write `func ParseAll(words []string) ([]int, error)` that stops at the first bad word and returns `fmt.Errorf("word %d: %w", i, err)`, so callers can still find the `strconv` error with `errors.Is(err, strconv.ErrSyntax)`.
// EXAMPLE: `ParseAll([]string{"1", "22"})` -> `[1 22], nil`; `ParseAll([]string{"1", "x"})` -> an error mentioning `word 1`
// BIGO: O(N).
func ParseAll(words []string) ([]int, error) {
	panic("todo")
}

// ==== END

// ==== EX 4.8 | break and continue | Core
// PROMPT: Write `func SumUntilNegative(a []int) int`: skip zeros with `continue`, stop at the first negative with `break`.
// EXAMPLE: `SumUntilNegative([]int{4, 0, 3, -1, 10})` -> `7`
// BIGO: O(N).
func SumUntilNegative(a []int) int {
	panic("todo")
}

// ==== END

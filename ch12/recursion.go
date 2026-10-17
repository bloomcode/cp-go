// Package ch12 - Recursion and memoization.
package ch12

// ==== EX 12.1 | Base case and recursive case | Core
// PROMPT: Write `func Power(base, exp int) int` recursively by halving the exponent.
// EXAMPLE: `Power(3, 13)` -> `1594323`; `Power(2, 0)` -> `1`
// BIGO: O(log exp).
func Power(base, exp int) int {
	panic("todo")
}

// ==== END

// ==== EX 12.2 | Memoize with a slice | Core
// PROMPT: Write `func Fib(n int) int` with a `[]int` memo filled with -1 and a recursive helper, so `Fib(90)` returns instantly.
// EXAMPLE: `Fib(10)` -> `55`; `Fib(90)` -> `2880067194370816120`
// BIGO: O(N). A slice beats a map when the states are small integers: no hashing.
func Fib(n int) int {
	panic("todo")
}

func fib(n int, memo []int) int {
	panic("todo")
}

// ==== END

// ==== EX 12.3 | Memoize with an array key | Core
// PROMPT: Write `func GridPaths(r, c int) int` counting right/down paths in an r x c grid, memoized in a `map[[2]int]int`. Arrays are comparable, so `[2]int{r, c}` is a valid key.
// EXAMPLE: `GridPaths(3, 3)` -> `6`; `GridPaths(17, 17)` -> `601080390`
// BIGO: O(R * C) states.
func GridPaths(r, c int) int {
	panic("todo")
}

func gridPaths(r, c int, memo map[[2]int]int) int {
	panic("todo")
}

// ==== END

// ==== EX 12.4 | Backtracking: all subsets | Core
// PROMPT: Write `func Subsets(a []int) [][]int` by choosing or skipping each element, appending to and truncating one shared `cur` slice. Record a copy, `append([]int{}, cur...)`, not `cur` itself.
// EXAMPLE: `Subsets([]int{1, 2})` -> `[[] [2] [1] [1 2]]`
// BIGO: O(N * 2^N).
// NOTE: Appending `cur` itself stores a view of a slice you keep overwriting, so every recorded subset ends up showing the last state. This is the most common Go backtracking bug. (`slices.Clone` would also copy, but it returns nil for an empty slice.)
func Subsets(a []int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 12.5 | Backtracking: permutations | Core
// PROMPT: Write `func Permutations(s string) []string` for distinct characters, using a `[]bool` of used positions and a `[]byte` buffer.
// EXAMPLE: `Permutations("abc")` -> `[abc acb bac bca cab cba]`
// BIGO: O(N * N!).
func Permutations(s string) []string {
	panic("todo")
}

// ==== END

// ==== EX 12.6 | Flood fill | Core
// PROMPT: Write `func CountIslands(g [][]byte) int` counting 4-connected groups of '#', marking visited cells by overwriting them with '.'.
// EXAMPLE: rows `"#.#"`, `"#.."`, `"..#"` -> `3`
// BIGO: O(R * C).
func CountIslands(g [][]byte) int {
	panic("todo")
}

// ==== END

// ==== EX 12.7 | Deep recursion is fine in Go | Practice
// PROMPT: Write `func DeepSum(n int) int` that sums 1..n recursively, and run it with n = 1,000,000.
// EXAMPLE: `DeepSum(1_000_000)` -> `500000500000`, with no stack overflow
// BIGO: O(N) time and stack. Goroutine stacks start small and grow on demand up to 1 GB, so a million frames is fine, unlike Java or Python. Infinite recursion still ends in a fatal "goroutine stack exceeds" error.
func DeepSum(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 12.8 | N-Queens count | Stretch
// PROMPT: Write `func NQueens(n int) int` counting placements with three `[]bool` slices for columns and both diagonals.
// EXAMPLE: `NQueens(4)` -> `2`; `NQueens(8)` -> `92`
// BIGO: O(N!) worst case, heavily pruned.
func NQueens(n int) int {
	panic("todo")
}

// ==== END

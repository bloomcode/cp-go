// Package ch07 - Sorting and searching.
package ch07


// ==== EX 7.1 | Sort a copy, not the caller's slice | Core
// PROMPT: Write `func SortedCopy(s []int) []int` with `slices.Clone` and `slices.Sort`, and `func SortedDesc(s []int) []int` that sorts a clone and then calls `slices.Reverse`.
// EXAMPLE: `SortedCopy([]int{3, 1, 2})` -> `[1 2 3]`; `SortedDesc` -> `[3 2 1]`; the input is unchanged
// BIGO: O(N log N). `slices.Sort` uses pattern-defeating quicksort, which cannot be pushed to O(N^2).
func SortedCopy(s []int) []int {
	panic("todo")
}

func SortedDesc(s []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 7.2 | Multi-key ordering with cmp.Compare | Core
// PROMPT: Declare `type Player struct{ Name string; Score int }` and write `func Leaderboard(ps []Player) []Player`: highest score first, ties alphabetical, using `slices.SortFunc` with `cmp.Compare` and `cmp.Or`.
// EXAMPLE: `[{bo 80} {al 95} {cy 80}]` -> `[{al 95} {bo 80} {cy 80}]`
// BIGO: O(N log N).
// NOTE: `cmp.Or` returns its first non-zero argument, which chains tie-breakers neatly. Swap the arguments of `cmp.Compare` for descending order.
type Player struct {
	Name  string
	Score int
}

func Leaderboard(ps []Player) []Player {
	panic("todo")
}

// ==== END

// ==== EX 7.3 | Stable sorting | Core
// PROMPT: Declare `type Person struct{ Name, Dept string; Age int }` and write `func ByDeptThenAge(ps []Person) []Person` by sorting by age, then stably by dept with `slices.SortStableFunc`.
// EXAMPLE: `[{a ops 40} {b dev 30} {c ops 25}]` -> names `[b c a]`
// BIGO: O(N log N). `slices.SortFunc` is not stable; only the `Stable` variants keep equal items in order.
type Person struct {
	Name, Dept string
	Age        int
}

func ByDeptThenAge(ps []Person) []Person {
	panic("todo")
}

// ==== END

// ==== EX 7.4 | Argsort | Core
// PROMPT: Write `func Argsort(a []int) []int` returning indices ordered by their values: build `0..n-1` and sort it with a comparator that looks values up in `a`.
// EXAMPLE: `Argsort([]int{30, 10, 20})` -> `[1 2 0]`
// BIGO: O(N log N).
func Argsort(a []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 7.5 | Count values in a range | Core
// PROMPT: Write `func CountInRange(sorted []int, lo, hi int) int` for the inclusive range [lo, hi] using `sort.SearchInts` twice (it returns the first index whose value is >= x).
// EXAMPLE: `CountInRange([]int{1, 3, 3, 5, 8}, 3, 5)` -> `3`; `CountInRange(..., 6, 7)` -> `0`
// BIGO: O(log N). `slices.BinarySearch` also returns the insertion point, plus whether the value was found.
func CountInRange(sorted []int, lo, hi int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.6 | sort.Search finds the first true | Core
// PROMPT: `sort.Search(n, f)` returns the smallest i in [0, n) for which f(i) is true, assuming f flips from false to true once. Write `func IntSqrt(n int) int` as `sort.Search(n+1, func(i int) bool { return i > n/i ... })` minus one, avoiding overflow.
// EXAMPLE: `IntSqrt(27)` -> `5`; `IntSqrt(36)` -> `6`; `IntSqrt(0)` -> `0`
// BIGO: O(log N). This "binary search on the answer" shape solves many contest problems (see 16.3).
func IntSqrt(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.7 | Largest number from concatenation | Practice
// PROMPT: Write `func LargestNumber(nums []int) string`: sort the numbers as strings so that x comes before y when `y + x < x + y`.
// EXAMPLE: `LargestNumber([]int{3, 30, 34, 5, 9})` -> `"9534330"`; `LargestNumber([]int{0, 0})` -> `"0"`
// BIGO: O(N log N) comparisons.
func LargestNumber(nums []int) string {
	panic("todo")
}

// ==== END

// ==== EX 7.8 | Merge overlapping intervals | Core
// PROMPT: Write `func MergeIntervals(iv [][2]int) [][2]int`: sort by start, then extend the last merged interval or start a new one.
// EXAMPLE: `[[1 3] [8 10] [2 6] [15 18]]` -> `[[1 6] [8 10] [15 18]]`
// BIGO: O(N log N) for the sort, O(N) for the sweep.
func MergeIntervals(iv [][2]int) [][2]int {
	panic("todo")
}

// ==== END

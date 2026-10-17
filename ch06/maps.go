// Package ch06 - Maps and sets.
package ch06


// ==== EX 6.1 | Counting with a map | Core
// PROMPT: Write `func WordCounts(text string) map[string]int` (lowercased, split with `strings.Fields`). `counts[w]++` works even for a new key because a missing key reads as the zero value.
// EXAMPLE: `WordCounts("To be or not to be")` -> `map[be:2 not:1 or:1 to:2]`
// BIGO: O(N) words, O(1) average per update.
func WordCounts(text string) map[string]int {
	panic("todo")
}

// ==== END

// ==== EX 6.2 | Comma-ok: missing versus zero | Core
// PROMPT: Write `func Lookup(m map[string]int, key string) (int, bool)` with the comma-ok form `v, ok := m[key]`, so a stored 0 is distinguishable from a missing key.
// EXAMPLE: for `{"a": 0}` -> `Lookup(m, "a")` is `0, true`; `Lookup(m, "b")` is `0, false`
// BIGO: O(1) average.
func Lookup(m map[string]int, key string) (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 6.3 | A set as map[T]struct{} | Core
// PROMPT: Write `func Unique(nums []int) []int` keeping first occurrences in order, with a `map[int]struct{}` as the "seen" set.
// EXAMPLE: `Unique([]int{3, 1, 3, 2, 1})` -> `[3 1 2]`
// BIGO: O(N). `struct{}` takes no memory, so the map stores only keys.
func Unique(nums []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 6.4 | Grouping into slices | Core
// PROMPT: Write `func GroupByFirst(words []string) map[byte][]string`. Appending to `groups[k]` works for a missing key because appending to a nil slice allocates one.
// EXAMPLE: `[apple bat avocado]` -> `map[97:[apple avocado] 98:[bat]]` (97 is 'a')
// BIGO: O(N).
func GroupByFirst(words []string) map[byte][]string {
	panic("todo")
}

// ==== END

// ==== EX 6.5 | Map order is random, so sort the keys | Core
// PROMPT: Write `func SortedKeys(m map[string]int) []string` with `slices.Sorted(maps.Keys(m))` (Go 1.23+), and `func Report(m map[string]int) string` printing "k=v" pairs in key order, joined by spaces.
// EXAMPLE: `Report(map[string]int{"b": 2, "a": 1})` -> `"a=1 b=2"` on every run
// BIGO: O(K log K). Go randomizes map iteration on purpose, so code that depends on it fails unpredictably.
func SortedKeys(m map[string]int) []string {
	panic("todo")
}

func Report(m map[string]int) string {
	panic("todo")
}

// ==== END

// ==== EX 6.6 | Structs and arrays as keys | Core
// PROMPT: Declare `type Point struct{ X, Y int }` and write `func Busiest(pts []Point) Point` returning the most frequent point (ties: smallest X, then Y). Any comparable type works as a key, including `[2]int`, but not slices.
// EXAMPLE: `[{1 2} {0 0} {1 2} {0 0} {3 3}]` -> `{0 0}`
// BIGO: O(N).
type Point struct{ X, Y int }

func Busiest(pts []Point) Point {
	panic("todo")
}

// ==== END

// ==== EX 6.7 | Two sum with a map | Core
// PROMPT: Write `func TwoSum(nums []int, target int) (int, int, bool)` returning the indices of two numbers that add to target.
// EXAMPLE: `TwoSum([]int{2, 7, 11, 15}, 9)` -> `0, 1, true`
// BIGO: O(N).
func TwoSum(nums []int, target int) (int, int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 6.8 | Set operations | Core
// PROMPT: Write `func Intersect(a, b map[int]struct{}) map[int]struct{}` (loop over the smaller set) and `func Difference(a, b map[int]struct{}) map[int]struct{}`, leaving both inputs unchanged.
// EXAMPLE: a = {1, 2, 3, 4}, b = {3, 4, 5} -> intersection `{3, 4}`, difference `{1, 2}`
// BIGO: O(min(N, M)) for the intersection, O(N) for the difference.
func Intersect(a, b map[int]struct{}) map[int]struct{} {
	panic("todo")
}

func Difference(a, b map[int]struct{}) map[int]struct{} {
	panic("todo")
}

// ==== END

// ==== EX 6.9 | Distinct values in every window | Practice
// PROMPT: Write `func DistinctInWindows(a []int, k int) []int` maintaining a count map as the window slides, and `delete`-ing a key when its count drops to zero so `len(m)` stays the distinct count.
// EXAMPLE: `DistinctInWindows([]int{1, 2, 1, 3, 4, 2, 3}, 4)` -> `[3 4 4 3]`
// BIGO: O(N).
func DistinctInWindows(a []int, k int) []int {
	panic("todo")
}

// ==== END

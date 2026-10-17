// Package ch05 - Arrays and slices.
package ch05


// ==== EX 5.1 | Arrays are values, slices are views | Core
// PROMPT: Write `func BumpArray(a [3]int) [3]int` and `func BumpSlice(s []int)` that each add 1 to element 0. Show that the caller's array is unchanged after `BumpArray` but the caller's slice changes after `BumpSlice`.
// EXAMPLE: `a := [3]int{1, 2, 3}; BumpArray(a)` -> `a[0]` is still `1`; `s := []int{1, 2, 3}; BumpSlice(s)` -> `s[0]` is `2`
// BIGO: Passing an array copies all of it, O(N); passing a slice copies a 3-word header, O(1).
func BumpArray(a [3]int) [3]int {
	panic("todo")
}

func BumpSlice(s []int) {
	panic("todo")
}

// ==== END

// ==== EX 5.2 | The append aliasing trap | Core
// PROMPT: Write `func AliasTrap() (int, int)`: make `a := make([]int, 3, 10)`, then `b := append(a, 4)` and `c := append(a, 5)`, and return `b[3], c[3]`. Then write `func AliasFixed() (int, int)` that uses the full slice expression `a[:3:3]` so each append copies.
// EXAMPLE: `AliasTrap()` -> `5, 5` (b changed when c was appended); `AliasFixed()` -> `4, 5`
// BIGO: O(1) per append while capacity lasts, which is exactly why both appends wrote into the same backing array.
// NOTE: Two slices made by appending to the same parent can share memory. Capping the capacity with `s[lo:hi:hi]` or `slices.Clip` forces the next append to allocate.
func AliasTrap() (int, int) {
	panic("todo")
}

func AliasFixed() (int, int) {
	panic("todo")
}

// ==== END

// ==== EX 5.3 | Deep copy of a 2D slice | Core
// PROMPT: Write `func Clone2D(g [][]int) [][]int` copying every row with `slices.Clone`. Show that `copy` on the outer slice alone shares the rows.
// EXAMPLE: after `d := Clone2D(g); d[0][0] = 9` -> `g[0][0]` is unchanged
// BIGO: O(R * C).
func Clone2D(g [][]int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 5.4 | Make a filled grid | Core
// PROMPT: Write `func MakeGrid(r, c, fill int) [][]int`. Allocate one backing slice of r*c ints and cut the rows from it, so the whole grid is two allocations.
// EXAMPLE: `MakeGrid(2, 3, 7)` -> `[[7 7 7] [7 7 7]]`
// BIGO: O(R * C). New slices are zero-filled, so `fill` only matters when it is not 0.
func MakeGrid(r, c, fill int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 5.5 | Prefix sums | Core
// PROMPT: Write `func PrefixSums(a []int) []int` (length N + 1, p[0] = 0) and `func RangeSum(p []int, l, r int) int` for the inclusive range.
// EXAMPLE: `PrefixSums([]int{3, 1, 4, 1, 5})` -> `[0 3 4 8 9 14]`; `RangeSum(p, 1, 3)` -> `6`
// BIGO: O(N) build, O(1) per query.
func PrefixSums(a []int) []int {
	panic("todo")
}

func RangeSum(p []int, l, r int) int {
	panic("todo")
}

// ==== END

// ==== EX 5.6 | Delete from the middle | Core
// PROMPT: Write `func RemoveAt(s []int, i int) []int` with `append(s[:i], s[i+1:]...)`, and `func RemoveAtCopy(s []int, i int) []int` that leaves the input untouched using `slices.Delete` on a clone.
// EXAMPLE: `RemoveAt([]int{1, 2, 9, 3}, 2)` -> `[1 2 3]`
// BIGO: O(N - i): the tail shifts left, just as in an array.
// NOTE: The in-place version overwrites the caller's backing array; always use its return value and stop using the old slice.
func RemoveAt(s []int, i int) []int {
	panic("todo")
}

func RemoveAtCopy(s []int, i int) []int {
	panic("todo")
}

// ==== END

// ==== EX 5.7 | Rotate right | Core
// PROMPT: Write `func RotateRight(a []int, k int) []int` returning a new slice built with two `copy` calls. k may exceed the length.
// EXAMPLE: `RotateRight([]int{1, 2, 3, 4, 5}, 2)` -> `[4 5 1 2 3]`
// BIGO: O(N).
func RotateRight(a []int, k int) []int {
	panic("todo")
}

// ==== END

// ==== EX 5.8 | Filter in place | Core
// PROMPT: Write `func KeepPositive(s []int) []int` that reuses the input's backing array: start with `out := s[:0]` and append the values to keep.
// EXAMPLE: `KeepPositive([]int{3, -1, 4, 0, 5})` -> `[3 4 5]`, with no new allocation
// BIGO: O(N) time, O(1) extra space. Safe because writes never overtake reads.
func KeepPositive(s []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 5.9 | Spiral order | Stretch
// PROMPT: Write `func Spiral(m [][]int) []int` returning values in clockwise spiral order from the top-left.
// EXAMPLE: `Spiral([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})` -> `[1 2 3 6 9 8 7 4 5]`
// BIGO: O(R * C).
func Spiral(m [][]int) []int {
	panic("todo")
}

// ==== END
